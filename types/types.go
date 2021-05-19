package types

import (
	"strings"
)

type BasicType int

const (
	Num BasicType = iota
	Char
	Bool
	Square
	Circle
	Image
	Text
	Background
	Null
)

func (t BasicType) convert() rune {
	if t == Num {
		return '1'
	}

	if t == Char {
		return '2'
	}

	if t == Bool {
		return '3'
	}

	if t == Square {
		return '4'
	}

	if t == Circle {
		return '5'
	}

	if t == Image {
		return '6'
	}

	if t == Text {
		return '7'
	}

	if t == Background {
		return '8'
	}
	return 'n'
}

// Type represents any type on the lambdish language.
// A type in the language can be either a basic type (num, bool, char), or a function type.
//
// - In the case of the function type, the following should be set
//		-function: true
//		-params: non null (might be empty anyways)
//		-basic: NULL
//
// - In the case of a basic type, the following should be set
// 		-function: false
//		-params: null
//		-basic: non null, the corresponing type
//
// Additional to this, the type might represent a list. If that is the case, all rules set above will
// remain true, and list will be set to a non-zero value, indicating the levels of nesting of the type
// in the list.
//
// For example, a value of list = 1 for a basic type num will consists of a list as following
// 		- [num]
//
// And a value of list = 3 for a function type could then look like this
// 		- [[[(num, num => bool)]]]
//
// Nontheless external users of this package should only construct Type structs using the
// predefined constructors provided below. When a function type is needed, the NewFuncType
// function should be called, and NewDataType with the basic type accordingly.
// This will ensure that the values are initialized correctly according to the rules set above.
type Type struct {
	basic    BasicType
	list     int
}

// String converts the type to its string representation which is used only in the dirfunc package
// to build the composite key of an entry
func (l Type) String() string {
	var builder strings.Builder

	for i := 0; i < l.list; i++ {
		builder.WriteRune('[')
	}
	
	builder.WriteRune(l.basic.convert())
	
	for i := 0; i < l.list; i++ {
		builder.WriteRune(']')
	}
	return builder.String()
}

// List
func (lt *Type) List() int {
	return lt.list
}

// Type
func (lt *Type) Basic() BasicType {
	return lt.basic
}

// List
func (lt *Type) DecreaseList() {
	lt.list = lt.list - 1
}

// List
func (lt *Type) IncreaseList() {
	lt.list = lt.list + 1
}

//Equal
func (lt *Type) Equal(lt2 *Type) bool {
	return lt.String() == lt2.String()
}

// NewDataType Declares a new, basic  type
func NewDataType(b BasicType, list int) *Type {
	return &Type{b, list}
}

