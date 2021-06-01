package semantics

import (
	"fmt"
	"strings"

	"github.com/sdkvictor/golang-compiler/types"
)


var reservedFunctions = []string{
	"KeyPressed",
	"CheckCollision",
	"Pow",
	"Sqrt",
    "Render",
    "Clear",
	"Update",
}

var reservedFunctionsReturn = map[string]*types.Type{
	"KeyPressed": types.NewDataType(types.Bool, 0, 0),
	"CheckCollision": types.NewDataType(types.Bool, 0, 0),
	"Pow": types.NewDataType(types.Float, 0, 0),
	"Sqrt": types.NewDataType(types.Float, 0, 0),
    "Render": types.NewDataType(types.Void, 0, 0),
    "Clear": types.NewDataType(types.Void, 0, 0),
	"Update": types.NewDataType(types.Void, 0, 0),
}

// Param amount cannot be greater than 2 under current
// architecture
var reservedFunctionsParamAmount = map[string]int{
	"KeyPressed": 1,
	"CheckCollision": 2,
	"Pow": 2,
	"Sqrt": 1,
    "Render": 1,
	"Clear": 0,
	"Update": 0,
}

var ObjectAttributes = []string{
	"height",
	"width",
	"x",
	"y",
	"size",
	"color",
	"message",
	"image",
}

var ObjectAttributesTypes = map[string]*types.Type{
	"height": types.NewDataType(types.Float, 0, 0),
	"width": types.NewDataType(types.Float, 0, 0),
	"x": types.NewDataType(types.Float, 0, 0),
	"y": types.NewDataType(types.Float, 0, 0),
	"size": types.NewDataType(types.Float, 0, 0),
	"color": types.NewDataType(types.String, 0, 0),
	"message": types.NewDataType(types.String, 0, 0),
	"image": types.NewDataType(types.String, 0, 0),
}

var objectTypeWithAttribute = map[string][]string{
	"7": {"height", "width", "x", "y", "color"}, // Square
	"8": {"height", "width", "x", "y", "color"}, // Circle
	"9": {"height", "width", "x", "y", "image"}, // Image
	"a": {"x", "y", "color", "size", "message"}, // Text
}

var ObjectSize = len(ObjectAttributes) + 1

type Operation int

const (
	Add Operation = iota
	Sub
	Mult
	Div
	Lt
	Gt
	Assign
	Equal
	NotEqual
	And
	Or
	Not
)

func (o Operation) convert() string {
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
		return "=="
	case NotEqual:
		return "!="
	case Assign:
		return "="
	case And:
		return "&&"
	case Or:
		return "||"
	case Not:
		return "!"
		
	}

	return ""
}

// GetOperation ...
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
	case "==":
		return Equal
	case "!=":
		return NotEqual
	case "=":
		return Assign
	case "&&":
		return And
	case "||":
		return Or
	case "!":
		return Not
	}

	return Add
}

// SemanticCube represents the semantic cube as a map of a key to its result type
type SemanticCube struct {
	operations map[string]types.BasicType
}

// NewSemanticCube creates a new semantic cube struct
func NewSemanticCube() *SemanticCube {

	return &SemanticCube{
		map[string]types.BasicType{
			//Arithmetical Operators
			"+@44": types.Int,
			"-@44": types.Int,
			"/@44": types.Int,
			"*@44": types.Int,
			"+@11": types.Float,
			"-@11": types.Float,
			"/@11": types.Float,
			"*@11": types.Float,
			
			//Relational Operators
			"<@44": types.Bool,
			">@44": types.Bool,
			"<@11": types.Bool,
			">@11": types.Bool,
			"==@44": types.Bool,
			"==@11": types.Bool,
			"==@33": types.Bool,
			"==@22": types.Bool,
			"==@55": types.Bool,
			"!=@11": types.Bool,
			"!=@55": types.Bool,
			"!=@33": types.Bool,
			"!=@22": types.Bool,
			"!=@44": types.Bool,
			//Logical Operators
			"&&@33": types.Bool,
			"||@33":  types.Bool,
			"!@3":    types.Bool,
		},
	}
}

// Get takes a key a checks if it exists in the semantic cube. If it does, it returns the result type
func (c *SemanticCube) Get(key string) (types.BasicType, bool) {
	typ, ok := c.operations[key]
	if !ok {
		return types.Null, false
	}
	return typ, true
}

//isOperationFromSemanticCube
func isOperationFromSemanticCube(s string) bool {
	switch s {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	case "%":
		return true
	case "<":
		return true
	case ">":
		return true
	case "==":
		return true
	case "=":
		return true
	case "&&":
		return true
	case "||":
		return true
	case "!":
		return true
	}

	return false
}




func GetSemanticCubeKey(operation string, params []*types.Type) string {
	var b strings.Builder

	for _, p := range params {
		b.WriteString(p.String())
	}

	return fmt.Sprintf("%s@%s", operation, b.String())
}

func GetObjectAttributeType(objectAtt string) *types.Type {
	t, _ := ObjectAttributesTypes[objectAtt]
	return t
}

func matchTypeWithAttr(ot *types.Type, att string) bool {
	availAttributes, ok := objectTypeWithAttribute[ot.String()]
	if !ok {
		return false
	}

	for _, a := range availAttributes {
		if att == a {
			return true
		}
	}

	return false
}

func GetReservedReturnType(id string) *types.Type {
	t, _ := reservedFunctionsReturn[id]
	return t
}

func GetReservedParamAmount(id string) int {
	t, _ := reservedFunctionsParamAmount[id]
	return t
}