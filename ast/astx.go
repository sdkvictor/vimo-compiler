package ast

import (
	"strconv"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/gocc/token"
	"github.com/sdkvictor/golang-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

// NewProgram creates a new program node which acts as the root of the tree
func NewProgram(id, vars, functions interface{}) (*Program, error) { //OK
	fs, ok := functions.([]*Function)
	if !ok {	
		return nil, errutil.Newf("Invalid type for functions. Expected []*Function")
	}

	idtok, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid program id. Expected string")
	}

	ids := string(idtok.Lit)

	v, ok := vars.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid type for variable declaration. Expected []*directories.VarEntry")
	}

	return &Program{fs, ids, v}, nil
}

// NewFunctionList creates a new function list of all the program's functions
func NewFunctionList(function *Function, err error) ([]*Function, error) { //OK
	if err != nil {
		return nil, err
	}
	return []*Function{function}, nil
}


func FirstFunction(typ, id, params, statements interface{}) ([]*Function, error) {
	return NewFunctionList(NewFunction(typ, id, params, statements))
}

func AppendFunction(typ, id, params, statements, flist interface{}) ([]*Function, error){
	list, ok := flist.([]*Function)
	if !ok {
		return nil, errutil.Newf("Invalid flist. Expected []*Function")
	}

	newf, err := NewFunction(typ, id, params, statements)
	return AppendFunctionList(newf, list, err)
}

// AppendFunctionList inserts a function into the program's list of functions
func AppendFunctionList(function *Function, flist []*Function, err error) ([]*Function, error) { //OK
	if err != nil {
		return nil, err
	}
	return append([]*Function{function}, flist...), nil
}

// NewFuncDecl returns a new function node
func NewFunction(typ, id, params, statements interface{}) (*Function, error) { //OK
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid program id. Expected string")
	}

	d := string(i.Lit)

	p, ok := params.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid program id. Expected string")
	}

	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.Newf("Invalid program id. Expected string")
	}

	s, ok := statements.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid program id. Expected string") 
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
		return nil, errutil.Newf("NewStatementList: Invalid type for statement. Expected Statement interface")
	}

	return []Statement{s}, nil
}

// -----OK-----
// AppendStatementList
func AppendStatementList(statement, list interface{}) ([]Statement, error) {
	l, ok := list.([]Statement)
	
	if !ok {
		return nil, errutil.Newf("Invalid type for statement list. Expected []Statement")
	}

	if s, ok := statement.(Statement); ok {
		return append([]Statement{s}, l...), nil
	}

	return nil, errutil.Newf("Invalid type for statement. Expected Statement interface got %v", statement)
}

// -----OK-----
// NewStatement creates a new Statement node which acts as the children of the function, which is the body of the function
func NewStatement(value interface{}) (Statement, error) {
	l, ok := value.(Statement)

	if !ok {
		return nil, errutil.Newf("Invalid type for statement. Expected Statement interface got %v", value)
	}

	return l, nil
}

// -----OK-----
// AppendParamsList
func AppendParamsList(typ, id, list interface{}) ([]*directories.VarEntry, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for params id. Expected token")
	}

	d := string(i.Lit)

	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.Type")
	}

	vlist, ok := list.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid type for parameters. Expected []*dir.VarEntry")
	}

	v := directories.NewVarEntry(d, t, i, len(vlist))

	vlist = append([]*directories.VarEntry{v}, vlist...)

	return vlist, nil
}

// -----OK-----
// NewParamsList
func NewParamsList(typ, id interface{}) ([]*directories.VarEntry, error) {

	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for params id. Expected token")
	}

	d := string(i.Lit)

	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.Type")
	}

	v := directories.NewVarEntry(d, t, i, 0)


	return []*directories.VarEntry{v}, nil
}

// NewType only basic types and void
func NewType(t interface{}) (*types.Type, error) { //OK
	typ, ok := t.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for type id. Expected token received %v" , t)
	}

	tstring := string(typ.Lit)

	if !ok {
		return nil, errutil.Newf("Invalid type for type. Expected string, got %v", t)
	}

	if tstring == "int" {
		return types.NewDataType(types.Int, 0, 0), nil
	}
	if tstring == "string" {
		return types.NewDataType(types.String, 0, 0), nil
	}
	if tstring == "bool" {
		return types.NewDataType(types.Bool, 0, 0), nil
	}
	if tstring == "float" {
		return types.NewDataType(types.Float, 0, 0), nil
	}
	if tstring == "char" {
		return types.NewDataType(types.Char, 0, 0), nil
	}
	if tstring == "void" {
		return types.NewDataType(types.Void, 0, 0), nil
	}
	if tstring == "Square" {
		return types.NewObjectType(types.Square, 0, 0), nil
	}
	if tstring == "Circle" {
		return types.NewObjectType(types.Circle, 0, 0), nil
	}
	if tstring == "Image" {
		return types.NewObjectType(types.Image, 0, 0), nil
	}
	if tstring == "Text" {
		return types.NewObjectType(types.Text, 0, 0), nil
	}
	if tstring == "Background" {
		return types.NewObjectType(types.Background, 0, 0), nil
	}

	return nil, errutil.Newf("Invalid type for type. Expected BasicType or ObjectType enum")
}


// AppendType
func NewTypeArray(typ, size interface{}) (*types.Type, error) { // OK?
	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.Type")
	}
	
	s, ok := size.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for size. Expected token")
	}

	sint, err := strconv.Atoi(string(s.Lit))
	if err != nil {
		return nil, errutil.Newf("Cannot parse %s to int", string(s.Lit)) 
	}

	if t.IsObject() {
		return types.NewObjectType(t.Object(), 1, sint), nil
	} else {
		return types.NewDataType(t.Basic(), 1, sint), nil
	}
}


// NewFunctionCall creates a new FunctionCall node which acts as the children of the program or function, which is the function call
func NewFunctionCall(id, exps interface{}) (*FunctionCall, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for function call id. expected Token")
	}
	
	d := string(i.Lit)

	e, ok := exps.([]*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for args. Expected []*Expression, got %v", exps)
	}

	return &FunctionCall{d, e, i}, nil
} 

// NewFunctionCallId creates a new FunctionCall node which acts as the children of the program or function, which is the function call
func NewFunctionCallId(id interface{}) (*FunctionCall, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for function call id. expected Token")
	}
	
	d := string(i.Lit)

	return &FunctionCall{d, make([]*Expression, 0), i}, nil
} 

func NewArgumentExpression(exp interface{}) ([]*Expression, error) {
	s, ok := exp.(*Expression) 
	if !ok {
		return nil, errutil.Newf("Invalid type for stmt. expected Statement")
	}

	return []*Expression{s}, nil
}

func AppendArgumentExpression(exp, list interface{}) ([]*Expression, error) {
	s, ok := exp.(*Expression) 
	if !ok {
		return nil, errutil.Newf("Invalid type for stmt. expected Statement")
	}

	slist, ok := list.([]*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for statementList. expected []Statement")
	}

	slist = append([]*Expression{s}, slist...)

	return slist, nil
}

// NewConstantBool
func NewConstantBool(value interface{}) (*ConstantValue, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for bool id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Bool, 0, 0), v, val}, nil
}

// NewConstantInt
func NewConstantInt(value interface{}) (*ConstantValue, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for int id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Int, 0, 0), v, val}, nil
}

func NewConstantId(value interface{}) (*ConstantValue, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for int id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Int, 0, 0), v, val}, nil
}

// NewConstantFloat
func NewConstantFloat(value interface{}) (*ConstantValue, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Float, 0, 0), v, val}, nil
}

// NewConstantChar
func NewConstantChar(value interface{}) (*ConstantValue, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	v := string(val.Lit)

	return &ConstantValue{types.NewDataType(types.Char, 0, 0), v, val}, nil
}

// NewConstantString
func NewConstantString(str interface{}) (*ConstantValue, error) {

    val, ok := str.(*token.Token)
    if !ok {
        return nil, errutil.Newf("Invalid type for string. Expected token")
    }

	valstr := string(val.Lit)
	
	return &ConstantValue{types.NewDataType(types.String, 0, 0), valstr, val}, nil
}

// NewReturn
func NewReturn(returnToken, exp interface{} ) (*Return, error) { //OK
	retTok, ok := returnToken.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for return keyword. Expected token")
	}
	
	if e, ok := exp.(*Expression); ok {
		return &Return{e, retTok}, nil
	}
	return nil, errutil.Newf("invalid return statement result type; expected *Expression, got %T", exp)
}

// NewWhile
func NewWhile(tok, exp, block interface{}) (*While, error) {
	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for while keyword. Expected token")
	}

	e, ok := exp.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for condition. Expected Expression")
	}

	b, ok := block.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for block. Expected []Statement")
	}


	return &While{e, b, t}, nil
}

// NewWrite
func NewWrite(tok, exp interface{}) (*Write, error) {
	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for write keyword. Expected token")
	}

	e, ok := exp.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for write. Expected Expression")
	}


	return &Write{e, t}, nil
}

// NewFor
func NewFor(tok, init, cond, op, block interface{}) (*For, error) {
	t, ok := tok.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for while keyword. Expected token")
	}

	i, ok := init.(*Assign)
	if !ok {
		return nil, errutil.Newf("Invalid type for init. Expected Assign")
	}

	c, ok := cond.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for cond. Expected Expression, got %T", cond)
	}

	o, ok := op.(*Assign)
	if !ok {
		return nil, errutil.Newf("Invalid type for operation. Expected Assign")
	}

	b, ok := block.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for block. Expected []Statement")
	}


	return &For{i, c, o, b, t}, nil
}

func NewVarsDec(vars interface{}) (*Vars, error) {
	t, ok := vars.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid type for vars keyword. Expected *[]directories.VarEntry")
	}


	return &Vars{t, t[0].Token()}, nil
}

// NewAssignWithoutAttr
func NewAssignWithoutAttr(id, exp interface{}) (*Assign, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	idstr := string(i.Lit)

	e, ok := exp.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for assign expression. Expected Expression")
	}

	a := &Attribute{idstr, "", i}

	return &Assign{a, e, i}, nil
}

// NewAssignWithAttr
func NewAssignWithAttr(attr, exp interface{}) (*Assign, error) {
	a, ok := attr.(*Attribute)
	if !ok {
		return nil, errutil.Newf("Invalid type for obj attribute. Expected Attribute")
	}

	e, ok := exp.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for assign expression. Expected Expression")
	}

	return &Assign{a, e, a.tok}, nil
}

func NewAttribute(id1, id2 interface{}) (*Attribute, error) {
	id1t, ok := id1.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for obj attribute. Expected token")
	}

	id2t, ok := id2.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for obj attribute. Expected token, got %T", id2)
	}

	return &Attribute{string(id1t.Lit), string(id2t.Lit), id1t}, nil
}

// NewCondition
func NewCondition(id, exp, stmts, elseStmts interface{}) (*Condition, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	e, ok := exp.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for obj attribute. Expected Attribute")
	}

	s, ok := stmts.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for equal assign token. Expected *token.Token")
	}

	els, ok := elseStmts.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for assign expression. Expected Expression")
	}

	return &Condition{e, s, els, i}, nil
}

// NewCondition
func NewConditionNoElseStmts(id, exp, stmts interface{}) (*Condition, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	e, ok := exp.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for obj attribute. Expected Attribute")
	}

	s, ok := stmts.([]Statement)
	if !ok {
		return nil, errutil.Newf("Invalid type for equal assign token. Expected *token.Token")
	}

	return &Condition{e, s, make([]Statement, 0), i}, nil
}

// New Expression
func NewExpression(exp interface{}) (*Expression, error) {
	e, ok := exp.(*Exp)
	if !ok {
		return nil, errutil.Newf("Invalid type for exp. Expected *Exp")
	}

	return &Expression{[]*Exp{e}, make([]string, 0), e.tok}, nil
}

// AppendExpression
func AppendExpression(exp, op, expre interface{}) (*Expression, error) {
	e, ok := exp.(*Exp)
	if !ok {
		return nil, errutil.Newf("Invalid type for exps. Expected []*Exp")
	}

	o, ok := op.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for ops. Expected []string")
	}

	ostr := string(o.Lit)

	expression, ok := expre.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for expression. Expected *Expression")
	}

	expression.exps = append([]*Exp{e}, expression.exps...)
	expression.ops = append([]string{ostr}, expression.ops...)

	return expression, nil
}

// NewExp
func NewExp(term interface{}) (*Exp, error) {
	t, ok := term.(*Term)
	if !ok {
		return nil, errutil.Newf("Invalid type for terms. Expected *Term")
	}

	return &Exp{[]*Term{t}, make([]string, 0), t.tok}, nil
}

// AppendExp
func AppendExp(term, op, exp interface{}) (*Exp, error) {
	t, ok := term.(*Term)
	if !ok {
		return nil, errutil.Newf("Invalid type for exps. Expected Term")
	}

	o, ok := op.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for ops. Expected *token.Token")
	}

	ostr := string(o.Lit)

	e, ok := exp.(*Exp)
	if !ok {
		return nil, errutil.Newf("Invalid type for exp. Expected *Exp")
	}

	e.terms = append([]*Term{t}, e.terms...)
	e.ops = append([]string{ostr}, e.ops...)

	return e, nil
}

// NewExp
func NewTerm(factor interface{}) (*Term, error) {
	f, ok := factor.(*Factor)
	if !ok {
		return nil, errutil.Newf("Invalid type for factors. Expected *Factor")
	}

	return &Term{[]*Factor{f}, make([]string, 0), f.tok}, nil
}

// AppendTerm
func AppendTerm(factor, op, term interface{}) (*Term, error) {
	f, ok := factor.(*Factor)
	if !ok {
		return nil, errutil.Newf("Invalid type for exps. Expected Factor")
	}

	o, ok := op.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for ops. Expected *token.Token")
	}

	ostr := string(o.Lit)

	t, ok := term.(*Term)
	if !ok {
		return nil, errutil.Newf("Invalid type for exp. Expected *Exp")
	}

	t.facs = append([]*Factor{f}, t.facs...)
	t.ops = append([]string{ostr}, t.ops...)

	return t, nil
}

// NewFactor
func NewFactor(expre interface{}) (*Factor, error) {
	expression, ok := expre.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for expression. Expected *Expression got %T", expre)
	}

	return &Factor{expression, nil, expression.tok}, nil
}

// NewVCFactor
func NewVCFactor(vcte interface{}) (*Factor, error) {
	vc, ok := vcte.(Constant)
	if !ok {
		return nil, errutil.Newf("Invalid type for ConstantValue. Expected Constant got %T", vcte)
	}

	return &Factor{nil, vc, vc.Token()}, nil

}

func NewVarsList(typ, ids interface{}) ([]*directories.VarEntry, error) {

	i, ok := ids.([]*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for vars id. Expected []*token.Token, got %T", ids)
	}


	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.Type")
	}

	varslst := make([]*directories.VarEntry, 0)

	for _, id := range i {
		idstr := string(id.Lit)
		v := directories.NewVarEntry(idstr, t, id, len(i))
		varslst = append([]*directories.VarEntry{v}, varslst...)
	}

	return varslst, nil
}

func AppendVarsList(typ, ids, idslist interface{}) ([]*directories.VarEntry, error) {
	i, ok := ids.([]*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for vars id. Expected token")
	}

	t, ok := typ.(*types.Type)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.Type")
	}

	l, ok := idslist.([]*directories.VarEntry)
	if !ok {
		return nil, errutil.Newf("Invalid type for typ. Expected *types.Type")
	}

	varslst := make([]*directories.VarEntry, 0)

	for _, id := range i {
		idstr := string(id.Lit)
		v := directories.NewVarEntry(idstr, t, id, len(i))
		varslst = append([]*directories.VarEntry{v}, varslst...)
	}

	return append(l, varslst...), nil
}

func NewIdList(id interface{}) ([]*token.Token, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	//idstr := string(i.Lit)
	

	return ([]*token.Token{i}), nil
}

func AppendIdList(id, list interface{}) ([]*token.Token, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for id. Expected token")
	}

	//idstr := string(i.Lit)

	idslist, ok := list.([]*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for parameters. Expected []*dir.VarEntry")
	}


	return append([]*token.Token{i}, idslist...), nil
}

func NewListElem(id, expression interface{}) (*ListElem, error) {
	i, ok := id.(*token.Token)
	if !ok {
		return nil, errutil.Newf("Invalid type for listelem id. Expected token")
	}

	idstr := string(i.Lit)

	expr, ok := expression.(*Expression)
	if !ok {
		return nil, errutil.Newf("Invalid type for list elem expression. Expected *Expresion")
	}

	return &ListElem{idstr, expr, i}, nil
}