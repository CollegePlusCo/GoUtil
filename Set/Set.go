package set

type Set[T comparable] interface {
	Add(T) bool
	Remove(T) bool
	Contains(T) bool
	Intersection(*Set[T]) *[]T
	Consume(*[]T)
	ToArray() *[]T
	Size() int
}
