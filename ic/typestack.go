//Package ic provides the generation of intermediate code
package ic

import (
	"fmt"
	"strings"

	"github.com/sdkvictor/golang-compiler/types"
)

// nodet is a data container
type nodet struct {
	val  *types.Type
	next *nodet
}

// TypeStack implements a stack for the FuncEntry data type
type TypeStack struct {
	head *nodet
}

// Empty returns true if TypeStack is empty
func (s *TypeStack) Empty() bool {
	return s.head == nil
}

// String returns the string representation of the stack
func (s *TypeStack) String() string {
	if s.Empty() {
		return "<Empty Stack>"
	}

	curr := s.head
	var result strings.Builder

	result.WriteString("< ")

	for curr != nil {
		value := fmt.Sprintf("%v", curr.val)
		if curr.next == nil {
			result.WriteString(value)
		} else {
			result.WriteString(value + ", ")
		}
		curr = curr.next
	}

	result.WriteString(" >")

	return result.String()
}

// Pop removes the first element in the container
func (s *TypeStack) Pop() {
	if s.Empty() {
		return
	}

	s.head = s.head.next
}

// Top returns the first element in the container
func (s *TypeStack) Top() *types.Type {
	if s.Empty() {
		return nil
	}
	return s.head.val
}

// Push adds an element to the top of the container
func (s *TypeStack) Push(val *types.Type) {
	newHead := &nodet{val, s.head}
	s.head = newHead
}

// NewTypeStack ...
func NewTypeStack() *TypeStack {
	return &TypeStack{nil}
}