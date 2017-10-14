package linkedlist_test

import (
	. "github.com/deanveloper/gostructures/linkedlist"
	"testing"
)

// Tests AsSlice with an empty list.
func TestAsSlice_Empty(t *testing.T) {
	newList := new(List)
	slice := newList.AsSlice()
	expected := make([]interface{}, 0)

	if !EqualSlice(slice, expected) {
		t.Fatalf("Actual: %q\nExpected: %q\n", slice, []interface{}{})
	}
}

// Tests AsSlice with a single item in the list.
// Assumes AddLast works properly.
func TestAsSlice_Single(t *testing.T) {
	newList := new(List)
	newList.AddLast(100)

	slice := newList.AsSlice()
	expected := []interface{}{100}
	if !EqualSlice(slice, expected) {
		t.Fatalf("Actual: %q, Expected: %q\n", slice, expected)
	}
}

// Tests AddLast with an empty list
func TestAddLast_Single(t *testing.T) {
	newList := new(List)
	node := newList.AddLast(100)
	if newList.Head != node {
		t.Fatalf("Node is not head of list")
	}
	if newList.Tail != node {
		t.Fatalf("Node is not tail of list")
	}
	if node.Value != 100 {
		t.Fatalf("Node value is incorrect")
	}
	if node.Prev != nil {
		t.Fatalf("Node should not have Prev")
	}
	if node.Next != nil {
		t.Fatalf("Node should not have Next")
	}
}

// Tests AddFirst with an emtpy list
func TestAddFirst_Single(t *testing.T) {
	newList := new(List)
	node := newList.AddFirst(100)
	if newList.Head != node {
		t.Fatalf("Node is not head of list")
	}
	if newList.Tail != node {
		t.Fatalf("Node is not tail of list")
	}
	if node.Value != 100 {
		t.Fatalf("Node value is incorrect")
	}
	if node.Prev != nil {
		t.Fatalf("Node should not have Prev")
	}
	if node.Next != nil {
		t.Fatalf("Node should not have Next")
	}
}

// Tests AddLast with 2 items
func TestAddLast_Double(t *testing.T) {
	newList := new(List)
	first := newList.AddLast(50)
	last := newList.AddLast(23)

	// test state
	if newList.Head != first {
		t.Errorf("Head should be %q, Actually %q\n", newList.Head, first)
	}
	if newList.Tail != last {
		t.Errorf("Tail should be %q, Actually %q\n", newList.Tail, last)
	}
	if first.Prev != nil {
		t.Errorf("first.Prev should be nil, actually %q\n", first.Prev)
	}
	if first.Next != last {
		t.Errorf("first.Next should be %q, actually %q\n", last, first.Next)
	}
	if last.Prev != first {
		t.Errorf("last.Prev should be %q, actually %q\n", first, last.Prev)
	}
	if last.Next != nil {
		t.Errorf("last.Next should be nil, actually %q\n", last.Next)
	}

}

// Tests AddFirst with 2 items
func TestAddFirst_Double(t *testing.T) {
	newList := new(List)
	last := newList.AddFirst(50)
	first := newList.AddFirst(23)

	// test state
	if newList.Head != first {
		t.Errorf("Head should be %q, Actually %q\n", newList.Head, first)
	}
	if newList.Tail != last {
		t.Errorf("Tail should be %q, Actually %q\n", newList.Tail, last)
	}
	if first.Prev != nil {
		t.Errorf("first.Prev should be nil, actually %q\n", first.Prev)
	}
	if first.Next != last {
		t.Errorf("first.Next should be %q, actually %q\n", last, first.Next)
	}
	if last.Prev != first {
		t.Errorf("last.Prev should be %q, actually %q\n", first, last.Prev)
	}
	if last.Next != nil {
		t.Errorf("last.Next should be nil, actually %q\n", last.Next)
	}
}

// Done with simple tests, let's do some complex testing!
func TestLarge(t *testing.T) {
	newList := new(List)
	node2 := newList.AddFirst(2)
	node1 := newList.AddFirst(1)
	node3 := newList.AddLast("3")
	temp := "4"
	ptr := &temp
	node4 := newList.AddLast(ptr)

	// test AsSlice with large list
	slice := newList.AsSlice()
	expected := []interface{}{1, 2, "3", ptr}
	if !EqualSlice(slice, expected) {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, newList.AsSlice())
	}

	// Test InsertAfter
	newList.InsertAfter(node2, node3)
	expected = []interface{}{1, "3", 2, ptr}
	if !EqualSlice(newList.AsSlice(), expected) {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, newList.AsSlice())
	}

	// Test InsertBefore, inserting at head
	newList.InsertBefore(node2, node1)
	expected = []interface{}{2, 1, "3", ptr}
	if !EqualSlice(newList.AsSlice(), expected) {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, newList.AsSlice())
	}

	// Test InsertAfter, inserting at tail
	newList.InsertAfter(node3, node4)
	expected = []interface{}{2, 1, ptr, "3"}
	if !EqualSlice(newList.AsSlice(), expected) {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, newList.AsSlice())
	}

	// Test Remove, removing head
	newList.Remove(newList.Head)
	expected = []interface{}{1, ptr, "3"}
	if !EqualSlice(newList.AsSlice(), expected) {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, newList.AsSlice())
	}

	// Test remove, removing middle
	newList.Remove(node3)
	expected = []interface{}{1, ptr}
	if !EqualSlice(newList.AsSlice(), expected) {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, newList.AsSlice())
	}

	// Test remove, removing tail
	newList.Remove(newList.Tail)
	if newList.Head != node1 || newList.Tail != node1 {
		t.Fatalf("Expected Head/Tail: %q, Actual Head: %q, Actual Tail: %q", node1, newList.Head, newList.Tail)
	}
	if node1.Next != nil || node1.Prev != nil {
		t.Fatalf("node1 Next/Prev should be nil, instead Prev=%q Next=%q", node1.Prev, node1.Next)
	}

	// Test removing final elem
	newList.Remove(node1)
	if newList.Head != nil || newList.Tail != nil {
		t.Fatalf("list should be empty, instead Head=%q, Tail=%q", newList.Head, newList.Tail)
	}
}

func EqualSlice(arr1, arr2 []interface{}) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, e := range arr1 {
		if e != arr2[i] {
			return false
		}
	}

	return true
}
