// A package that provides an implementation of an ArrayList.
package arraylist

import (
	"errors"
	"fmt"
)

var ErrOutOfBounds error = errors.New("index out of bounds")

// Initialize with new(List) or using any of the New*() functions
type List struct {
	arr []interface{}
}

// === CONSTRUCTORS ===

// Returns an arraylist with initial capacity 16
func New() *List {
	return NewWithCap(16)
}

// Returns an arraylist with any capacity that you wish.
// Panics if capacity is negative.
func NewWithCap(capacity int) *List {
	if capacity < 0 {
		panic(ErrOutOfBounds)
	}
	newList := new(List)
	newList.arr = make([]interface{}, 0, capacity)
	return newList
}

// Gets the length of the arraylist.
func (l *List) Len() int {
	return len(l.arr)
}

// Gets the value at an index in the arraylist.
//
// Returns the value at the given index, or ErrOutOfBounds if
// the given index is < 0, or >= l.Len()
func (l *List) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.Len() {
		return nil, ErrOutOfBounds
	}

	return l.arr[index], nil
}

// Gets the value at the given index in the arraylist.
//
// Returns the value at the given index, and panics if the
// given index is < 0, or >= l.Len().
//
// The advantage of this method is that it only returns
// one value, allowing for the method to be inlined in
// things like if-statements.
func (l *List) GetForce(index int) interface{} {
	return l.arr[index]
}

// Add an element to the end of the list.
func (l *List) AddLast(elem interface{}) {
	l.arr = append(l.arr, elem)
}

// Add an element to the beginning of the list.
func (l *List) AddFirst(elem interface{}) {
	l.AddAt(0, elem)
}

// Add an element at a specific index. Returns ErrOutOfBounds if index < 0 or >= l.Len().
func (l *List) AddAt(index int, elem interface{}) error {
	if index < 0 || index >= l.Len() {
		return ErrOutOfBounds
	}

	l.arr = append(l.arr, nil)
	l.shift(index, 1)

	l.arr[index] = elem

	return nil
}

// Add each element of a slice at the end of the list.
// Uses copy() for maximum efficiency.
//
// To convert your array to []interface, see https://stackoverflow.com/a/27689178/3396646
func (l *List) AddSliceLast(elem []interface{}) {
	nils := make([]interface{}, len(elem))
	l.arr = append(l.arr, nils...)

	copy(l.arr[len(l.arr)-len(elem):], elem)
}

// Add each element of a slice at the beginning of the list.
// Uses copy() for maximum efficiency.
//
// To convert your array to []interface, see https://stackoverflow.com/a/27689178/3396646
func (l *List) AddSliceFirst(elem []interface{}) {
	l.AddSliceAt(0, elem)
}

func (l *List) AddSliceAt(index int, elems []interface{}) error {
	if index < 0 || index >= l.Len() {
		return ErrOutOfBounds
	}

	nils := make([]interface{}, len(elems))
	l.arr = append(l.arr, nils...)

	l.shift(index, len(elems))
	copy(l.arr[index:], elems)

	return nil
}

// Removes the value at an index.
// Returns (value, error) where value is the removed
// value, and error is ErrOutOfBounds if the given index < 0 or >= l.Len()
func (l *List) Remove(index int) (interface{}, error) {
	if index < 0 || index >= l.Len() {
		return nil, ErrOutOfBounds
	}

	value := l.arr[index]
	l.shift(index+1, -1)
	l.arr = l.arr[:len(l.arr)-1]

	return value, nil
}

// Copies the list into a slice.
func (l *List) AsSlice() []interface{} {
	result := make([]interface{}, l.Len())
	copy(result, l.arr)
	return result
}

// implement fmt.Stringer
func (l *List) String() string {
	return fmt.Sprint(l.arr)
}

// === PRIVATE METHODS ===

// Copies the array into itself with an offset of shamt.
// Uses the copy() builtin to do so. This method assumes that
// l.arr is large enough.
func (l *List) shift(start, shamt int) {
	if shamt == 0 {
		return
	}

	if shamt > 0 {
		copy(l.arr[start+shamt:], l.arr[start:len(l.arr)-shamt+1])
	} else {
		copy(l.arr[start+shamt:], l.arr[start:len(l.arr)])
	}
}
