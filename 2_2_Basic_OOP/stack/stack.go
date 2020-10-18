package stack

type Stack struct {
	data []interface{}
}

func New() Stack {
	return Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	value := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return value
}

func (s Stack) Size() int {
	return len(s.data)
}

func (s Stack) GetAt(idx int) interface{} {
	return s.data[idx]
}
