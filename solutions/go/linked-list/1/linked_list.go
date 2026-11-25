// Reference solution for the Exercism Go linked-list exercise.
// Write your own linked_list.go first; compare after.
//
// Run:  go test ./solutions/ -v
//
// Design note vs your Python version: Python used sentinel head/tail nodes.
// Here the tests require First().Prev() == nil and Last().Next() == nil, so
// sentinels would leak through Next()/Prev(). Plain nil-terminated ends are
// simpler: the only special cases are "list was empty" and "list becomes empty".
package linkedlist

import "errors"

// ErrEmptyList is returned by Shift and Pop on an empty list.
var ErrEmptyList = errors.New("empty list")

// Node is one element. Value is exported because the tests read it.
type Node struct {
	Value any
	next  *Node
	prev  *Node
}

// List is a doubly linked list. The zero value List{} is a valid empty list.
type List struct {
	head  *Node
	tail  *Node
	count int
}

// NewList builds a list from elements, in order.
func NewList(elements ...any) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

// Next returns the following node, or nil at the end.
func (n *Node) Next() *Node { return n.next }

// Prev returns the preceding node, or nil at the start.
func (n *Node) Prev() *Node { return n.prev }

// Unshift adds v at the front.
func (l *List) Unshift(v any) {
	n := &Node{Value: v, next: l.head}
	if l.head == nil {
		l.tail = n // was empty: n is also the last node
	} else {
		l.head.prev = n
	}
	l.head = n
	l.count++
}

// Push adds v at the back.
func (l *List) Push(v any) {
	n := &Node{Value: v, prev: l.tail}
	if l.tail == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
	l.count++
}

// Shift removes and returns the first value.
func (l *List) Shift() (any, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}
	n := l.head
	l.unlink(n)
	return n.Value, nil
}

// Pop removes and returns the last value.
func (l *List) Pop() (any, error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}
	n := l.tail
	l.unlink(n)
	return n.Value, nil
}

// Reverse reverses the list in place: swap every node's links, then swap the ends.
func (l *List) Reverse() {
	for n := l.head; n != nil; n = n.prev { // prev is the old next after the swap
		n.next, n.prev = n.prev, n.next
	}
	l.head, l.tail = l.tail, l.head
}

// First returns the first node, or nil if the list is empty.
func (l *List) First() *Node { return l.head }

// Last returns the last node, or nil if the list is empty.
func (l *List) Last() *Node { return l.tail }

// Count returns the number of nodes. O(1): kept up to date on every change.
func (l *List) Count() int { return l.count }

// Delete removes the first node whose Value == v. Reports whether one was removed.
func (l *List) Delete(v any) bool {
	for n := l.head; n != nil; n = n.next {
		if n.Value == v {
			l.unlink(n)
			return true
		}
	}
	return false
}

// unlink removes n from the list. Every removal goes through here, so the
// four "is n at an end?" cases are written once.
func (l *List) unlink(n *Node) {
	if n.prev == nil {
		l.head = n.next
	} else {
		n.prev.next = n.next
	}
	if n.next == nil {
		l.tail = n.prev
	} else {
		n.next.prev = n.prev
	}
	n.next, n.prev = nil, nil // a removed node must not still point into the list
	l.count--
}