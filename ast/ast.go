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

func (a *Assign) Id() string {
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
}

// .... CONTINUE

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

