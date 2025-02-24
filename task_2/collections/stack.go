package collections

type Stack[T string | float64] struct {
	buffer []T
}

func (s *Stack[T]) Push(str T) {
	s.buffer = append(s.buffer, str)
}

func (s *Stack[T]) Pop() T {
	popedElement := s.buffer[len(s.buffer)-1]
	s.buffer = s.buffer[:len(s.buffer)-1]
	return popedElement
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Len() int {
	return len(s.buffer)
}

func (s *Stack[T]) Top() T {
	return s.buffer[s.Len()-1]
}
