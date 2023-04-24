package hashset

import (
	set "github.com/senseisub/collegepluscollegefinderbackend-backend-api/util/Set"
)

type HashSet[T comparable] struct {
	m map[T]bool // TODO: make our own hash structure later
}

func (pset *HashSet[T]) Add(elem T) bool {
	if _, contains := pset.m[elem]; contains {
		return true
	}
	pset.m[elem] = true
	return false
}

func (pset *HashSet[T]) Remove(elem T) bool {
	if _, contains := pset.m[elem]; contains {
		delete(pset.m, elem)
		return true
	}
	return false
}

func (pset *HashSet[T]) Contains(elem T) bool {
	_, contains := pset.m[elem]
	return contains
}

func (pset *HashSet[T]) Intersection(set2 *set.Set[T]) *[]T {
	returningArr := make([]T, 0, len(pset.m))
	for key := range pset.m {
		if (*set2).Contains(key) {
			returningArr = append(returningArr, key)
		}
	}
	return &returningArr
}

func (pset *HashSet[T]) Consume(incoming *[]T) {
	for _, key := range *incoming {
		pset.m[key] = true
	}
}

func (pset *HashSet[T]) ToArray() *[]T {
	returningArr := make([]T, 0, len(pset.m))
	for key := range pset.m {
		returningArr = append(returningArr, key)
	}
	return &returningArr
}

func (pset *HashSet[T]) Size() int {
	return len(pset.m)
}

func New[T comparable]() set.Set[T] {
	return &HashSet[T]{m: make(map[T]bool)}
}

func NewWithInput[T comparable](incoming *[]T) set.Set[T] {
	pset := HashSet[T]{m: make(map[T]bool)}
	pset.Consume(incoming)
	return &pset
}
