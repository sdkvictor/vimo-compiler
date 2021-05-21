package ast

import (
	"fmt"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/gocc/token"
	"github.com/sdkvictor/golang-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// NewProgram creates a new program node which acts as the root of the tree
func NewProgram(id, vars, functions interface{}) (*Program, error) { //OK
	fs, ok := functions.([]*Function)
	if !ok {	
		return nil, errutil.NewNoPosf("Invalid type for functions. Expected []*Function")
	}

	idtok, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid program id. Expected string")
	}

	id := string(idtok.Lit)

	v, ok := call.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for variable declaration. Expected []*directories.VarEntry")
	}

	return &Program{fs, id, v}, nil
}

// NewFunctionList creates a new function list of all the program's functions
func NewFunctionList(function interface{}) ([]*Function) { //OK
	f, ok := function.(*Function)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for function. Expected *Function")
	}

	return []*Function{f}, nil
}


func FirstFunction(id, params, typ, statements interface{}) (([]*Function, error)) {
	return NewFunctionList(NewFunction(id, params, typ, statements))
}

// AppendFunctionList inserts a function into the program's list of functions
func AppendFunctionList(function, list interface{}) ([]*Function, error) { //OK
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
func NewFunction(id, params, typ, statements interface{}) (*Function, error) { //OK
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

	s, ok := statements.([]*Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement")
	}

	f := &Function{d, "", p, t, s, i}
	f.CreateKey()

	return f, nil
}

// OK

// -----OK-----
// NewStatementList
func NewStatementList(statement interface{}) ([]Statement, error) {
	s, ok := statement.(Statement)
	if !ok {
		return nil, errutil.NewNoPosf("NewStatementList: Invalid type for statement. Expected Statement interface")
	}

	return []Statement{s}, nil
}

// -----OK-----
// AppendStatementList
func AppendStatementList(statement, list interface{}) ([]Statement, error) {
	l, ok := list.([]Statement)
	
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement list. Expected []Statement")
	}

	if s, ok := statement.(Statement); ok {
		return append([]Statement{s}, l...), nil
	}

	return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement interface got %v", statement)
}

// -----OK-----
// NewStatement creates a new Statement node which acts as the children of the function, which is the body of the function
func NewStatement(value interface{}) (Statement, error) {
	l, ok := value.(Statement)

	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for statement. Expected Statement interface got %v", value)
	}

	return l, nil
}

// -----OK-----
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

// -----OK-----
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
func NewType(t interface{}) (*types.Type, error) { //OK
	typ, ok := t.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	tstring := string(typ.Lit)

	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for type. Expected string, got %v", t)
	}

	if tstring == "int" {
		return types.NewDataType(types.Int, 0), nil
	}
	if tstring == "float" {
		return types.NewDataType(types.Float, 0), nil
	}
	if tstring == "char" {
		return types.NewDataType(types.Char, 0), nil
	}
	if tstring == "void" {
		return types.NewDataType(types.Void, 0), nil
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

	return nil, errutil.NewNoPosf("Invalid type for type. Expected BasicType or ObjectType enum")
}


// AppendType
func AppendType(typ interface{}) (*types.Type, error) { // OK?
	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}

	if t.IsObject() {
		return types.NewObjectType(t.Object(), t.List()+1), nil
	} else {
		return types.NewDataType(t.Basic(), t.List()+1), nil
	}
}


// NewFunctionCall creates a new FunctionCall node which acts as the children of the program or function, which is the function call
func NewFunctionCall(id, exps interface{}) (*FunctionCall, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. expected Token")
	}
	
	d := string(id.Lit)

	e, ok := exps.([]*Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for args. Expected []*Expression, got %v", exps)
	}

	return &FunctionCall{d, e, i}, nil
}

// NewFunctionCallClass
func NewFunctionCallClass(id, obj, exps interface{}) (*ClassFunctionlCall, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected Token")
	}
	d := string(id.Lit)

	o, ok := obj.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected Token")
	}
	b := string(obj.Lit)

	e, ok := exps.([]*Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for args. Expected []Statement, got %v", exps)
	}

	return &ClassFunctionCall{d, b, e, i}, nil
}

// NewConstantBool
func NewConstantBool(value interface{}) (*ConstantValue, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Bool, 0), v, val}, nil
}

// NewConstantInt
func NewConstantInt(value interface{}) (*ConstantValue, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Int, 0), v, val}, nil
}

// NewConstantFloat
func NewConstantFloat(value interface{}) (*ConstantValue, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Float, 0), v, val}, nil
}

// NewConstantChar
func NewConstantChar(value interface{}) (*ConstantValue, error) {
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

// Posiblemente está mal
func AppendStringConstant(str interface{}) ([]*ConstantValue, error) {
    val, ok := str.(*token.Token)
    if !ok {
        return nil, errutil.NewNoPosf("Invalid type for string. Expected token")
    }

	return append([]*ConstantValue, str...), nil
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


func NewReturn(returnToken, exp interface{} ) (*Return, error) { //OK
	retTok, ok := returnToken.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for return keyword. Expected token")
	}
	
	if exp, ok := exp.(Expression); ok {
		return &Return{exp, d}, nil
	}
	return nil, errutil.Newf("invalid return statement result type; expected Expression, got %T", exp)
}



func NewWhile(tok, exp, block interface{}) (*While, error) {
	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for while keyword. Expected token")
	}

	exp, ok := exp.(Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for condition. Expected Expression")
	}

	block, ok := typ.([]Statement])
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}


	return &While{exp, block, t}, nil
}


func NewFor(tok, init, cond, op, block interface{}) (*For, error) {
	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for while keyword. Expected token")
	}

	i, ok := init.(Assign)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for init. Expected Assign")
	}

	c, ok := cond.(Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for cond. Expected Expression")
	}

	o, ok := op.(Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for operation. Expected Expression")
	}

	b, ok := block.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for block. Expected []Statement")
	}


	return &For{i, c, o, b, t}, nil
}

func NewAssign(id, attr, tok, exp interface{}) (*Assign, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	idstr := string(i.Lit)

	a, ok := attr.(Attribute)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for obj attribute. Expected Attribute")
	}

	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for equal assign token. Expected *token.Token")
	}

	e, ok := exp.(Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for assign expression. Expected Expression")
	}

	return &Assign{idstr, a, e, t}, nil
}

func NewCondition(id, exp, stmts, elseStmts interface{}) (*Condition, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	e, ok := exp.(Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for obj attribute. Expected Attribute")
	}

	s, ok := stmts.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for equal assign token. Expected *token.Token")
	}

	els, ok := elseStmts.([]Statement)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for assign expression. Expected Expression")
	}

	return &Condition{e, s, els, i}, nil
}

func NewExpression(exps, ops, tok interface{}) (*Expression, error) {
	e, ok := exps.([]*Exp)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exps. Expected []*Exp")
	}

	o, ok := ops.([]string)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ops. Expected []string")
	}

	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for tok. Expected token")
	}

	return &Expression{exps, ops, tok}, nil
}

// Así jala con la struct de expression pero no c kpdo porque en el bnf 
// donde carlos puso que se crea el objeto Expression solo tiene un
// parametro que es ExpressionAux
func NewExpression(exps, ops, tok interface{}) (*Expression, error) {
	e, ok := exps.([]*Exp)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exps. Expected []*Exp")
	}

	o, ok := ops.([]string)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ops. Expected []string")
	}

	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for tok. Expected token")
	}

	return &Expression{exps, ops, tok}, nil
}

func NewExp(terms, ops, tok interface{}) (*Exp, error) {
	t, ok := terms.([]*Term)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for terms. Expected []*Term")
	}

	o, ok := ops.([]string)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ops. Expected []string")
	}

	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for tok. Expected token")
	}

	return &Exp{terms, ops, tok}, nil
}

func NewTerm(facs, ops, tok interface{}) (*Term, error) {
	f, ok := facs.([]*Factor)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for facs. Expected []*Factor")
	}

	o, ok := ops.([]string)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ops. Expected []string")
	}

	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for tok. Expected token")
	}

	return &Exp{terms, ops, tok}, nil
}

func NewFactor(exp, cv, op, tok interface{}) (*Factor, error) {
	e, ok := exp.(Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exp. Expected Expression")
	}

	c, ok := ops.(ConstantValue)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for cv. Expected ConstantValue")
	}

	o, ok := ops.(string)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for op. Expected string")
	}

	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for tok. Expected token")
	}

	return &Factor{exp, cv, op, tok}, nil
}