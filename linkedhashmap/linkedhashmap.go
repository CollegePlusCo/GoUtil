package linkedhashmap

import (
	"encoding/json"
	"github.com/CollegePlusCo/GoUtil/linkedlist"
	"unsafe"
)

type LinkedHashMap[K comparable, V any] struct {
	keys      *linkedlist.LinkedList[K]
	m         *map[K]V
	keysIndex *map[K]*linkedlist.Node[K]
	cap       int
	maxSize   int
	currSize  int
}

func New[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{keys: linkedlist.New[K](), m: &map[K]V{}, keysIndex: &map[K]*linkedlist.Node[K]{}}
}

func (l *LinkedHashMap[K, V]) Init() {
	l.m = &map[K]V{}
	l.keysIndex = &map[K]*linkedlist.Node[K]{}
	l.keys = linkedlist.New[K]()
}

func (l *LinkedHashMap[K, V]) Has(key K) bool {
	if l.m == nil {
		l.Init()
	}
	_, has := (*l.m)[key]
	return has
}

func (l *LinkedHashMap[K, V]) Put(key K, value V) {
	if l.m == nil {
		l.Init()
	}
	node := l.keys.AddWithNode(key)
	if _, has := (*l.keysIndex)[key]; has {
		// constant time op
		l.keys.RemoveRef((*l.keysIndex)[key])
		l.currSize -= int(unsafe.Sizeof((*l.m)[key]))
	} else {
		if l.maxSize > 0 && l.currSize >= l.maxSize {
			first := l.keys.PopFirstWithNode()
			l.currSize -= int(unsafe.Sizeof((*l.m)[first.GetData()]))
			delete(*l.keysIndex, first.GetData())
			delete(*l.m, first.GetData())
		} else if l.cap > 0 && len(*l.m) == l.cap {
			first := l.keys.PopFirstWithNode()
			l.currSize -= int(unsafe.Sizeof((*l.m)[first.GetData()]))
			delete(*l.keysIndex, first.GetData())
			delete(*l.m, first.GetData())
		}
	}
	(*l.keysIndex)[key] = node
	(*l.m)[key] = value
	l.currSize += int(unsafe.Sizeof(value))
}

func (l *LinkedHashMap[K, V]) Remove(key K) bool {
	if l.m == nil {
		l.Init()
	}
	_, exists := (*l.m)[key]
	if exists {
		if (*l.keysIndex)[key] != nil {
			// constant time op
			l.keys.RemoveRef((*l.keysIndex)[key])
		}
		l.currSize -= int(unsafe.Sizeof((*l.m)[key]))
		delete(*l.m, key)
		delete(*l.keysIndex, key)
	}
	return exists
}

func (l *LinkedHashMap[K, V]) Get(key K) V {
	var x V
	if l.Has(key) {
		x = (*l.m)[key]
	}
	return x
}

func (l *LinkedHashMap[K, V]) Keys() []K {
	if l.m == nil {
		l.Init()
	}
	return l.keys.ToArray()
}

func (l *LinkedHashMap[K, V]) Values() []V {
	if l.m == nil {
		l.Init()
	}
	values := make([]V, 0, 10)
	for _, value := range *l.m {
		values = append(values, value)
	}
	return values
}

func (l *LinkedHashMap[K, V]) Len() int {
	if l.m == nil {
		l.Init()
	}
	return len(*l.m)
}

func (l *LinkedHashMap[K, V]) SetCap(cap int) {
	l.cap = cap
}

// max size of bytes
func (l *LinkedHashMap[K, V]) SetMaxSize(maxSize int) {
	l.maxSize = maxSize
}

func (l *LinkedHashMap[K, V]) StartIterator() {
	l.keys.StartIterator()
}

func (l *LinkedHashMap[K, V]) GetNext() K {
	return l.keys.GetNext()
}

func (l *LinkedHashMap[K, V]) Done() bool {
	return l.keys.Done()
}

func (l *LinkedHashMap[K, V]) ToJSON() []byte {
	marsh, marshErr := json.Marshal(l.m)
	if marshErr != nil {
		return nil
	}
	return marsh
}

func (l *LinkedHashMap[K, V]) Clear() {
	l.Init()
}
