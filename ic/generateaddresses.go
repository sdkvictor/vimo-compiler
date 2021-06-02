package ic

import (
	"sort"

	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Implement check to figure out if the variable space has been filled and return error

func generateAddressesProgram(program *ast.Program, ctx *GenerationContext) error {
	for _, fe := range ctx.funcdir.Table() {
		if err := generateAddressesFuncEntry(fe, ctx); err != nil {
			return err
		}
	}

	if err := generateAddressesGlobals(ctx.Globals(), ctx); err != nil {
		return err
	}

	for _, f := range program.Functions() {
		if err := generateAddressesFunction(f, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesFuncEntry(fe *directories.FuncEntry, ctx *GenerationContext) error {
	ctx.vm.ResetLocal()

	// Before we assign an address to each parameter of the function, we need to sort the
	// entries by their position in the definition. This will keep the order in runtime
	ves := make([]*directories.VarEntry, 0)

	// We extract every single var entry
	for _, ve := range fe.VarDir().Table() {
		ves = append(ves, ve)
	}

	// We sort them by their position in the function declaration
	sort.SliceStable(ves, func(i, j int) bool {
		return ves[i].Pos() < ves[j].Pos()
	})

	// Now we can request address assigning using the sorted array which
	// will guarantee that the addresses are given in order
	for _, ve := range ves {
		addr, err := ctx.vm.GetNextLocal(ve.Type())
		if err != nil {
			return err
		}
		ve.SetAddress(addr)
	}

	return nil
}

func generateAddressesGlobals(globals *directories.VarDirectory, ctx *GenerationContext) error {
	for _, ve := range globals.Table() {
		addr, err := ctx.vm.GetNextGlobal(ve.Type())
		if err != nil {
			return err
		}
		ve.SetAddress(addr)
	}

	return nil
}

func generateAddressesFunction(function *ast.Function, ctx *GenerationContext) error {
	for _, s := range function.Statements() {
		if err := generateAddressesStatement(s, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesStatement(statement ast.Statement, ctx *GenerationContext) error {
	if _, ok := statement.(*ast.Vars); ok {
		return nil
	} else if assign, ok := statement.(*ast.Assign); ok {
		return generateAddressesAssign(assign, ctx)
	} else if condition, ok := statement.(*ast.Condition); ok {
		return generateAddressesCondition(condition, ctx)
	} else if write, ok := statement.(*ast.Write); ok {
		return generateAddressesWrite(write, ctx)
	} else if ret, ok := statement.(*ast.Return); ok {
		return generateAddressesReturn(ret, ctx)
	} else if forstr, ok := statement.(*ast.For); ok {
		return generateAddressesFor(forstr, ctx)
	} else if while, ok := statement.(*ast.While); ok {
		return generateAddressesWhile(while, ctx)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return generateAddressesFunctionCall(fcall, ctx)
	}

	return errutil.NewNoPosf("Cannot cast statement to valid form")
}

func generateAddressesAssign(assign *ast.Assign, ctx *GenerationContext) error {
	if err := generateAddressesAttribute(assign.Attribute(), ctx); err != nil {
		return err
	}

	if err := generateAddressesExpression(assign.Expression(), ctx); err != nil {
		return err
	}

	return nil
}

func generateAddressesCondition(condition *ast.Condition, ctx *GenerationContext) error {
	if err := generateAddressesExpression(condition.Expression(), ctx); err != nil {
		return err
	}

	for _, s := range condition.Statements() {
		if err := generateAddressesStatement(s, ctx); err != nil {
			return err
		}
	}

	for _, s := range condition.ElseStatements() {
		if err := generateAddressesStatement(s, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesWrite(write *ast.Write, ctx *GenerationContext) error {
	if err := generateAddressesExpression(write.Expression(), ctx); err != nil {
		return err
	}

	return nil
}

func generateAddressesReturn(ret *ast.Return, ctx *GenerationContext) error {
	if err := generateAddressesExpression(ret.Expression(), ctx); err != nil {
		return err
	}

	return nil
}

func generateAddressesFor(f *ast.For, ctx *GenerationContext) error {
	if err := generateAddressesAssign(f.Init(), ctx); err != nil {
		return err
	}

	if err := generateAddressesExpression(f.Condition(), ctx); err != nil {
		return err
	}

	if err := generateAddressesAssign(f.Operation(), ctx); err != nil {
		return err
	}

	for _, s := range f.Block() {
		if err := generateAddressesStatement(s, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesWhile(w *ast.While, ctx *GenerationContext) error {
	if err := generateAddressesExpression(w.Expression(), ctx); err != nil {
		return err
	}

	for _, s := range w.Block() {
		if err := generateAddressesStatement(s, ctx); err != nil {
			return err
		}
	}

	return nil
}



func generateAddressesAttribute(att *ast.Attribute, ctx *GenerationContext) error {
	if att.Index() != nil {
		if err := generateAddressesExpression(att.Index(), ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesExpression(expression *ast.Expression, ctx *GenerationContext) error {
	for _, e := range expression.Exps() {
		if err := generateAddressesExp(e, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesExp(exp *ast.Exp, ctx *GenerationContext) error {
	for _, t := range exp.Terms() {
		if err := generateAddressesTerm(t, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesTerm(term *ast.Term, ctx *GenerationContext) error {
	for _, f := range term.Factors() {
		if err := generateAddressesFactor(f, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesFactor(factor *ast.Factor, ctx *GenerationContext) error {
	if factor.Expression() != nil {
		if err := generateAddressesExpression(factor.Expression(), ctx); err != nil {
			return err
		}

		return nil
	} else {
		if err := generateAddressesConstant(factor.Constant(), ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesConstant(c ast.Constant, ctx *GenerationContext) error {
	if cv, ok := c.(*ast.ConstantValue); ok {
		return generateAddressesConstantValue(cv, ctx)
	} else if att, ok := c.(*ast.Attribute); ok {
		return generateAddressesAttribute(att, ctx)
	} else if fc, ok := c.(*ast.FunctionCall); ok {
		return generateAddressesFunctionCall(fc, ctx)
	} else if le, ok := c.(*ast.ListElem); ok {
		return generateAddressesListElem(le, ctx)
	}

	return errutil.Newf("Cannot cast constant to any valid form %+v", c.Token())
}

func generateAddressesConstantValue(cv *ast.ConstantValue, ctx *GenerationContext) error {
	if !ctx.vm.ConstantExists(cv.Value()) {
		ctx.vm.AddConstant(cv.Value(), cv.Type())
	}

	return nil
}

func generateAddressesFunctionCall(fcall *ast.FunctionCall, ctx *GenerationContext) error {
	for _, args := range fcall.Params() {
		if err := generateAddressesExpression(args, ctx); err != nil {
			return err
		}
	}

	return nil
}

func generateAddressesListElem(le *ast.ListElem, ctx *GenerationContext) error {
	if err := generateAddressesExpression(le.Index(), ctx); err != nil {
		return err
	}

	return nil
}

