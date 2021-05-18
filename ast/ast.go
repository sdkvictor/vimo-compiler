package ast

import (
    "foo/token"
)

type Attrib interface {}

type Program struct {
    functions 	[]*Function
	id 			string
	vars 		[]*dir.VarEntry
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
	params    []*dir.VarEntry
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

