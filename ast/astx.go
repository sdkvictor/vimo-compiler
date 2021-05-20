package ast

import (
	"fmt"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/gocc/token"
	"github.com/sdkvictor/golang-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// NewProgram creates a new program node which acts as the root of the tree
func NewProgram(id, vars, functions interface{}) (*Program, error) {
	fs, ok := functions.([]*Function)
	if !ok {	
		return nil, errutil.NewNoPosf("Invalid type for functions. Expected []*Function")
	}

	idtok, ok := id.(*Id)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid program id. Expected string")
	}

	v, ok := call.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for variable declaration. Expected []*directories.VarEntry")
	}

	return &Program{fs, id, v}, nil
}

// NewFunctionList creates a new function list of all the program's functions
func NewFunctionList(function interface{}) ([]*Function, error) {
	f, ok := function.(*Function)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for function. Expected *Function")
	}

	return []*Function{f}, nil
}

// AppendFunctionList inserts a function into the program's list of functions
func AppendFunctionList(function, list interface{}) ([]*Function, error) {
	f, ok := function.(*Function)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for function. Expected *Function")
	}

	flist, ok := list.([]*Function)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for functions. Expected []*Function")
	}

	return append([]*Function{f}, flist...), nil
}

// NewFuncDecl returns a new function node
func NewFunction(id, params, typ, statement interface{}) (*Function, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	p, ok := params.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for params. Expected []*directories.VarEntry")
	}

	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement")
	}

	f := &Function{d, "", p, t, s, i}
	f.CreateKey()

	return f, nil
}

// OK


// NewStatementList
func NewStatementList(statement interface{}) ([]Statement, error) {
	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.NewNoPosf("NewStatementList: Invalid type for statement. Expected Statement interface")
	}

	return []Statement{s}, nil
}

// AppendStatementList
func AppendStatementList(statement, list interface{}) ([]Statement, error) {
	l, ok := list.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement list. Expected []Statement")
	}

	// Check if the value is an id and cast it fist to a token
	if s, ok := statement.(*token.Token); ok {
		id := &Id{string(s.Lit), s}
		return append([]Statement{id}, l...), nil
		// If not, cast the value to a statement interface
	} else if s, ok := statement.(Statement); ok {
		return append([]Statement{s}, l...), nil
	}

	return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement interface got %v", statement)
}

// NewStatement creates a new Statement node which acts as the children of the function, which is the body of the function
func NewStatement(value interface{}) (Statement, error) {
	// Check if the value is an id and cast it fist to a token
	if t, ok := value.(*token.Token); ok {
		return &Id{string(t.Lit), t}, nil
		// If not, cast the value to a statement interface
	} else if s, ok := value.(Statement); ok {
		return s, nil
	}

	return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement interface got %v", value)
}

// AppendParamsList
func AppendParamsList(typ, id, list interface{}) ([]*directories.VarEntry, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	vlist, ok := list.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for parameters. Expected []*dir.VarEntry")
	}

	v := directories.NewVarEntry(d, t, i, len(vlist))

	return append([]*directories.VarEntry{v}, vlist...), nil
}

// NewParamsList
func NewParamsList(typ, id interface{}) ([]*directories.VarEntry, error) {

	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	v := directories.NewVarEntry(d, t, i, 0)

	return []*directories.VarEntry{v}, nil
}

// NewType only basic types and void
func NewType(t interface{}) (*types.Type, error) {
	typ, ok := t.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	tstring := string(typ.Lit)

	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for type. Expected string, got %v", t)
	}

	if tstring == "num" {
		return types.NewDataType(types.Num, 0), nil
	}
	if tstring == "bool" {
		return types.NewDataType(types.Bool, 0), nil
	}
	if tstring == "char" {
		return types.NewDataType(types.Char, 0), nil
	}
	if tstring == "void" {
		return types.NewDataType(types.Void, 0), nil
	}

	return nil, errutil.NewNoPosf("Invalid type for type. Expected BasicType enum")
}

// NewObjectType
func NewObjectType(params, ret interface{}) (*types.Type, error) {
	typ, ok := t.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	tstring := string(typ.Lit)

	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for type. Expected string, got %v", t)
	}

	if tstring == "Square" {
		return types.NewObjectType(types.Square, 0), nil
	}
	if tstring == "Circle" {
		return types.NewObjectType(types.Circle, 0), nil
	}
	if tstring == "Image" {
		return types.NewObjecType(types.Image, 0), nil
	}
	if tstring == "Text" {
		return types.NewObjectType(types.Text, 0), nil
	}
	if tstring == "Background" {
		return types.NewObjectType(types.Background, 0), nil
	}

	return nil, errutil.NewNoPosf("Invalid type for type. Expected ObjectType enum")
}

// AppendType
func AppendType(typ interface{}) (*types.Type, error) {
	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	if t.IsObject() {
		return types.NewObjectType(t.Retval(), t.Params(), t.List()+1), nil
	} else {
		return types.NewDataType(t.Basic(), t.List()+1), nil
	}
}

// NewFuncType
func NewFuncTypeList(typ interface{}) ([]*types.Type, error) {
	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	return []*types.Type{t}, nil
}

// NewFuncType
func AppendFuncTypeList(typ, list interface{}) ([]*types.Type, error) {
	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	l, ok := list.([]*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected []*types.Type")
	}

	return append([]*types.Type{t}, l...), nil
}

// NewFunctionCall creates a new FunctionCall node which acts as the children of the program or function, which is the function call
func NewFunctionCall(id, args interface{}) (*FunctionCall, error) {
	i, ok := id.(Id)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected statement")
	}

	e, ok := exp.(Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for args. Expected []Statement, got %v", args)
	}

	return &FunctionCall{i, e}, nil
}

// Revisar
func NewFunctionReservedCall(id, args interface{}) (*FunctionCall, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected statement")
	}

	v := string(i.Lit)
	idstruct := Id{v, i}

	a, ok := args.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for args. Expected []Statement, got %v", args)
	}

	return &FunctionCall{&idstruct, a}, nil
}


// NewConstantBool
func NewConstantBool(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Bool, 0), v, val}, nil
}

// NewConstantNum
func NewConstantNum(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Num, 0), v, val}, nil
}

// NewConstantChar
func NewConstantChar(value interface{}) (Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Char, 0), v, val}, nil
}

// AppendConstant
func AppendConstant(start, list interface{}) (*ConstantList, error) {
	val, ok := start.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for start. Expected token")
	}
	l, ok := list.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement list. Expected []Statement")
	}

	return &ConstantList{l, val, nil}, nil
}

// AppendEmptyConstant creates an empty list with a given type
func AppendEmptyConstant(start, t interface{}) (*ConstantList, error) {
	val, ok := start.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for start. Expected token")
	}

	typ, ok := t.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	// We increase the list content because the read type is always one list level lower
	// than the actual list
	typ.IncreaseList()

	return &ConstantList{make([]Statement, 0), val, typ}, nil
}

func AppendStringConstant(str interface{}) (*ConstantList, error) {
	val, ok := str.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for start. Expected token")
	}

	substr := val.Lit[1 : len(val.Lit)-1]

	chars := make([]Statement, 0)

	for _, c := range []byte(substr) {
		chars = append(chars, &ConstantValue{types.NewDataType(types.Char, 0), fmt.Sprintf("'%c'", c), val})
	}

	return &ConstantList{[]Statement(chars), val, nil}, nil
}


func NewReturn(exp, result interface{}) (*Return, error) {
	retTok, ok := returnToken.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for start. Expected token")
	}
	d := string(retTok.Lit)
	if result == nil {
		return &Return{Return: retTok.Offset}, nil
	}
	if result, ok := result.(Expression); ok {
		return &ast.ReturnStmt{Return: retTok.Offset, Result: result}, nil
	}
	return nil, errutil.Newf("invalid return statement result type; expected ast.Expr, got %T", result)
}

func NewFunction(id, params, typ, statement interface{}) (*Function, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	d := string(i.Lit)

	p, ok := params.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for params. Expected []*directories.VarEntry")
	}

	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement")
	}

	f := &Function{d, "", p, t, s, i}
	f.CreateKey()

	return f, nil
}

