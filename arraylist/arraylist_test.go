package arraylist_test

import (
	"reflect"
	"testing"

	. "github.com/deanveloper/gostructures/arraylist"
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
