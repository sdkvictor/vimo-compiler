package quad

import (
	"fmt"

	"github.com/sdkvictor/golang-compiler/mem"
)

// Operation The operation or instruction of the system is a number
type Operation int

const (
	Add Operation = iota
	Sub
	Mult
	Div
	Lt
	Gt
	Equal
	Assign
	And
	Or
	Not
	GotoT
	GotoF
	Goto
	Ret
	Era
	Param
	Call
	CheckBound
	AddAddr
	AssignIndex // This assigns a value to a given index
	AssignIndexInv // This assigns an index to a given value
	Init
		
	Render 
    
    KeyPressed
    Print
    CheckCollision
    Pow
    Sqrt
	Clear
	Update

	Invalid
)

func (o Operation) String() string {
	switch o {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mult:
		return "*"
	case Div:
		return "/"
	case Lt:
		return "<"
	case Gt:
		return ">"
	case Equal:
		return "Equal"
	case And:
		return "And"
	case Or:
		return "Or"
	case Not:
		return "!"
	case GotoT:
		return "GotoT"
	case GotoF:
		return "GotoF"
	case Goto:
		return "Goto"
	case Ret:
		return "Ret"
	case Call:
		return "Call"
	case Era:
		return "Era"
	case Param:
		return "Param"
	case Print:
		return "Print"
	case Assign:
		return "Assign"
	case CheckBound:
		return "CheckBound"
	case AddAddr:
		return "AddAddr"
	case AssignIndex:
		return "AssignIndex"
	case AssignIndexInv:
		return "AssignIndexInv"
	case Init:
		return "Init"
	case Render:
		return "Render"
	case KeyPressed:
		return "KeyPressed"
	case CheckCollision:
		return "CheckCollision"
	case Pow:
		return "Pow"
	case Sqrt:
		return "Sqrt"
	case Clear:
		return "Clear"
	case Update:
		return "Update"
	}

	return ""
}

func GetOperation(s string) Operation {
	switch s {
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mult
	case "/":
		return Div
	case "<":
		return Lt
	case ">":
		return Gt
	case "Equal":
		return Equal
	case "And":
		return And
	case "Or":
		return Or
	case "!":
		return Not
	case "GotoT":
		return GotoT
	case "GotoF":
		return GotoF
	case "Goto":
		return Goto
	case "Ret":
		return Ret
	case "Call":
		return Call
	case "Era":
		return Era
	case "Param":
		return Param
	case "Print":
		return Print
	case "Assign":
		return Assign
	case "Render":
		return Render
	case "CheckBound":
		return CheckBound
	case "AddAddr":
		return AddAddr
	case "AssignIndex":
		return AssignIndex
	case "AssignIndexInv":
		return AssignIndexInv
	case "Init":
		return Init
	case "KeyPressed":
		return KeyPressed
	case "CheckCollision":
		return CheckCollision
	case "Pow":
		return Pow
	case "Sqrt":
		return Sqrt
	case "Clear":
		return Clear
	case "Update":
		return Update
	}

	return Invalid
}

func StringToOperation(s string) Operation {
	switch s {
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mult
	case "/":
		return Div
	case "<":
		return Lt
	case ">":
		return Gt
	case "Equal":
		return Equal
	case "And":
		return And
	case "Or":
		return Or
	case "!":
		return Not
	case "GotoT":
		return GotoT
	case "GotoF":
		return GotoF
	case "Goto":
		return Goto
	case "Ret":
		return Ret
	case "Call":
		return Call
	case "Era":
		return Era
	case "Param":
		return Param
	case "Print":
		return Print
	case "Assign":
		return Assign
	case "Render":
		return Render
	case "KeyPressed":
		return KeyPressed
	case "CheckCollision":
		return CheckCollision
	case "Pow":
		return Pow
	case "Sqrt":
		return Sqrt
	case "Clear":
		return Clear
	case "Update":
		return Update
	}

	return Invalid
}

/*
Exceptions:
– Unary operators: no arg2
– Operators like param: no arg2, no result
– (Un)conditional jumps: target label is the result
*/
/*
	Structure: Defines the main structure of code generation
	It contains the operation, 2 operands that holds the addresses and the third one
	is usually the result of the operation.
*/
type Quadruple struct {
	op Operation
	a1 mem.Address
	a2 mem.Address
	r  mem.Address
}

func (q *Quadruple) SetR(addr mem.Address) {
	q.r = addr
}

func (q *Quadruple) SetLop(addr mem.Address) {
	q.a1 = addr
}

func (q *Quadruple) SetRop(addr mem.Address) {
	q.a2 = addr
}

func (q *Quadruple) Op() Operation {
	return q.op
}

func (q *Quadruple) Lop() mem.Address {
	return q.a1
}

func (q *Quadruple) Rop() mem.Address {
	return q.a2
}

func (q *Quadruple) R() mem.Address {
	return q.r
}

//String is used in the creation of the file object that has the quadruples
func (q Quadruple) String() string {
	return fmt.Sprintf("%s %s %s %s", q.op, q.a1, q.a2, q.r)
}

//NewQuadruple Creates new quadruple with the addresses given and the number of operation
func NewQuadruple(op Operation, a1, a2, r mem.Address) *Quadruple {
	return &Quadruple{op, a1, a2, r}
}