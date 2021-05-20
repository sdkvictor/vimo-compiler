package types

import (
	"strings"
)

type BasicType int

const (
	Num BasicType = iota
	Char
	Bool
	Void
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
	if t == Void {
		return '4'
	}

	return 'n'
}

type ObjType int

const (
	Square ObjType = iota
	Circle
	Image
	Text
	Background
	NullObj
)

func (t ObjType) convert() rune {
	if t == Square {
		return '1'
	}

	if t == Circle {
		return '2'
	}

	if t == Image {
		return '3'
	}

	if t == Text {
		return '4'
	}

	if t == Background {
		return '5'
	}

	return 'n'
}


type Type struct {
	basic   	BasicType
	object		ObjType
	list    	int
	isObject	bool
}

// String converts the type to its string representation which is used only in the dirfunc package
// to build the composite key of an entry
func (l Type) String() string {
	var builder strings.Builder

	for i := 0; i < l.list; i++ {
		builder.WriteRune('[')
	}
	
	if l.isObject {
		builder.WriteRune(l.object.convert())

	} else {
		builder.WriteRune(l.basic.convert())
	}
	
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

// Type
func (lt *Type) Object() ObjType {
	return lt.object
}

// Object
func (lt *Type) IsObject() bool {
	return lt.isObject
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
	return &Type{b, NullObj, list, false}
}

// NewDataType Declares a new, object  type
func NewObjectType(o ObjType, list int) *Type {
	return &Type{Null, o, list, true}
}

