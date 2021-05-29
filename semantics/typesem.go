package sem

import (
	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

//typeCheckProgram starts the type checking in the whole program
func typeCheckProgram(program *ast.Program, ctx *SemanticContext) error {
	for _, f := range program.Functions() {
		if err := typeCheckFunction(f, ctx); err != nil {
			return err
		}
	}

	return nil
}

//typeCheckFunction verifies its statement is of the same type of the return type of the function
func typeCheckFunction(function *ast.Function, ctx *SemanticContext) error {

	fe := ctx.FuncDir().Get(function.Key())
	if fe == nil {
		return errutil.NewNoPosf("%+v: Cannot get function %s from Func Directory", function.Token(), function.Id())
	}

	if err := typeCheckStatements(function.Statements(), ctx, fe); err != nil {
		return err
	}

	return nil
}

func typeCheckStatements(statements []ast.Statement, ctx *SemanticContext, fe *directories.FuncEntry) error {
	for _, s := range statements {
		if err := typeCheckStatement(s, fe, ctx); err != nil {
			return err
		}
	}

	return nil
}

func typeCheckStatement(statement ast.Statement, fe *directories.FuncEntry, ctx *SemanticContext) error {
	if _, ok := statement.(*ast.Vars); ok {
		return nil
	} else if assign, ok := statement.(*ast.Assign); ok {
		return typeCheckAssign(assign, ctx, fe)
	} else if condition, ok := statement.(*ast.Condition); ok {
		return typeCheckCondition(condition, ctx, fe)
	} else if write, ok := statement.(*ast.Write); ok {
		return typeCheckWrite(write, ctx, fe)
	} else if ret, ok := statement.(*ast.Return); ok {
		return typeCheckReturn(ret, ctx, fe)
	} else if forstr, ok := statement.(*ast.For); ok {
		return typeCheckFor(forstr, ctx, fe)
	} else if while, ok := statement.(*ast.While); ok {
		return typeCheckWhile(while, ctx, fe)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		_, err := typeCheckFunctionCall(fcall, ctx, fe)
		return err
	}

	return errutil.Newf("Statement cannot be casted to any valid form")
}

func typeCheckAssign(assign *ast.Assign, ctx *SemanticContext, fe *directories.FuncEntry) error {
	tAtt, err := typeCheckAttribute(assign.Attribute(), ctx, fe); 
	if err != nil {
		return err
	}

	tExp, err := typeCheckExpression(assign.Expression(), ctx, fe)
	if err != nil {
		return err
	}
	
	if !tAtt.Equal(tExp) {
		return errutil.Newf("%+v: Expression does not match variable type in assignment", assign.Token())
	}

	return nil
}

func typeCheckCondition(condition *ast.Condition, ctx *SemanticContext , fe *directories.FuncEntry) error {
	tExp, err := typeCheckExpression(condition.Expression(), ctx, fe)
	if err != nil {
		return err
	}

	if tExp.Basic() != types.Bool {
		return errutil.Newf("Expression does not match variable type in condition in %+v", condition.Token())
	}

	err = typeCheckStatements(condition.Statements(), ctx, fe)
	if err != nil {
		return err
	}

	 err = typeCheckStatements(condition.ElseStatements(), ctx, fe);
	if err != nil {
		return err
	}
	return nil

}

func typeCheckWrite(write *ast.Write, ctx *SemanticContext , fe *directories.FuncEntry) error {
	_, err := typeCheckExpression(write.Expression(), ctx, fe)
	if err != nil {
		return err
	}
	return nil

}

func typeCheckReturn(ret *ast.Return, ctx *SemanticContext , fe *directories.FuncEntry) error {
	tExp, err := typeCheckExpression(ret.Expression(), ctx, fe)
	if err != nil {
		return err
	}
	
	if !tExp.Equal(fe.ReturnType()) {
		return errutil.Newf("Expression does not match function type in return in %+v", ret.Token())
	}

	return nil
}

func typeCheckFor(f *ast.For, ctx *SemanticContext , fe *directories.FuncEntry) error {
	err := typeCheckAssign(f.Init(), ctx, fe)
	if err != nil {
		return err
	}

	tCond, err := typeCheckExpression(f.Condition(), ctx, fe)
	if err != nil {
		return err
	}
	
	if tCond.Basic() != types.Bool {
		return errutil.Newf("Expression does not match variable type in for condition in %+v", f.Token())
	}

	err = typeCheckAssign(f.Operation(), ctx, fe)
	if err != nil {
		return err
	}

	if err := typeCheckStatements(f.Block(), ctx, fe); err != nil {
		return err
	}

	return nil
}

func typeCheckWhile(while *ast.While, ctx *SemanticContext , fe *directories.FuncEntry) error {
	tExp, err := typeCheckExpression(while.Expression(), ctx, fe)
	if err != nil {
		return err
	}

	if tExp.Basic() != types.Bool {
		return errutil.Newf("Expression does not match variable type in while condition in %+v", while.Token())
	}

	
	if err := typeCheckStatements(while.Block(), ctx, fe); err != nil {
		return err
	}

	return nil
}

func typeCheckAttribute(att *ast.Attribute, ctx *SemanticContext, fe *directories.FuncEntry) (*types.Type, error) {
	// Caso donde es variable normal sin object attribute
	if att.VarId() == "" {
		ve := fe.VarDir().Get(att.ObjId())
		if ve == nil {
			ve = ctx.Globals().Get(att.ObjId())
			if ve == nil {
				return nil, errutil.Newf("Id %s could not be found in local or global scope. (check scopecheck?)", att.VarId())
			}
		}
		return ve.Type(), nil
	} else {
		// Caso donde es un object attribute
		t := GetObjectAttributeType(att.VarId())
		return t, nil
	}
}

func typeCheckExpression(expression *ast.Expression, ctx *SemanticContext, fe *directories.FuncEntry) (*types.Type, error) {
	prevExp := expression.Exps()[0]
	prevType, err := typeCheckExp(prevExp, ctx, fe)
	if err != nil {
		return nil, err
	}
	
	for i, op := range expression.Operations() {
		nextExp := expression.Exps()[i+1]
		nextType, err := typeCheckExp(nextExp, ctx, fe)
		if err != nil {
			return nil, err
		}
		semCubeKey := GetSemanticCubeKey(op, []*types.Type{prevType, nextType})
		basic, ok := ctx.SemCube().Get(semCubeKey)
		if !ok {
			return nil, errutil.Newf("%+v: Invalid operation %v %s %v", nextExp.Token(), prevType, op, nextType)
		}

		prevType = types.NewDataType(basic, 0, 0)
		prevExp = nextExp
	}

	return prevType, nil
}

func typeCheckExp(exp *ast.Exp, ctx *SemanticContext, fe *directories.FuncEntry) (*types.Type, error) {
	prevTerm := exp.Terms()[0]
	prevType, err := typeCheckTerm(prevTerm, ctx, fe)
	if err != nil {
		return nil, err
	}
	
	for i, op := range exp.Operations() {
		nextTerm := exp.Terms()[i+1]
		nextType, err := typeCheckTerm(nextTerm, ctx, fe)
		if err != nil {
			return nil, err
		}
		semCubeKey := GetSemanticCubeKey(op, []*types.Type{prevType, nextType})
		basic, ok := ctx.SemCube().Get(semCubeKey)
		if !ok {
			return nil, errutil.Newf("%+v: Invalid operation %v %s %v, key: %s", nextTerm.Token(), prevType, op, nextType, semCubeKey)
		}

		prevType = types.NewDataType(basic, 0, 0)
		prevTerm = nextTerm
	}

	return prevType, nil
}

func typeCheckTerm(term *ast.Term, ctx *SemanticContext, fe *directories.FuncEntry) (*types.Type, error) {
	prevFactor := term.Factors()[0]
	prevType, err := typeCheckFactor(prevFactor, ctx, fe)
	if err != nil {
		return nil, err
	}
	
	for i, op := range term.Operations() {
		nextFactor := term.Factors()[i+1]
		nextType, err := typeCheckFactor(nextFactor, ctx, fe)
		if err != nil {
			return nil, err
		}
		semCubeKey := GetSemanticCubeKey(op, []*types.Type{prevType, nextType})
		basic, ok := ctx.SemCube().Get(semCubeKey)
		if !ok {
			return nil, errutil.Newf("%+v: Invalid operation %v %s %v", prevType, op, nextType, nextFactor.Token())
		}

		prevType = types.NewDataType(basic, 0, 0)
		prevFactor = nextFactor
	}

	return prevType, nil
}

func typeCheckFactor(factor *ast.Factor, ctx *SemanticContext, fe *directories.FuncEntry) (*types.Type, error) {
	if factor.Expression() != nil {
		return typeCheckExpression(factor.Expression(), ctx, fe)
	}

	if factor.Constant() != nil {
		return typeCheckConstant(factor.Constant(), ctx, fe)
	}

	return nil, errutil.Newf("Factor expression or constant not set %+v", factor.Token())
}

func typeCheckConstant(constant ast.Constant, ctx *SemanticContext, fe *directories.FuncEntry) (*types.Type, error) {
	if cv, ok := constant.(*ast.ConstantValue); ok {
		return cv.Type(), nil
	} else if att, ok := constant.(*ast.Attribute); ok {
		return typeCheckAttribute(att, ctx, fe)
	} else if fc, ok := constant.(*ast.FunctionCall); ok {
		return typeCheckFunctionCall(fc, ctx, fe)
	} else if le, ok := constant.(*ast.ListElem); ok {
		return typeCheckListElem(le, ctx, fe)
	}

	return nil, errutil.Newf("Cannot cast constant to any valid form %+v", constant.Token())
}

func typeCheckFunctionCall(fc *ast.FunctionCall, ctx *SemanticContext, fe *directories.FuncEntry) (*types.Type, error) {
	currFe := ctx.FuncDir().Get(fc.Id())
	if currFe == nil {
		return nil, errutil.Newf("Cannot get function %s from FuncEntry", fc.Id())
	}

	argTypes := make([]*types.Type, 0)

	for _, param := range fc.Params() {
		t, err := typeCheckExpression(param, ctx, fe)
		if err != nil {
			return nil, err
		}

		argTypes = append(argTypes, t)
	}

	if len(argTypes) != len(currFe.Params()) {
		return nil, errutil.Newf("Function %s expected %v arguments, got %v", fc.Id(), len(currFe.Params()), len(argTypes))
	}

	for i := range argTypes {
		if !argTypes[i].Equal(currFe.Params()[i]) {
			return nil, errutil.Newf("In function call %s, expected type %v for position %v, got %v", fc.Id(), currFe.Params()[i], i+1, argTypes[i])
		}
	}

	return currFe.ReturnType(), nil
}

func typeCheckListElem(le *ast.ListElem, ctx *SemanticContext, fe *directories.FuncEntry) (*types.Type, error) {
	t, err := typeCheckExpression(le.Index(), ctx, fe)
	if err != nil {
		return nil, err
	}

	if t.Basic() != types.Int {
		return nil, errutil.Newf("%+v: Index for array must be type integer", le.Token())
	}

	ve := fe.VarDir().Get(le.Id())
	if ve == nil {
		ve = ctx.Globals().Get(le.Id())
	}
	if ve == nil {
		return nil, errutil.Newf("Cannot find %s in local or global scope", le.Id())
	}

	return ve.Type(), nil
}