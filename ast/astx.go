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
func NewFunctionCallClass(id, obj, exps interface{}) (*ClassFunctionCall, error) {
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

// NewFunctionCallNoAux
func NewFunctionCallNoAux(id, obj interface{}) (*ClassFunctionCall, error) {
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

	return &ClassFunctionCall{d, b, make([]*Expression, 0), i}, nil
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

func NewStringConstant(str interface{}) (*ConstantValue, error) {
    val, ok := str.(*token.Token)
    if !ok {
        return nil, errutil.NewNoPosf("Invalid type for string. Expected token")
    }

	valstr := string(val.Lit)
	
	return &ConstantValue{types.NewDataType(types.String, 0), valstr, val}, nil
}

// NewReturn
func NewReturn(returnToken, exp interface{} ) (*Return, error) { //OK
	retTok, ok := returnToken.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for return keyword. Expected token")
	}
	
	if e, ok := exp.(Expression); ok {
		return &Return{e, retTok}, nil
	}
	return nil, errutil.Newf("invalid return statement result type; expected Expression, got %T", exp)
}

// NewWhile
func NewWhile(tok, exp, block interface{}) (*While, error) {
	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for while keyword. Expected token")
	}

	e, ok := exp.(Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for condition. Expected Expression")
	}

	b, ok := typ.([]Statement])
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for typ. Expected *types.Type")
	}


	return &While{e, b, t}, nil
}

// NewFor
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

// NewAssignNoAttr probablement está mal lo del return por el attribute
func NewAssignNoAttr(id, tok, exp interface{}) (*Assign, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for id. Expected token")
	}

	idstr := string(i.Lit)

	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for equal assign token. Expected *token.Token")
	}

	e, ok := exp.(*Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for assign expression. Expected Expression")
	}

	a := new(Attribute)

	return &Assign{idstr, a, e, t}, nil
}

// NewAssignNoId probablement está mal lo del return por el id
func NewAssignNoId(attr, tok, exp interface{}) (*Assign, error) {
	a, ok := attr.(Attribute)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for obj attribute. Expected Attribute")
	}

	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for equal assign token. Expected *token.Token")
	}

	e, ok := exp.(*Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for assign expression. Expected Expression")
	}

	i := new(string)

	return &Assign{i, a, e, t}, nil
}

// NewCondition
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

// NewCondition probablemente está mal el return por el els
func NewConditionNoElseStmts(id, exp, stmts interface{}) (*Condition, error) {
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

	els := new([Statement])

	return &Condition{e, s, els, i}, nil
}

// New Expression
func NewExpression(exp interface{}) (*Expression, error) {
	e, ok := exp.(*Exp)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exp. Expected *Exp")
	}

	return &Expression{[]*Exp{e}, make(string, 0), e.tok}, nil
}

// AppendExpression
func AppendExpression(exp, op, expre interface{}) (*Expression, error) {
	e, ok := exp.(Exp)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exps. Expected []*Exp")
	}

	o, ok := op.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ops. Expected []string")
	}

	ostr := string(o.Lit)

	expression, ok := expre.(*Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for expression. Expected *Expression")
	}

	append(expression.exps, e)
	append(expression.ops, ostr)

	return expression, nil
}

// NewExp
func NewExp(term interface{}) (*Exp, error) {
	t, ok := term.(*Term)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for terms. Expected *Term")
	}

	return &Exp{[]*Term{t}, make(string, 0), t.tok}
}

// AppendExp
func AppendExp(term, op, exp interface{}) (*Exp, error) {
	t, ok := term.(Term)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exps. Expected Term")
	}

	o, ok := op.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ops. Expected *token.Token")
	}

	ostr := string(o.Lit)

	e, ok := exp.(*Exp)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exp. Expected *Exp")
	}

	append(e.terms, t)
	append(e.ops, ostr)

	return e, nil
}

// NewExp
func NewTerm(factor interface{}) (*Term, error) {
	f, ok := factor.(*Factor)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for factors. Expected *Factor")
	}

	return &Term{[]*Factor{f}, make(string, 0), f.tok}
}

// AppendTerm
func AppendTerm(factor, op, term interface{}) (*Term, error) {
	f, ok := factor.(Factor)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exps. Expected Factor")
	}

	o, ok := op.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ops. Expected *token.Token")
	}

	ostr := string(o.Lit)

	t, ok := term.(*Term)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for exp. Expected *Exp")
	}

	append(t.factors, f)
	append(t.ops, ostr)

	return t, nil
}

// NewFactor probablemente está mal el return lo de cv y op
func NewFactor(expre interface{}) (*Factor, error) {
	expression, ok := expre.(*Expression)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for expression. Expected *Expression")
	}

	cv := new(ConstantValue)
	op := new(string)

	return &Factor{expression, cv, op, expression.tok}, nil
}

// AppendFactor probablemente está mal el return lo de expre y op
func AppendFactor(op, vcte interface{}) (*Factor, error) {
	o, ok := op.(*token.Token)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for op. Expected *token.Token")
	}

	ostr := string(o.Lit)

	vc, ok := vcte.(*ConstantValue)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ConstantValue. Expected *ConstantValue")
	}

	expre := new(Expression)
	op := new(string)

	return &Factor(expre, vc, op, o), nil
}

func AppendVCFactor(vcte interface{}) (*Factor, error) {
	vc, ok := vcte.(*ConstantValue)
	if !ok {
		return nil, errutil.NewNoPosf("Invalid type for ConstantValue. Expected *ConstantValue")
	}

	expre := new(Expression)
	op := new(string)

	return &Factor(expre, vc, op, vc.tok), nil
}