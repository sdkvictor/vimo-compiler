package ast

import (
    "github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/gocc/token"
	"github.com/sdkvictor/golang-compiler/types"
)



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

func (p *Program) Vars() []*directories.VarEntry {
	return p.vars
}

/*

type Object struct {
	id string
	key string
	t *types.Type
	tok *token.Token
}
*/
type Function struct {
	id        	string
	key       	string
	params    	[]*directories.VarEntry
	t         	*types.Type
	statements 	[]Statement
	tok       	*token.Token
}


func (f *Function) Id() string {
	return f.id
}

func (f *Function) Params() []*directories.VarEntry {
	return f.params
}

func (f *Function) Type() *types.Type {
	return f.t
}

func (f *Function) Statements() []Statement {
	return f.statements
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
	isVars() 				bool
	isAssign() 				bool
	isCondition() 			bool
	isWrite() 				bool
	isReturn()				bool
	isFor() 				bool
	isWhile() 				bool
	isFunctionCall() 		bool
	isPredefinedFunction() 	bool
	Token() 				*token.Token
}



type Factor struct {
	exp 	*Expression
	cv 		Constant	
	tok 	*token.Token
}

func (f *Factor) Expression() *Expression {
	return f.exp
}

func (f *Factor) Constant() Constant {
	return f.cv
}

func (f *Factor) Token() *token.Token {
	return f.tok
}

type Term struct {
	facs 	[]*Factor
	ops		[]string
	tok 	*token.Token
}

func (t *Term) Factors() []*Factor {
	return t.facs
}

func (t *Term) Operations() []string {
	return t.ops
}

func (t *Term) Token() *token.Token {
	return t.tok
}

type Exp struct {
	terms 	[]*Term
	ops		[]string
	tok 	*token.Token
}

func (e *Exp) Terms() []*Term {
	return e.terms
}

func (e *Exp) Operations() []string {
	return e.ops
}

func (e *Exp) Token() *token.Token {
	return e.tok
}

type Expression struct {
	exps 	[]*Exp
	ops 	[]string
	tok 	*token.Token
}

func (e *Expression) Exps() []*Exp {
	return e.exps
}

func (e *Expression) Operations() []string {
	return e.ops
}

func (e *Expression) Token() *token.Token {
	return e.tok
}

type Attribute struct {
	objId string
	varId string
	tok *token.Token
}

func (a *Attribute) ObjId() string{
	return a.objId
}

func (a *Attribute) VarId() string{
	return a.varId
}

func (a Attribute) isConstantValue() bool {
	return false
}

func (a Attribute) isAttribute() bool {
	return true
}

func (a Attribute) isFunctionCall() bool {
	return false
}

func (a Attribute) isListElem() bool {
	return false
}


func (a *Attribute) Token() *token.Token{
	return a.tok
}

type ListElem struct{
	id 		string
	index 	*Expression
	tok 	*token.Token
}

func (a *ListElem) Id() string{
	return a.id
}

func (a *ListElem) Index() *Expression{
	return a.index
}

func (a ListElem) isConstantValue() bool {
	return false
}

func (a ListElem) isAttribute() bool {
	return false
}

func (a ListElem) isFunctionCall() bool {
	return false
}

func (a ListElem) isListElem() bool {
	return true
}

func (a *ListElem) Token() *token.Token{
	return a.tok
}

type Vars struct {
	variables 	[]*directories.VarEntry
	tok 		*token.Token
}

func (a Vars) Variables() []*directories.VarEntry {
	return a.variables
}

func (a Vars) isVars() bool {
	return true
}

func (a Vars) isAssign() bool {
	return false
}

func (a Vars) isCondition() bool {
	return true
}

func (a Vars) isWrite() bool {
	return true
}

func (a Vars) isReturn() bool {
	return false
}

func (a Vars) isFor() bool {
	return false
}

func (a Vars) isWhile() bool {
	return false
}

func (a Vars) isFunctionCall() bool {
	return false
}

func (a Vars) isPredefinedFunction() bool {
	return false
}

func (a *Vars) Token() *token.Token {
	return a.tok
}


type Assign struct {
	attr 	*Attribute
	exp 	*Expression
	tok 	*token.Token
}


func (a *Assign) Attribute() *Attribute {
	return a.attr
}

func (a *Assign) Expression() *Expression {
	return a.exp
}

func (a Assign) isVars() bool {
	return false
}

func (a Assign) isAssign() bool {
	return true
}

func (a Assign) isCondition() bool {
	return false
}

func (a Assign) isWrite() bool {
	return false
}

func (a Assign) isReturn() bool {
	return false
}

func (a Assign) isFor() bool {
	return false
}

func (a Assign) isWhile() bool {
	return false
}

func (a Assign) isFunctionCall() bool {
	return false
}

func (a Assign) isPredefinedFunction() bool {
	return false
}


func (a *Assign) Token() *token.Token {
	return a.tok
}


type Write struct {
	exp 	*Expression
	tok 	*token.Token
}

func (a *Write) Expression() *Expression {
	return a.exp
}

func (a Write) isVars() bool {
	return false
}

func (a Write) isAssign() bool {
	return false
}

func (a Write) isCondition() bool {
	return true
}

func (a Write) isWrite() bool {
	return true
}

func (a Write) isReturn() bool {
	return false
}

func (a Write) isFor() bool {
	return false
}

func (a Write) isWhile() bool {
	return false
}

func (a Write) isFunctionCall() bool {
	return false
}

func (a Write) isPredefinedFunction() bool {
	return false
}

func (a *Write) Token() *token.Token {
	return a.tok
}

type Condition struct {
	exp 		*Expression
	stmts  		[]Statement
	elseStmts 	[]Statement
	tok 		*token.Token
}

func (c *Condition) Expression() *Expression {
	return c.exp
}

func (c *Condition) Statements() []Statement {
	return c.stmts
}

func (c *Condition) ElseStatements() []Statement {
	return c.elseStmts
}

func (c *Condition) Token() *token.Token {
	return c.tok
}


func (s Condition) isVars() bool {
	return false
}

func (s Condition) isAssign() bool {
	return false
}

func (s Condition) isCondition() bool {
	return true
}

func (s Condition) isWrite() bool {
	return false
}

func (s Condition) isReturn() bool {
	return false
}

func (s Condition) isFor() bool {
	return false
}

func (s Condition) isWhile() bool {
	return false
}

func (s Condition) isFunctionCall() bool {
	return false
}

func (s Condition) isPredefinedFunction() bool {
	return false
}



type Return struct {
	exp *Expression
	tok *token.Token
}

func (r *Return) Expression() *Expression {
	return r.exp
}

func (r Return) isVars() bool {
	return false
}

func (r Return) isAssign() bool {
	return false
}

func (r Return) isCondition() bool {
	return false
}

func (r Return) isWrite() bool {
	return false
}

func (r Return) isReturn() bool {
	return true
}

func (r Return) isFor() bool {
	return false
}

func (r Return) isWhile() bool {
	return false
}

func (r Return) isFunctionCall() bool {
	return false
}

func (r Return) isPredefinedFunction() bool {
	return false
}

func (r *Return) Token() *token.Token {
	return r.tok
}

type For struct {
	init 	*Assign
	cond 	*Expression
	op 		*Assign
	blck 	[]Statement
	tok 	*token.Token
}

func (f *For) Init() *Assign {
	return f.init
}

func (f *For) Condition() *Expression {
	return f.cond
}

func (f *For) Operation() *Assign {
	return f.op
}

func (f *For) Block() []Statement {
	return f.blck
}

func (f For) isVars() bool {
	return false
}

func (f For) isAssign() bool {
	return false
}

func (f For) isCondition() bool {
	return false
}

func (f For) isWrite() bool {
	return false
}

func (f For) isReturn() bool {
	return false
}

func (f For) isFor() bool {
	return true
}

func (f For) isWhile() bool {
	return false
}

func (f For) isFunctionCall() bool {
	return false
}

func (f For) isPredefinedFunction() bool {
	return false
}

func (f *For) Token() *token.Token {
	return f.tok
}

type While struct {
	exp 	*Expression
	blck 	[]Statement
	tok 	*token.Token
}

func (w *While) Expression() *Expression {
	return w.exp
}

func (w *While) Block() []Statement {
	return w.blck
}


func (w *While) isVars() bool {
	return false
}


func (w *While) isAssign() bool {
	return false
}

func (w *While) isCondition() bool {
	return false
}

func (w *While) isWrite() bool {
	return false
}

func (w *While) isReturn() bool {
	return false
}

func (w *While) isFor() bool {
	return false
}

func (w *While) isWhile() bool {
	return true
}

func (w While) isFunctionCall() bool {
	return false
}


func (w While) isPredefinedFunction() bool {
	return false
}

func (w *While) Token() *token.Token {
	return w.tok
}

type FunctionCall struct {
	id 		string
	params 	[]*Expression
	tok 	*token.Token
}

func (fc *FunctionCall) Id() string {
	return fc.id
}

func (fc *FunctionCall) Params() []*Expression {
	return fc.params
}

func (fc FunctionCall) isVars() bool {
	return false
}


func (fc *FunctionCall) isAssign() bool {
	return false
}

func (fc *FunctionCall) isCondition() bool {
	return false
}

func (fc *FunctionCall) isWrite() bool {
	return false
}

func (fc *FunctionCall) isReturn() bool {
	return false
}

func (fc *FunctionCall) isFor() bool {
	return false
}

func (fc *FunctionCall) isWhile() bool {
	return false
}

func (fc *FunctionCall) isFunctionCall() bool {
	return true
}

func (fc *FunctionCall) isPredefinedFunction() bool {
	return false
}

func (fc *FunctionCall) isConstantValue() bool {
	return true
}

func (fc *FunctionCall) isAttribute() bool {
	return false
}

func (fc *FunctionCall) isListElem() bool {
	return false
}

func (fc *FunctionCall) Token() *token.Token {
	return fc.tok
}
/*
type ClassFunctionCall struct {
	id 		string
	obj 	string
	params 	[]*Expression
	tok 	*token.Token
}

func (fc *ClassFunctionCall) Id() string {
	return fc.id
}

func (fc *ClassFunctionCall) Obj() string {
	return fc.obj
}

func (fc *ClassFunctionCall) Params() []*Expression {
	return fc.params
}

func (fc *ClassFunctionCall) Token() *token.Token {
	return fc.tok
}
*/
// ConstantValue defines a type with a single basic value
type Constant interface {
	isConstantValue() 			bool
	isAttribute()				bool
	isFunctionCall()			bool
	isListElem()				bool
	Token()   					*token.Token
}

type ConstantValue struct {
	t     	*types.Type
	value 	string
	tok   	*token.Token
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

func (c ConstantValue) isConstantValue() bool {
	return true
}

func (c ConstantValue) isAttribute() bool {
	return false
}

func (c ConstantValue) isFunctionCall() bool {
	return false
}

func (c ConstantValue) isListElem() bool {
	return false
}


/*

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

type 


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
*/
