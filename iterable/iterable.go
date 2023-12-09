package iterable

type Iterable[V any] interface {
	StartIterator()
	GetNext() V
	Done() bool
}
