package linkedlist

import (
	"unsafe"
)

type Node[V any] struct {
	prev *Node[V]
	next *Node[V]
	data V
}

func (l *Node[V]) GetData() V {
	return l.data
}

type LinkedList[V any] struct {
	head *Node[V]
	tail *Node[V]
	len  int
	refs *map[uintptr]bool
	iter **Node[V]
}

func New[V any]() *LinkedList[V] {
	return &LinkedList[V]{refs: &map[uintptr]bool{}, len: 0}
}

func (l *LinkedList[V]) Has(value V, compareFunc func(a V, b V) bool) bool {
	curr := l.head
	for curr != nil {
		if compareFunc(curr.data, value) {
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *LinkedList[V]) Add(value V) {
	if l.head == nil {
		l.head = &Node[V]{data: value}
		l.tail = l.head
	} else {
		l.tail.next = &Node[V]{data: value, prev: l.tail}
		l.tail = l.tail.next
	}
	(*l.refs)[uintptr(unsafe.Pointer(&(*l.tail)))] = true
	l.len++
}

func (l *LinkedList[V]) AddWithNode(value V) *Node[V] {
	if l.head == nil {
		l.head = &Node[V]{data: value}
		l.tail = l.head
	} else {
		l.tail.next = &Node[V]{data: value, prev: l.tail}
		l.tail = l.tail.next
	}
	(*l.refs)[uintptr(unsafe.Pointer(&(*l.tail)))] = true
	l.len++
	return l.tail
}

func (l *LinkedList[V]) RemoveRef(node *Node[V]) (*Node[V], *Node[V]) {
	curr := node
	var prev *Node[V]
	var next *Node[V]
	prev = curr.prev
	next = curr.next
	if curr.prev != nil {
		if l.tail == curr {
			l.tail = curr.prev
		} else {
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
		}
	} else if curr.next != nil {
		l.head = curr.next
		curr.next.prev = nil
	} else {
		l.head = nil
		l.tail = nil
	}
	curr.prev = nil
	curr.next = nil
	if _, has := (*l.refs)[uintptr(unsafe.Pointer(&(*curr)))]; has {
		l.len--
		delete(*l.refs, uintptr(unsafe.Pointer(&(*curr))))
	}

	return prev, next
}

func (l *LinkedList[V]) ToArray() []V {
	values := make([]V, 0, 10)
	curr := l.head
	for curr != nil {
		values = append(values, curr.data)
		curr = curr.next
	}
	return values
}

func (l *LinkedList[V]) Len() int {
	return l.len
}

func (l *LinkedList[V]) StartIterator() {
	if l.head != nil {
		l.iter = &l.head
	}
}

func (l *LinkedList[V]) GetNext() V {
	var x V
	if l.head != nil {
		x = (*l.iter).data
		*l.iter = (*l.iter).next
	}
	return x
}

func (l *LinkedList[V]) Done() bool {
	return l.iter == nil || *l.iter == nil
}

func (l *LinkedList[V]) PopFirstWithNode() *Node[V] {
	if l.head != nil {
		x := l.head
		l.RemoveRef(l.head)
		return x
	}
	return nil
}

func (l *LinkedList[V]) PopLastWithNode() *Node[V] {
	if l.tail != nil {
		x := l.tail
		l.RemoveRef(l.tail)
		return x
	}
	return nil
}
