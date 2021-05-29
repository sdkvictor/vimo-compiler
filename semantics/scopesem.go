package sem

import (
	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/mewkiz/pkg/errutil"
)

//scopeCheckProgram starts the scope checking for the whole program
func scopeCheckProgram(program *ast.Program, ctx *SemanticContext) error {
	for _, f := range program.Functions() {
		if err := scopeCheckFunction(f, ctx); err != nil {
			return err
		}
	}

	return nil
}

//scopeCheckFunction verifies the function is added to the func directory and the checks its statements
func scopeCheckFunction(function *ast.Function, ctx *SemanticContext) error {
	fe := ctx.FuncDir().Get(function.Key())
	if fe == nil {
		return errutil.NewNoPosf("%+v: Function entry %+v not found in FuncDirectory", function.Token(), fe)
	}

	if err := scopeCheckStatements(function.Statements(), fe, ctx); err != nil {
		return err
	}

	return nil
}

//scopeCheckStatements 
func scopeCheckStatements(statements []ast.Statement, fe *directories.FuncEntry, ctx *SemanticContext) error {
	for _, statement := range statements {
		if err := scopeCheckStatement(statement, fe, ctx); err != nil {
			return err
		}
	}

	return nil
}

//scopeCheckStatement calls the corresponding function to check the statement depending on the type
func scopeCheckStatement(statement ast.Statement, fe *directories.FuncEntry, ctx *SemanticContext) error {
	if vars, ok := statement.(*ast.Vars); ok {
		return createVarsInFuncEntry(vars, fe)
	} else if assign, ok := statement.(*ast.Assign); ok {
		return scopeCheckAssign(assign, ctx, fe)
	} else if condition, ok := statement.(*ast.Condition); ok {
		return scopeCheckCondition(condition, ctx, fe)
	} else if write, ok := statement.(*ast.Write); ok {
		return scopeCheckWrite(write, ctx, fe)
	} else if ret, ok := statement.(*ast.Return); ok {
		return scopeCheckReturn(ret, ctx, fe)
	} else if forstr, ok := statement.(*ast.For); ok {
		return scopeCheckFor(forstr, ctx, fe)
	} else if while, ok := statement.(*ast.While); ok {
		return scopeCheckWhile(while, ctx, fe)
	} else if fcall, ok := statement.(*ast.FunctionCall); ok {
		return scopeCheckFCall(fcall, ctx, fe)
	}

	return errutil.Newf("Statement cannot be casted to any valid form")
}

func scopeCheckAssign(assign *ast.Assign, ctx *SemanticContext, fe *directories.FuncEntry) error {
	if err := scopeCheckAttribute(assign.Attribute(), ctx, fe); err != nil {
		return err
	}

	if err := scopeCheckExpression(assign.Expression(), ctx, fe); err != nil {
		return err
	}
	
	return nil
}

func scopeCheckCondition(condition *ast.Condition, ctx *SemanticContext, fe *directories.FuncEntry) error {
	//Check Expressions
	if err := scopeCheckExpression(condition.Expression(), ctx, fe) ; err != nil {
		return err
	}
	//Check if statements
	if err := scopeCheckStatements(condition.Statements(), fe, ctx) ; err != nil {
		return err
	}
	 
	//Check else statements
	if err := scopeCheckStatements(condition.ElseStatements(), fe, ctx) ; err != nil {
		return err
	}

	return nil
}

func scopeCheckWrite(write *ast.Write, ctx *SemanticContext, fe *directories.FuncEntry) error {
	return scopeCheckExpression(write.Expression(), ctx, fe)
}

func scopeCheckReturn(ret *ast.Return, ctx *SemanticContext, fe *directories.FuncEntry) error {
	return scopeCheckExpression(ret.Expression(), ctx, fe)
}

func scopeCheckFor(forstr *ast.For, ctx *SemanticContext, fe *directories.FuncEntry) error {
	//Check Assign
	if err := scopeCheckAssign(forstr.Init(), ctx, fe) ; err != nil {
		return err
	}

	//Check Condition
	if err := scopeCheckExpression(forstr.Condition(), ctx, fe) ; err != nil {
		return err
	}

	//Check Operation
	if err := scopeCheckAssign(forstr.Operation(), ctx, fe) ; err != nil {
		return err
	}

	//Check Block Statements
	if err := scopeCheckStatements(forstr.Block(), fe, ctx) ; err != nil {
		return err
	}

	return nil
}

func scopeCheckWhile(while *ast.While, ctx *SemanticContext, fe *directories.FuncEntry) error {
	//Check expression
	if err := scopeCheckExpression(while.Expression(), ctx, fe) ; err != nil {
		return err
	}
	//Check statements
	if err := scopeCheckStatements(while.Block(), fe, ctx) ; err != nil {
		return err
	}

	return nil

}	

func scopeCheckFCall(fcall *ast.FunctionCall, ctx *SemanticContext, fe *directories.FuncEntry) error {
	//Check func directory
	if !ctx.FuncDir().Exists(fcall.Id()) {
		return errutil.Newf("Function %s not declared.", fcall.Id())
	}
	//Check params
	for _, p := range fcall.Params() {
		if err := scopeCheckExpression(p, ctx, fe); err != nil {
			return err
		}
	}
	return nil
}

func scopeCheckExpression(expr *ast.Expression, ctx *SemanticContext, fe *directories.FuncEntry) error {
	if expr == nil {
		return nil;
	}
	for _, e := range expr.Exps() {
		for _, t := range e.Terms() {
			for _, f := range t.Factors() {
				if err := scopeCheckExpression(f.Expression(), ctx, fe); err != nil {
					return err
				}
				if _, ok := f.Constant().(*ast.ConstantValue); ok {
					return nil
				} else if attr, ok := f.Constant().(*ast.Attribute); ok {
					return scopeCheckAttribute(attr, ctx, fe)
				} else if fc, ok := f.Constant().(*ast.FunctionCall); ok {
					return scopeCheckFCall(fc, ctx, fe)
				} else if le, ok := f.Constant().(*ast.ListElem); ok {
					//Check if its inside the scope
					if !fe.VarDir().Exists(le.Id()) {
						return errutil.Newf("Id %s not declared in local scope.", le.Id())
					}
					//Check if its a global variable
					if !ctx.Globals().Exists(le.Id()) {
						return errutil.Newf("Id %s not declared in local or global scope.", le.Id())
					}
				}		
			}
		}
	}
	return nil
}


func scopeCheckAttribute(attr *ast.Attribute, ctx *SemanticContext, fe *directories.FuncEntry ) error{
	//Check if objId is inside the scope
	if !fe.VarDir().Exists(attr.ObjId()) {
		//Check if objId is a global variable
		if !ctx.Globals().Exists(attr.ObjId()) {
			return errutil.Newf("Id %s not declared in local or global scope.", attr.ObjId())
		}
	}

	if(attr.VarId()!= ""){
		ok := false
		for _, att := range objectAttributes {
			if attr.VarId() == att {
				ok = true
			}
		}

		if !ok {
			return errutil.Newf("Invalid object attribute %s", attr.VarId())
		}
	}

	return nil
}


func checkVarsReserved(vars *ast.Vars) error {
	for _, res := range reservedFunctions {
		for _, v := range vars.Variables() {
			if v.Id() == res {
				return errutil.Newf("Cannot declare variable as %v since it is a reserved function", v.Id())
			}
		}
	}

	return nil
}

func createVarsInFuncEntry(vars *ast.Vars, fe *directories.FuncEntry) error {
	if err := checkVarsReserved(vars); err != nil {
		return err
	}
	if err := checkVarsInFuncEntry(vars, fe); err != nil {
		return err
	}

	for _, ve := range vars.Variables() {
		if ok := fe.VarDir().Add(ve); !ok {
			return errutil.Newf("Cannot add var %s to function directory", ve.Id())
		}
	}

	return nil
}

func checkVarsInGlobals(vars *ast.Vars, globals *directories.VarDirectory) error {
	for _, v := range vars.Variables() {
		if !globals.Exists(v.Id()) {
			return errutil.Newf("Id %s already declared in global scope.", v.Id())
		}
	}

	return nil
}

func checkVarsInFuncEntry(vars *ast.Vars, fe *directories.FuncEntry) error {
	for _, v := range vars.Variables() {
		if fe.VarDir().Exists(v.Id()) {
			return errutil.Newf("Id %s already declared in local scope.", v.Id())
		}
	}

	return nil
}
/*
//scopeCheckID
func scopeCheckID(id *ast.Id, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory) error {
	if idExistsInFuncStack(id, fes) {
		return nil
	}
	if idExistsInFuncDir(id, funcdir) {
		return nil
	}
	if IsReservedFunction(id.String()) {
		return nil
	}
	return errutil.NewNoPosf("%+v: Id %s not declared in this scope", id.Token(), id.String())
}

//scopeCheckFunctionCall
func scopeCheckFunctionCall(fcall *ast.FunctionCall, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	if err := scopeCheckStatement(fcall.Statement(), fes, funcdir, semcube); err != nil {
		return err
	}

	for _, arg := range fcall.Args() {
		if err := scopeCheckStatement(arg, fes, funcdir, semcube); err != nil {
			return err
		}
	}
	return nil
}

//scopeCheckLambda
func scopeCheckLambda(lambda *ast.Lambda, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	// Get the FuncEntry for the lambda in the top of the FuncEntry stack
	if lambdaEntry := fes.Top().GetLambdaEntryById(lambda.Id()); lambdaEntry != nil {
		// Add the new func entry to the top of the stack and then call scopeCheckStatement
		// with the updated stack
		fes.Push(lambdaEntry)
		if err := scopeCheckStatement(lambda.Statement(), fes, funcdir, semcube); err != nil {
			return err
		}
		// After checking pop the FuncEntry
		fes.Pop()

		return nil
	}
	return errutil.NewNoPosf("%+v: Lambda not in lambda list of top of FuncEntry stack", lambda.Token())
}

//scopeCheckConstantList
func scopeCheckConstantList(cl *ast.ConstantList, fes *dir.FuncEntryStack, funcdir *dir.FuncDirectory, semcube *SemanticCube) error {
	for _, s := range cl.Contents() {
		if err := scopeCheckStatement(s, fes, funcdir, semcube); err != nil {
			return err
		}
	}
	return nil
}
*/