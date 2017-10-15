package arraylist_test

import (
	"reflect"
	"testing"

	. "github.com/deanveloper/gostructures/list/arraylist"
)

func TestEmpty(t *testing.T) {
	list := New()

	if list.Len() != 0 {
		t.Errorf("Len(): Expected %v, Actual %v", 0, list.Len())
	}

	slice := list.AsSlice()

	if !reflect.DeepEqual([]interface{}{}, slice) {
		t.Errorf("AsSlice(): Expected %v, Actual %v", []interface{}{}, slice)
	}
}

func TestSingle(t *testing.T) {
	list := New()
	list.AddLast("Hello World!")

	if list.Len() != 1 {
		t.Errorf("Len(): Expected %v, Actual %v", 1, list.Len())
	}

	if list.GetForce(0) != "Hello World!" {
		t.Errorf("GetForce(0): Expected %v, Actual %v", "Hello World!", list.GetForce(0))
	}

	expectedSlice := []interface{}{"Hello World!"}
	if !reflect.DeepEqual(expectedSlice, list.AsSlice()) {
		t.Errorf("AsSlice(): Expected %v, Actual %v", expectedSlice, list.AsSlice())
	}
}

func TestMega(t *testing.T) {
	list := New()
	list.AddLast("Last")
	list.Remove(0)

	val, err := list.Get(0)
	if err != ErrOutOfBounds || val != nil {
		t.Fatalf("Remove -> Get: Expected (val %v, err %v), Actual (val %v, err %v)", nil, ErrOutOfBounds, val, err)
	}

	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	list.AddLast("four")
	list.AddLast("FiVe")
	list.AddFirst("zero")

	if list.Len() != 6 {
		t.Errorf("Len(): Expected %d, Actual %d", 6, list.Len())
	}

	expectedSlice := []interface{}{"zero", 1, 2, 3, "four", "FiVe"}
	if !reflect.DeepEqual(list.AsSlice(), expectedSlice) {
		t.Errorf("Slices: Expected %v, Actual %v", expectedSlice, list.AsSlice())
	}

	list.Remove(2)
	expectedSlice2 := []interface{}{"zero", 1, 3, "four", "FiVe"}
	if !reflect.DeepEqual(list.AsSlice(), expectedSlice2) {
		t.Errorf("Slices: Expected %v, Actual %v", expectedSlice2, list.AsSlice())
	}

	list.Remove(0)
	expectedSlice3 := []interface{}{1, 3, "four", "FiVe"}
	if !reflect.DeepEqual(list.AsSlice(), expectedSlice3) {
		t.Errorf("Slices: Expected %v, Actual %v", expectedSlice3, list.AsSlice())
	}

	list.Remove(3)
	expectedSlice4 := []interface{}{1, 3, "four"}
	if !reflect.DeepEqual(list.AsSlice(), expectedSlice4) {
		t.Errorf("Slices: Expected %v, Actual %v", expectedSlice4, list.AsSlice())
	}

	list.AddFirst(nil)
	list.AddFirst(10)
	expectedSlice5 := []interface{}{10, nil, 1, 3, "four"}
	if !reflect.DeepEqual(list.AsSlice(), expectedSlice5) {
		t.Errorf("Slices: Expected %v, Actual %v", expectedSlice5, list.AsSlice())
	}
	val2, err2 := list.Get(1)
	if val2 != nil || err2 != nil {
		t.Errorf("Get nil: Expected (val %v, err %v), Actual (val %v, err %v)")
	}
}
