// A package that provides an implementation of an ArrayList.
package arraylist

import "errors"

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

// Returns an arraylist with any capacity that you wish
func NewWithCap(capacity uint) *List {
	newList := new(List)
	newList.arr = make([]interface{}, 0, capacity)
	return newList
}

// Gets the length of the arraylist.
func (l *List) Len() uint {
	return uint(len(l.arr))
}

// Add an element to the end of the list.
func (l *List) AddLast(elem interface{}) {
	l.arr = append(l.arr, elem)
}

// Add an element to the beginning of the list.
func (l *List) AddFirst(elem interface{}) {
	l.AddAt(0, elem)
}

// Add an element at a specific index. Returns ErrOutOfBounds if index is too large.
func (l *List) AddAt(index uint, elem interface{}) error {
	if index > l.Len() {
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

func (l *List) AddSliceAt(index uint, elems []interface{}) error {
	if index > l.Len() {
		return ErrOutOfBounds
	}

	nils := make([]interface{}, len(elems))
	l.arr = append(l.arr, nils...)

	l.shift(index, len(elems))
	copy(l.arr[index:], elems)

	return nil
}

// === PRIVATE METHODS ===

// Copies the array into itself with an offset of shamt.
// Uses the copy() builtin to do so. This method assumes that
// l.arr is large enough.
func (l *List) shift(start uint, shamt int) {
	if shamt == 0 {
		return
	}

	iStart := int(start)

	if shamt > 0 {
		copy(l.arr[iStart+shamt:], l.arr[start:len(l.arr)-shamt+1])
	} else {
		copy(l.arr[iStart+shamt:], l.arr[start:len(l.arr)])
	}
}
