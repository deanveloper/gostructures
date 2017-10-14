package linkedlist

import "errors"

var ErrNilNode = errors.New("Node cannot be nil")
var ErrNilList = errors.New("List cannot be nil")

// === NODE STUFF ===

// A node in a list. Each node is "linked" to the node that comes before and
// after it.
type Node struct {
	Prev  *Node
	Value interface{}
	Next  *Node
}

// Puts this node after the "other" node.
// Panics if either node is nil.
// This function does NOT replace the head/tail of a list.
func (n *Node) putAfter(other *Node) {
	if n == nil || other == nil {
		panic(ErrNilNode)
	}

	n.remove()

	if other.Next != nil {
		other.Next.Prev = n
	}

	other.Next = n
}

// Puts this node before the "other" node.
// Panics if either node is nil.
// This function does NOT replace the head/tail of a list.
func (n *Node) putBefore(other *Node) {
	if n == nil || other == nil {
		panic(ErrNilNode)
	}

	n.remove()

	if other.Prev != nil {
		other.Prev.Next = n
	}

	other.Prev = n
}

// Removes the node. Panics if nil.
// This function does NOT replace the head/tail of a list.
func (n *Node) remove() {
	if n == nil {
		panic(ErrNilNode)
	}

	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
}

// === LIST STUFF ===

// A Linked-List data structure. Linked Lists are made of Nodes, where each
// node is "linked" to the previous and next nodes in the list.
type List struct {
	Head *Node
	Tail *Node
	size int
}

// Returns the size of the list. The size is kept track of,
// so this operation works in constant time.
func (l *List) Size() int {
	return l.size
}

// Adds an element to the end of the list.
// Panics if the list is nil.
// Returns the node which represents the added element.
func (l *List) AddLast(elem interface{}) *Node {
	newNode := new(Node)
	newNode.Value = elem
	newNode.Prev = l.Tail

	if l.Tail != nil {
		l.Tail.Next = newNode
	}

	if l.Head == nil {
		l.Head = newNode
	}

	l.Tail = newNode

	return newNode
}

// Adds an element to the beginning of the list.
// Panics if the list is nil.
// Returns the node which represents the added element.
func (l *List) AddFirst(elem interface{}) *Node {
	newNode := new(Node)
	newNode.Value = elem
	newNode.Next = l.Head

	if l.Head != nil {
		l.Head.Next = newNode
	}

	if l.Tail == nil {
		l.Tail = newNode
	}

	l.Head = newNode

	return newNode
}

// Inserts/Moves a node after another one. Panics if list or nodes are nil.
func (l *List) InsertAfter(toInsert, target *Node) {
	if l == nil {
		panic(ErrNilList)
	}
	if toInsert == nil || target == nil {
		panic(ErrNilNode)
	}

	if l.Head == toInsert {
		l.Head = toInsert.Next
	}
	if l.Tail == toInsert {
		l.Tail = toInsert.Prev
	}

	toInsert.putAfter(target)
	if l.Tail == target {
		l.Tail = toInsert
	}
}

// Inserts/Moves a node before another one. Panics if list or nodes are nil.
func (l *List) InsertBefore(toInsert, target *Node) {
	if l == nil {
		panic(ErrNilList)
	}
	if toInsert == nil || target == nil {
		panic(ErrNilNode)
	}

	if l.Head == toInsert {
		l.Head = toInsert.Next
	}
	if l.Tail == toInsert {
		l.Tail = toInsert.Prev
	}

	toInsert.putAfter(target)

	if l.Head == target {
		l.Head = toInsert
	}
}

// Removes a given node. Panics if list or node is nil.
func (l *List) Remove(n *Node) {
	if l == nil {
		panic(ErrNilList)
	}
	if n == nil {
		panic(ErrNilNode)
	}

	if n == l.Head {
		l.Head = n.Next
	}
	if n == l.Tail {
		l.Tail = n.Prev
	}
	if n.Next != nil {
		n.Next.Prev = nil
	}
	if n.Prev != nil {
		n.Prev.Next = nil
	}
	n.Next = nil
	n.Prev = nil
	n.Value = nil
}

// Returns the list as a slice.
func (l *List) AsSlice() []interface{} {
	slice := make([]interface{}, l.Size())
	current := l.Head

	for i := 0; i < l.Size(); i++ {
		slice[i] = current.Value
		current = current.Next
	}

	return slice
}
