package stack

// Stack
type Stack []string

// Empty
func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s Stack) Clear() {
	for !s.Empty() {
		s.Pop()
	}
}

// Push
func (s Stack) Push(v string) Stack {
	return append(s,v)
}

// Pop
func (s Stack) Pop() (Stack, bool) {
	if s.Empty() {
		return s, false
	}

	length := len(s)
	return s[:length-1], true
}

// Top
func (s Stack) Top() (string, bool) {
	if s.Empty() {
		return "", false
	}

	length := len(s)
	return s[length-1], true
}

func (s Stack) Size() (int) {
	return len(s);
}