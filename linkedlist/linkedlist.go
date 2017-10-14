package linkedlist

import (
	"bytes"
	"errors"
	"fmt"
)

var ErrSelfLinkNode = errors.New("attempted to self-link node")
var ErrNilNode = errors.New("node cannot be nil")
var ErrNilList = errors.New("list cannot be nil")

// === NODE STUFF ===

// A node in a list. Each node is "linked" to the node that comes before and
// after it.
type Node struct {
	Prev  *Node
	Value interface{}
	Next  *Node
}

// Puts this node after the "other" node.
// Panics if either node is nil, or if node is equal.
// This function does NOT replace the head/tail of a list.
func (n *Node) putAfter(other *Node) {
	if n == nil || other == nil {
		panic(ErrNilNode)
	}
	if n == other {
		panic(ErrSelfLinkNode)
	}

	n.remove()

	if other.Next != nil {
		other.Next.Prev = n
	}

	n.Prev = other
	n.Next = other.Next
	other.Next = n
}

// Puts this node before the "other" node.
// Panics if either node is nil, or if node is equal.
// This function does NOT replace the head/tail of a list.
func (n *Node) putBefore(other *Node) {
	if n == nil || other == nil {
		panic(ErrNilNode)
	}
	if n == other {
		panic(ErrSelfLinkNode)
	}

	n.remove()

	if other.Prev != nil {
		other.Prev.Next = n
	}

	n.Prev = other.Prev
	n.Next = other
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
	n.Next = nil
	n.Prev = nil
}

// Implement fmt.Stringer
func (n *Node) String() string {
	if n == nil {
		return "Node(nil)"
	}
	if s, ok := n.Value.(string); ok {
		return fmt.Sprintf("Node(%q)", s)
	}
	return fmt.Sprintf("Node(%v)", n.Value)
}

// === LIST STUFF ===

// A Linked-List data structure. Linked Lists are made of Nodes, where each
// node is "linked" to the previous and next nodes in the list.
type List struct {
	Head *Node
	Tail *Node
	Size int
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

	l.Size++

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
		l.Head.Prev = newNode
	}

	if l.Tail == nil {
		l.Tail = newNode
	}

	l.Head = newNode

	l.Size++

	return newNode
}

// Inserts/Moves a node after another one. Does NOT update the size if you insert a node
// from outside the list. Panics if list or nodes are nil, or if the nodes are equal.
func (l *List) InsertAfter(toInsert, target *Node) {
	if l == nil {
		panic(ErrNilList)
	}
	if toInsert == nil || target == nil {
		panic(ErrNilNode)
	}

	l.Remove(toInsert)

	toInsert.putAfter(target)
	l.Size++ // Increment as we removed the node and added it back in.

	if l.Tail == target {
		l.Tail = toInsert
	}
}

// Inserts/Moves a node before another one. Does NOT update the size if you insert a node
// from outside the list. Panics if list or nodes are nil, or if the nodes are equal.
func (l *List) InsertBefore(toInsert, target *Node) {
	if l == nil {
		panic(ErrNilList)
	}
	if toInsert == nil || target == nil {
		panic(ErrNilNode)
	}

	l.Remove(toInsert)

	toInsert.putBefore(target)
	l.Size++ // Increment as we removed the node and added it back in.
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
	n.remove()
	l.Size--
}

// Returns the list as a slice.
func (l *List) AsSlice() []interface{} {
	slice := make([]interface{}, l.Size)
	current := l.Head

	for i := 0; i < l.Size; i++ {
		slice[i] = current.Value
		current = current.Next
	}

	return slice
}

// Implement fmt.Stringer
func (l *List) String() string {
	if l == nil {
		return "{nil}"
	}

	counter := 0
	var buffer bytes.Buffer
	buffer.WriteString("{")
	current := l.Head
	for current != nil {
		if current != l.Head {
			buffer.WriteString(",")
		}
		buffer.WriteString(current.String())
		if counter++; counter > 25 {
			buffer.WriteString("...")
			break
		}
		current = current.Next
	}

	return buffer.String()
}
