package ast

import (
    "github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/gocc/token"
	"github.com/sdkvictor/golang-compiler/types"
)

type Attrib interface {}

type Program struct {
    functions 	[]*Function
	id 			string
	vars 		[]*directories.VarEntry
}

func (p *Program) Functions() []*Function {
	return p.functions
}

func (p *Program) Id() string {
	return p.id
}

type Function struct {
	id        string
	key       string
	params    []*directories.VarEntry
	t         *types.Type
	statement Statement
	tok       *token.Token
}

func (f *Function) Id() string {
	return f.id
}

func (f *Function) Params() []*dir.VarEntry {
	return f.params
}

func (f *Function) Type() *types.LambdishType {
	return f.t
}

func (f *Function) Statement() Statement {
	return f.statement
}

func (f *Function) Key() string {
	return f.key
}

func (f *Function) CreateKey() {
	f.key = f.id
}

func (f *Function) Token() *token.Token {
	return f.tok
}

// Statement interface represents the body of the function
type Statement interface {
	isAssign() bool
	isCondition() bool
	isWrite() bool
	isReturn() bool
	isFor() bool
	isWhile() bool
	isFunctionCall() bool
	isPredefinedFunction() bool
	Token() *token.Token
}

func (s *Statement) isAssign() bool {
	return false
}

func (s *Statement) isCondition() bool {
	return false
}

func (s *Statement) isWrite() bool {
	return false
}

func (s *Statement) isReturn() bool {
	return false
}

func (s *Statement) isFor() bool {
	return false
}

func (s *Statement) isWhile() bool {
	return false
}

func (s *Statement) isFunctionCall() bool {
	return false
}

func (s *Statement) isPredefinedFunction() bool {
	return false
}

func (s *Statement) Token() *token.Token {
	return i.tok
}

// Id is a wrapper for a string to represent an id for a variable as a statement
type Id struct {
	id  string
	tok *token.Token
}

func (i *Id) Id() string {
	return i.id
}

func (i *Id) Token() *token.Token {
	return i.tok
}

type Assign struct {
	id Id
	attr Attribute
	exp Expression
	tok *token.Token
}

func (a *Assign) Id() Id {
	return a.id
}

func (a *Assign) Attribute() Attribute {
	return a.attr
}

func (a *Assign) Expression() Expression {
	return a.exp
}

func (a *Assign) Token() *token.Token {
	return a.tok
}

type Condition struct {
	exp Expression
	stmts []Statement
	elseStmts []Statement
	tok *token.Token
}

func (c *Condition) Exp() string {
	return c.exp
}

func (c *Condition) Statements() string {
	return c.stmts
}

func (c *Condition) ElseStatements() string {
	return c.elseStmts
}

func (c *Condition) Token() *token.Token {
	return c.tok
}

type Write struct {
	exp Expression
	str string
	tok *token.Token
}

func (w *Write) Expression() Expression {
	return w.exp
}

func (w *Write) String() string {
	return w.str
}

func (w *Write) Token() *token.Token {
	return w.tok
}

type Return struct {
	exp Expression
	tok *token.Token
}

func (r *Return) Expression() Expression {
	return r.exp
}

func (r *Return) Token() *token.Token {
	return r.tok
}

type For struct {
	ass Assign
	expC Expression
	expA Expression
	blck []Statement
	tok *token.Token
}

func (f *For) Assign() Assign {
	return f.ass
}

func (f *For) ExpressionCondition() Expression {
	return f.expC
}

func (f *For) ExpressionAfter() Expression {
	return f.expA
}

func (f *For) Block() []Statement {
	return f.blck
}

func (f *For) Token() *token.Token {
	return f.tok
}

type While struct {
	exp Expression
	blck []Statement
	tok *token.Token
}

func (w *While) Expression() Expression {
	return w.exp
}

func (w *While) Block() []Statement {
	return w.blck
}

func (w *While) Token() *token.Token {
	return w.tok
}

type FunctionCall struct {
	id Id
	exp Expression
	tok *token.Token
}

func (fc *FunctionCall) Id() Id {
	return fc.id
}

func (fc *FunctionCall) Expression() Expression {
	return fc.exp
}

func (fc *FunctionCall) Token() *token.Token {
	return fc.tok
}

// ConstantValue defines a type with a single basic value
type ConstantValue struct {
	t     *types.Type
	value string
	tok   *token.Token
}

func (c *ConstantValue) Type() *types.Type {
	return c.t
}

func (c *ConstantValue) Token() *token.Token {
	return c.tok
}

func (c *ConstantValue) Value() string {
	return c.value
}