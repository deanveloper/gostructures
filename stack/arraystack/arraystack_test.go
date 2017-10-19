package arraystack_test

import (
	. "github.com/deanveloper/gostructures/stack/arraystack"
	"testing"
)

// Tests some stuff on an empty stack.
func TestEmpty(t *testing.T) {
	st := New()
	ExpectedActual(t, 0, st.Len())
	val, err := st.Peek()
	ExpectedActual(t, nil, val)
	ExpectedActual(t, ErrEmptyStack, err)
}

func TestStack(t *testing.T) {
	st := New()
	st.Push(1)
	st.Push("lol")
	stacksOnStacks := New()
	st.Push(stacksOnStacks)

	// "val" is used to extract the value from a peek/pop
	val, _ := st.Peek()
	ExpectedActual(t, stacksOnStacks, val)
	val, _ = st.Peek()
	ExpectedActual(t, stacksOnStacks, val)

	val, _ = st.Pop()
	ExpectedActual(t, stacksOnStacks, val)
	val, _ = st.Pop()
	ExpectedActual(t, "lol", val)

	st.Push("newVal")
	val, _ = st.Pop()
	ExpectedActual(t, "newVal", val)

	val, _ = st.Pop()
	ExpectedActual(t, 1, val)

	var err error
	val, err = st.Peek()
	ExpectedActual(t, nil, val)
	ExpectedActual(t, ErrEmptyStack, err)
}

func ExpectedActual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v")
	}
}
