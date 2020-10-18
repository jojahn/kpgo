package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	s.Push(1)
	res := s.Pop()
	if res.(int) != 1 {
		t.Error("Stack did not return on put")
	}

	res = s.Pop()
	if res != nil {
		t.Error("Stack did not actually pop")
	}

	s.Push("Hello, World!")
	res = s.Pop()
	if res != "Hello, World!" {
		t.Error("Stack cannot use strings")
	}

	size := s.Size()
	if size != 0 {
		t.Error("Size must 0 if empty")
	}

	s.Push(1)
	s.Push("Hello, World!")
	size = s.Size()
	if size != 2 {
		t.Error("Size did not evaluate to the right amount")
	}

	if s.GetAt(0).(int) != 1 || s.GetAt(1).(string) != "Hello, World!" {
		t.Error("Failed to get correct item")
	}
}