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

type Object struct {
	id string
	key string
	t *types.Type
	tok *token.Token
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


type Write struct {
	exp Expression
	str string
	tok *token.Token
}


// .... CONTINUE

type PredefinedFunction interface {
	isSetColor()		bool
	isSetSize() 		bool
	isSetPosition() 	bool
	isMove() 			bool
	isSetText() 		bool
	isKeyPressed() 		bool
	isSetFontSize()		bool
	isRender() 			bool
	isLoadImage() 		bool
	isCheckCollision() 	bool
	isSetImage() 		bool
	isPow() 			bool
	isSquareRoot() 		bool
	Token() *token.Token
}

func (s *Statement) isSetColor() bool {
	return false
}

func (s *Statement) isMove() bool {
	return false
}
func (s *Statement) isSetText() bool {
	return false
}
func (s *Statement) isKeyPressed() bool {
	return false
}
func (s *Statement) isSetFontSize() bool {
	return false
}
func (s *Statement) isRender() bool {
	return false
}
func (s *Statement) isLoadImage() bool {
	return false
}
func (s *Statement) isCheckCollision() bool {
	return false
}
func (s *Statement) isSetImage() bool {
	return false
}
func (s *Statement) isPow() bool {
	return false
}
func (s *Statement) isSquareRoot() bool {
	return false
}

type SetColor struct {
	color string
	tok *token.Token
}

func (c *SetColor) Color() string {
	return c.color
}

func (c *SetColor) Token() *token.Token {
	return c.tok
}

type SetSize struct {
	width int
	height int
	tok *token.Token
}

func (s *SetSize) Width() int {
	return s.width
}

func (s *SetSize) Height() int {
	return s.height
}

func (s *SetSize) Token() *token.Token {
	return s.tok
}

type SetPosition struct {
	x int
	y int
	tok *token.Token
}

func (p *SetPosition) X() int {
	return p.x
}

func (p *SetPosition) Y() int {
	return p.y
}

func (p *SetPosition) Token() *token.Token {
	return p.tok
}

type Move struct {
	x int
	y int
	tok *token.Token
}

func (m *Move) X() int {
	return m.x
}

func (m *Move) Y() int {
	return m.y
}

func (m *Move) Token() *token.Token {
	return m.tok
}

type SetText struct {
	txt string
	tok *token.Token
}

func (t *SetText) Text() string {
	return t.txt
}

func (t *SetText) Token() *token.Token {
	return t.tok
}

type KeyPressed struct {
	key string
	tok *token.Token
}

func (k *KeyPressed) Key() string {
	return k.key
}

func (t *SetText) Token() *token.Token {
	return t.tok
}

type SetFontSize struct {
	size int
	tok *token.Token
}

func (f *SetFontSize) Size() int {
	return f.size
}

func (f *SetFontSize) Token() *token.Token {
	return f.tok
}

type Render struct {
	obj Object
	tok *token.Token
}

func (r *Render) Object() Object {
	return r.obj
}

func (r *Render) Token() *token.Token {
	return r.tok
}



type LoadImage struct {
	image string
	tok *token.Token
}

func (l *LoadImage) Image() string {
	return l.image
}

func (l *LoadImage) Token() *token.Token {
	return l.tok
}

type CheckCollision struct {
	obj1 Object
	obj2 Object
	tok *token.Token
}

func (l *CheckCollision) Obj1() Object {
	return l.obj1
}

func (l *CheckCollision) Obj2() Object {
	return l.obj2
}

func (l *CheckCollision) Token() *token.Token {
	return l.tok
}

type SetImage struct {
	image string
	tok *token.Token
}

func (i *SetImage) Image() string {
	return i.image
}

func (i *SetImage) Token() *token.Token {
	return i.tok
}

type Pow struct {
	base 	float32
	exp 	float32
	tok *token.Token
}

func (p *Pow) Base() float32 {
	return p.base
}

func (p *Pow) Exp() float32 {
	return p.exp
}

func (p *Pow) Token() *token.Token {
	return p.tok
}

type SquareRoot struct {
	base float32
	tok *token.Token
}

func (r *SquareRoot) Base() float32 {
	return r.base
}

func (r *SquareRoot) Token() *token.Token {
	return r.tok
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

