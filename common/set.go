package common

type Set[T comparable] struct {
	data map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{data: map[T]bool{}}
}

func NewPointSet() Set[Point] {
	return NewSet[Point]()
}

func (s *Set[T]) Add(element T) {
	if !s.Contains(element) {
		s.data[element] = true
	}
}

func (s *Set[T]) Contains(element T) bool {
	_, ok := s.data[element]

	return ok
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Merge(other Set[T]) {
	for element := range other.data {
		s.Add(element)
	}
}
