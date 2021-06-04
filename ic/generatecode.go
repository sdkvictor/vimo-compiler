package ic

import (
	"strconv"
	"fmt"
	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/quad"
	"github.com/sdkvictor/golang-compiler/semantics"
	"github.com/sdkvictor/golang-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// generateCodeProgram starts the code generation for the whole program
func generateCodeProgram(program *ast.Program, ctx *GenerationContext) error {

	if err := generateCodeGlobals(ctx.Globals(), ctx); err != nil {
		return err
	}

	// This statement adds the initial goto to the pending jumps
	ctx.gen.AddPendingFuncAddr(ctx.gen.ICounter(), "main")

	// The first instruction we generate jumps the execution
	ctx.gen.Generate(quad.Goto, mem.Address(-1), mem.Address(-1), mem.Address(-1))

	for _, function := range program.Functions() {
		if err := generateCodeFunction(function, ctx); err != nil {
			return err
		}
	}

	ctx.gen.FillPendingFuncAddr(ctx.funcdir)

	return nil
}

func generateCodeGlobals(globals *directories.VarDirectory, ctx *GenerationContext) error {
	for _, ve := range globals.Table() {
		if ve.Type().List() > 0 {
			nt := ve.Type().Copy()
			nt.DecreaseList()
			typeint, err := strconv.ParseInt(nt.String(), 16, 64)
			if err != nil {
				return err
			}
			ctx.gen.Generate(quad.Init, ve.Address(), mem.Address(ve.Type().Size()), mem.Address(typeint))
		} else {
			nt := ve.Type().Copy()
			nt.DecreaseList()
			typeint, err := strconv.ParseInt(nt.String(), 16, 64)
			if err != nil {
				return err
			}
			ctx.gen.Generate(quad.Init, ve.Address(), mem.Address(1), mem.Address(typeint))
		}
	}
	return nil
}

// generateCodeFunction takes a function node and generate its corresponding quadruples.
// Function nodes always end with a RET quad
func generateCodeFunction(function *ast.Function, ctx *GenerationContext) error {
	ctx.vm.ResetTemp()

	fe := ctx.funcdir.Get(function.Key())
	fe.SetLocation(ctx.gen.ICounter())

	for _, s := range function.Statements() {
		if err := generateCodeStatement(s, ctx, fe); err != nil {
			return err
		}
	}

	ctx.gen.Generate(quad.Ret, mem.Address(-1), mem.Address(-1), ctx.vm.GetDefaultAddress(fe.ReturnType()))

	return nil
}

// generateCodeStatement checks the type of the statement and calls the specific function to generate the code
func generateCodeStatement(statement ast.Statement, ctx *GenerationContext, fe *directories.FuncEntry) error {
	if vars, ok := statement.(*ast.Vars); ok {
		return generateCodeVars(vars, ctx, fe)
	} else if assign, ok := statement.(*ast.Assign); ok {
		return generateCodeAssign(assign, ctx, fe)
	} else if condition, ok := statement.(*ast.Condition); ok {
		return generateCodeCondition(condition, ctx, fe)
	} else if write, ok := statement.(*ast.Write); ok {
		return generateCodeWrite(write, ctx, fe)
	} else if ret, ok := statement.(*ast.Return); ok {
		return generateCodeReturn(ret, ctx, fe)
	} else if forstr, ok := statement.(*ast.For); ok {
		return generateCodeFor(forstr, ctx, fe)
	} else if while, ok := statement.(*ast.While); ok {
		return generateCodeWhile(while, ctx, fe)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		_, err := generateCodeFunctionCall(fcall, ctx, fe)
		if err != nil {
			return err
		}
		return nil
	}

	return errutil.Newf("Cannot cast statement to valid form: %T", statement)
}

func generateCodeVars(vars *ast.Vars, ctx *GenerationContext, fe *directories.FuncEntry) error {
	for _, ve := range vars.Variables() {
		if ve.Type().List() > 0 {
			nt := ve.Type().Copy()
			nt.DecreaseList()
			typeint, err := strconv.ParseInt(nt.String(), 16, 64)
			if err != nil {
				return err
			}
			ctx.gen.Generate(quad.Init, ve.Address(), mem.Address(ve.Type().Size()), mem.Address(typeint))
			fmt.Printf("GENERATING LIST SIZE: %d\n", ve.Type().Size())
		} else {
			nt := ve.Type().Copy()
			nt.DecreaseList()
			typeint, err := strconv.ParseInt(nt.String(), 16, 64)
			if err != nil {
				return err
			}
			ctx.gen.Generate(quad.Init, ve.Address(), mem.Address(1), mem.Address(typeint))
		}
	}
	return nil
}


func generateCodeAssign(assign *ast.Assign, ctx *GenerationContext, fe *directories.FuncEntry) error {
	addrE, err := generateCodeExpression(assign.Expression(), ctx, fe)
	if err != nil {
		return err
	}

	if err := generateCodeAttributeAssign(assign.Attribute(), ctx, fe, addrE); err != nil {
		return err
	}

	return nil
}

func generateCodeAttributeAssign(att *ast.Attribute, ctx *GenerationContext, fe *directories.FuncEntry, expTmp mem.Address) error {
	ve := fe.VarDir().Get(att.ObjId())
	if ve == nil {
		ve = ctx.Globals().Get(att.ObjId())
		if ve == nil {
			return errutil.Newf("Id %s could not be found in local or global scope. (check scopecheck?)", att.ObjId())
		}
	}

	// Caso que variable es indexada
	if att.Index() != nil {
		addrE, err := generateCodeExpression(att.Index(), ctx, fe)
		if err != nil {
			return err
		}
		ctx.gen.Generate(quad.CheckBound, mem.Address(ve.Type().Size()), mem.Address(-1), addrE)

		addrTmp, err := ctx.vm.GetNextTemp(types.NewDataType(types.Int, 0, 0))
		if err != nil {
			return err
		}

		ctx.gen.Generate(quad.AddAddr, ve.Address(), addrE, addrTmp)
		ctx.gen.Generate(quad.AssignIndex, expTmp, mem.Address(-1), addrTmp)
		// TODO: Checar estos quads en la maquina virtual

		return nil
	}

	// Caso de variable no indexada

	// Caso de variable con objeto
	if att.VarId() != "" {
		actualAddr := mem.Address(int(ve.Address()) + types.GetAttributeOffset(att.VarId()))
		ctx.gen.Generate(quad.Assign, expTmp, mem.Address(-1), actualAddr)
		return nil
	}
	
	// Caso de variable simple
	ctx.gen.Generate(quad.Assign, expTmp, mem.Address(-1), ve.Address())
	return nil
}

func generateCodeAttributeConstant(att *ast.Attribute, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	ve := fe.VarDir().Get(att.ObjId())
	if ve == nil {
		ve = ctx.Globals().Get(att.ObjId())
		if ve == nil {
			return mem.Address(-1), errutil.Newf("Id %s could not be found in local or global scope. (check scopecheck?)", att.ObjId())
		}
	}

	// Caso que variable es indexada
	if att.Index() != nil {
		addrE, err := generateCodeExpression(att.Index(), ctx, fe)
		if err != nil {
			return mem.Address(-1), err
		}
		ctx.gen.Generate(quad.CheckBound, mem.Address(ve.Type().Size()), mem.Address(-1), addrE)

		addrTmp, err := ctx.vm.GetNextTemp(types.NewDataType(types.Int, 0, 0))
		if err != nil {
			return mem.Address(-1), err
		}

		ctx.gen.Generate(quad.AddAddr, ve.Address(), addrE, addrTmp)

		finalTmp, err := ctx.vm.GetNextTemp(ve.Type())
		if err != nil {
			return mem.Address(-1), err
		}

		ctx.gen.Generate(quad.AssignIndexInv, addrTmp, mem.Address(-1), finalTmp)
		// TODO: Checar estos quads en la maquina virtual

		nt := ve.Type().Copy()
		nt.DecreaseList()
		ctx.gen.PushToTypeStack(nt)
		return finalTmp, nil
	}

	// Caso de variable no indexada

	// Caso de variable con objeto
	if att.VarId() != "" {
		// TODO: resolver cuando se resuelva addresses con objetos
		actualAddr := mem.Address(int(ve.Address()) + types.GetAttributeOffset(att.VarId()))
		ctx.gen.PushToTypeStack(semantics.GetObjectAttributeType(att.VarId()))
		return actualAddr, nil
	}
	
	// Caso de variable simple
	ctx.gen.PushToTypeStack(ve.Type())
	return ve.Address(), nil
}

func generateCodeListElem(le *ast.ListElem, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	ve := fe.VarDir().Get(le.Id())
	if ve == nil {
		ve = ctx.Globals().Get(le.Id())
		if ve == nil {
			return mem.Address(-1), errutil.Newf("Id %s could not be found in local or global scope. (check scopecheck?)", le.Id())
		}
	}

	addrE, err := generateCodeExpression(le.Index(), ctx, fe)
	if err != nil {
		return mem.Address(-1), err
	}
	ctx.gen.Generate(quad.CheckBound, mem.Address(ve.Type().Size()), mem.Address(-1), addrE)

	addrTmp, err := ctx.vm.GetNextTemp(types.NewDataType(types.Int, 0, 0))
	if err != nil {
		return mem.Address(-1), err
	}

	ctx.gen.Generate(quad.AddAddr, ve.Address(), addrE, addrTmp)

	finalTmp, err := ctx.vm.GetNextTemp(ve.Type())
	if err != nil {
		return mem.Address(-1), err
	}

	ctx.gen.Generate(quad.AssignIndexInv, addrTmp, mem.Address(-1), finalTmp)
	// TODO: Checar estos quads en la maquina virtual

	nt := ve.Type().Copy()
	nt.DecreaseList()
	ctx.gen.PushToTypeStack(nt)
	return finalTmp, nil
}

func generateCodeExpression(expression *ast.Expression, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	prevExp := expression.Exps()[0]
	prevAddr, err := generateCodeExp(prevExp, ctx, fe)
	if err != nil {
		return mem.Address(-1), err
	}
	
	for i, op := range expression.Operations() {
		nextExp := expression.Exps()[i+1]
		nextAddr, err := generateCodeExp(nextExp, ctx, fe)
		if err != nil {
			return mem.Address(-1), err
		}

		prevType := ctx.gen.GetFromTypeStack()
		nextType := ctx.gen.GetFromTypeStack()

		semCubeKey := semantics.GetSemanticCubeKey(op, []*types.Type{prevType, nextType})
		basic, ok := ctx.SemCube().Get(semCubeKey)
		if !ok {
			return mem.Address(-1), errutil.Newf("%+v: Invalid operation %v %s %v", nextExp.Token(), prevType, op, nextType)
		}

		genType := types.NewDataType(basic, 0, 0)

		tmp, err := ctx.vm.GetNextTemp(genType)
		if err != nil {
			return mem.Address(-1), err
		}

		if err := generateOperationQuad(op, prevAddr, nextAddr, tmp, ctx); err != nil {
			return mem.Address(-1), err
		}

		prevAddr = tmp
		prevExp = nextExp

		ctx.gen.PushToTypeStack(genType)
	}

	return prevAddr, nil
}

func generateCodeExp(exp *ast.Exp, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	prevTerm := exp.Terms()[0]
	prevAddr, err := generateCodeTerm(prevTerm, ctx, fe)
	if err != nil {
		return mem.Address(-1), err
	}
	
	for i, op := range exp.Operations() {
		nextTerm := exp.Terms()[i+1]
		nextAddr, err := generateCodeTerm(nextTerm, ctx, fe)
		if err != nil {
			return mem.Address(-1), err
		}

		prevType := ctx.gen.GetFromTypeStack()
		nextType := ctx.gen.GetFromTypeStack()

		semCubeKey := semantics.GetSemanticCubeKey(op, []*types.Type{prevType, nextType})
		basic, ok := ctx.SemCube().Get(semCubeKey)
		if !ok {
			return mem.Address(-1), errutil.Newf("%+v: Invalid operation %v %s %v", nextTerm.Token(), prevType, op, nextType)
		}

		genType := types.NewDataType(basic, 0, 0)

		tmp, err := ctx.vm.GetNextTemp(genType)
		if err != nil {
			return mem.Address(-1), err
		}

		if err := generateOperationQuad(op, prevAddr, nextAddr, tmp, ctx); err != nil {
			return mem.Address(-1), err
		}

		prevAddr = tmp
		prevTerm = nextTerm

		ctx.gen.PushToTypeStack(genType)
	}

	return prevAddr, nil
}

func generateCodeTerm(term *ast.Term, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	prevFactor := term.Factors()[0]
	prevAddr, err := generateCodeFactor(prevFactor, ctx, fe)
	if err != nil {
		return mem.Address(-1), err
	}
	
	for i, op := range term.Operations() {
		nextFactor := term.Factors()[i+1]
		nextAddr, err := generateCodeFactor(nextFactor, ctx, fe)
		if err != nil {
			return mem.Address(-1), err
		}

		prevType := ctx.gen.GetFromTypeStack()
		nextType := ctx.gen.GetFromTypeStack()

		semCubeKey := semantics.GetSemanticCubeKey(op, []*types.Type{prevType, nextType})
		basic, ok := ctx.SemCube().Get(semCubeKey)
		if !ok {
			return mem.Address(-1), errutil.Newf("%+v: Invalid operation %v %s %v", nextFactor.Token(), prevType, op, nextType)
		}

		genType := types.NewDataType(basic, 0, 0)

		tmp, err:= ctx.vm.GetNextTemp(genType)
		if err != nil {
			return mem.Address(-1), err
		}

		if err := generateOperationQuad(op, prevAddr, nextAddr, tmp, ctx); err != nil {
			return mem.Address(-1), err
		}

		prevAddr = tmp
		prevFactor = nextFactor

		ctx.gen.PushToTypeStack(genType)
	}

	return prevAddr, nil
}

func generateCodeFactor(factor *ast.Factor, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	if factor.Expression() != nil {
		expAddr, err := generateCodeExpression(factor.Expression(), ctx, fe)
		if err != nil {
			return mem.Address(-1), err
		}

		return expAddr, nil
	}

	return generateCodeConstant(factor.Constant(), ctx, fe)
}

func generateCodeConstant(c ast.Constant, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	if cv, ok := c.(*ast.ConstantValue); ok {
		return generateCodeConstantValue(cv, ctx, fe)
	} else if att, ok := c.(*ast.Attribute); ok {
		return generateCodeAttributeConstant(att, ctx, fe)
	} else if fc, ok := c.(*ast.FunctionCall); ok {
		return generateCodeFunctionCall(fc, ctx, fe)
	} else if le, ok := c.(*ast.ListElem); ok {
		return generateCodeListElem(le, ctx, fe)
	}

	return mem.Address(-1), errutil.Newf("Cannot cast constant to any valid form %+v", c.Token())
}

func generateCodeConstantValue(cv *ast.ConstantValue, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	ctx.gen.PushToTypeStack(cv.Type())
	return ctx.vm.GetConstantAddress(cv.Value()), nil
}

func generateCodeFunctionCall(fc *ast.FunctionCall, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	currFe := ctx.FuncDir().Get(fc.Id())

	if currFe == nil {
		if semantics.IdIsReserved(fc.Id()) {
			return generateCodeFunctionCallReserved(fc, ctx, fe)
		}
		return mem.Address(-1), errutil.Newf("Cannot find function %s in FuncDirectory", fc.Id())
	}

	ctx.gen.Generate(quad.Era, mem.Address(-1), mem.Address(-1), mem.Address(-1))

	for _, e := range fc.Params() {
		tmp, err := generateCodeExpression(e, ctx, fe)

		expType := ctx.gen.GetFromTypeStack()
		if err != nil {
			return mem.Address(-1), err
		}

		if expType.List() > 0 {
			ctx.gen.Generate(quad.Param, tmp, mem.Address(expType.Size()), mem.Address(-1))
		} else {
			ctx.gen.Generate(quad.Param, tmp, mem.Address(-1), mem.Address(-1))
		}

	}

	rettmp, err := ctx.vm.GetNextTemp(currFe.ReturnType())
	if err != nil {
		return mem.Address(-1), err
	}

	ctx.gen.AddPendingFuncAddr(ctx.gen.ICounter(), currFe.Id())
	ctx.gen.Generate(quad.Call, mem.Address(-1), mem.Address(-1), rettmp)

	ctx.gen.PushToTypeStack(currFe.ReturnType())
	
	return rettmp, nil
}

func generateCodeFunctionCallReserved(fc *ast.FunctionCall, ctx *GenerationContext, fe *directories.FuncEntry) (mem.Address, error) {
	rettype := semantics.GetReservedReturnType(fc.Id())

	addresses := []mem.Address{mem.Address(-1), mem.Address(-1)}

	for i, e := range fc.Params() {
		tmp, err := generateCodeExpression(e, ctx, fe)

		expType := ctx.gen.GetFromTypeStack()
		if err != nil {
			return mem.Address(-1), err
		}

		if expType.List() > 0 {
			return mem.Address(-1), errutil.Newf("Cannot call function %s with argument of type array", fc.Id())
		}

		addresses[i] = tmp
	}

	retTmp, err := ctx.vm.GetNextTemp(rettype)
	if err != nil {
		return mem.Address(-1), err
	}

	ctx.gen.Generate(quad.GetOperation(fc.Id()), addresses[0], addresses[1], retTmp)

	return retTmp, nil
}

func generateCodeCondition(cond *ast.Condition, ctx *GenerationContext, fe *directories.FuncEntry) error {
	exp, err := generateCodeExpression(cond.Expression(), ctx, fe)
	_ = ctx.gen.GetFromTypeStack() // Remove unused type from gen expression
	if err != nil {
		return err
	}

	ctx.gen.PushToJumpStack(mem.Address(ctx.gen.ICounter()))
	ctx.gen.Generate(quad.GotoF, exp, mem.Address(-1) , mem.Address(-1))
	for _, s := range cond.Statements() {
		if err := generateCodeStatement(s, ctx, fe); err != nil {
			return err
		}
	}
	
	if len(cond.ElseStatements()) > 0 {
		end :=  ctx.gen.GetFromJumpStack()
		ctx.gen.PushToJumpStack(mem.Address(ctx.gen.ICounter()))
		ctx.gen.Generate(quad.Goto, mem.Address(-1), mem.Address(-1) , mem.Address(-1))
		ctx.gen.FillJumpQuadruple(end, mem.Address(ctx.gen.ICounter()))
		for _, es := range cond.ElseStatements() {
			if err := generateCodeStatement(es, ctx, fe); err != nil {
				return err
			}
		}
		ctx.gen.FillJumpQuadruple(ctx.gen.GetFromJumpStack(), mem.Address(ctx.gen.ICounter()))
	} else{
		ctx.gen.FillJumpQuadruple(ctx.gen.GetFromJumpStack(), mem.Address(ctx.gen.ICounter()))
	}
	
	return nil
}

func generateCodeReturn(ret *ast.Return, ctx *GenerationContext, fe *directories.FuncEntry) error {
	expTmp, err := generateCodeExpression(ret.Expression(), ctx, fe)
	_ = ctx.gen.GetFromTypeStack() // Remove unused type from gen expression
	if err != nil {
		return err
	}

	ctx.gen.Generate(quad.Ret, mem.Address(-1), mem.Address(-1), expTmp)

	return nil
}

func generateCodeWrite(write *ast.Write, ctx *GenerationContext, fe *directories.FuncEntry) error {
	expTmp, err := generateCodeExpression(write.Expression(), ctx, fe)
	if err != nil {
		return err
	}

	_ = ctx.gen.GetFromTypeStack()

	ctx.gen.Generate(quad.Print, mem.Address(-1), mem.Address(-1), expTmp)

	return nil
}

func generateCodeFor(f *ast.For, ctx *GenerationContext, fe *directories.FuncEntry) error {
	if err := generateCodeAssign(f.Init(), ctx, fe); err != nil {
		return err
	}
	ctx.gen.PushToJumpStack(mem.Address(ctx.gen.ICounter())) // address cond
	expTmp, err := generateCodeExpression(f.Condition(), ctx, fe)
	if err != nil {
		return err
	}
	_ = ctx.gen.GetFromTypeStack()

	ctx.gen.PushToJumpStack(mem.Address(ctx.gen.ICounter())) // address gotof
	ctx.gen.Generate(quad.GotoF, expTmp, mem.Address(-1) , mem.Address(-1))
	
	for _, s := range f.Block() {
		if err := generateCodeStatement(s, ctx, fe); err != nil {
			return err
		}
	}

	if err := generateCodeAssign(f.Operation(), ctx, fe); err != nil {
		return err
	}

	GOTOF := ctx.gen.GetFromJumpStack()
	COND := ctx.gen.GetFromJumpStack()

	ctx.gen.Generate(quad.Goto, mem.Address(-1) , mem.Address(-1) , COND)
	ctx.gen.FillJumpQuadruple(GOTOF, mem.Address(ctx.gen.ICounter()))

	return nil
}

func generateCodeWhile(w *ast.While, ctx *GenerationContext, fe *directories.FuncEntry) error {
	ctx.gen.PushToJumpStack(mem.Address(ctx.gen.ICounter())) // address while
	expTmp, err := generateCodeExpression(w.Expression(), ctx, fe)
	if err != nil {
		return err
	}
	_ = ctx.gen.GetFromTypeStack()

	ctx.gen.PushToJumpStack(mem.Address(ctx.gen.ICounter())) // address gotof
	ctx.gen.Generate(quad.GotoF, expTmp, mem.Address(-1) , mem.Address(-1))
	
	for _, s := range w.Block() {
		if err := generateCodeStatement(s, ctx, fe); err != nil {
			return err
		}
	}

	GOTOF := ctx.gen.GetFromJumpStack()
	WHILE := ctx.gen.GetFromJumpStack()

	ctx.gen.Generate(quad.Goto, mem.Address(-1), mem.Address(-1) , WHILE)
	ctx.gen.FillJumpQuadruple(GOTOF, mem.Address(ctx.gen.ICounter()))
	
	return nil
}

func generateOperationQuad(op string, lop, rop, res mem.Address, ctx *GenerationContext) error {
	switch op {
	case "+":
		ctx.gen.Generate(quad.Add, lop, rop, res)
	case "-":
		ctx.gen.Generate(quad.Sub, lop, rop, res)
	case "*":
		ctx.gen.Generate(quad.Mult, lop, rop, res)
	case "/":
		ctx.gen.Generate(quad.Div, lop, rop, res)
	case "<":
		ctx.gen.Generate(quad.Lt, lop, rop, res)
	case ">":
		ctx.gen.Generate(quad.Gt, lop, rop, res)
	case "==":
		ctx.gen.Generate(quad.Equal, lop, rop, res)
	case "&&":
		ctx.gen.Generate(quad.And, lop, rop, res)
	case "||":
		ctx.gen.Generate(quad.Or, lop, rop, res)
	default:
		return errutil.Newf("Cannot get valid operation")
	}

	return nil
}
