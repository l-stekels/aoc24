package common

import "iter"

type void struct{}

var empty void

type Set[T comparable] struct {
	data map[T]void
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{data: map[T]void{}}
}

func NewPointSet() Set[Point] {
	return NewSet[Point]()
}

func (s *Set[T]) Add(element T) {
	s.data[element] = empty
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.data[element]

	return exists
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Merge(other Set[T]) {
	for element := range other.data {
		s.data[element] = empty
	}
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for element := range s.data {
			if !yield(element) {
				return
			}
		}
	}
}
