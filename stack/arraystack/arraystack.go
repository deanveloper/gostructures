package arraystack

var ErrEmptyStack error = errors.New("the stack is empty")

type Stack struct {
	list *arraylist.List
}

// Creates a new stack with an initial capacity of 16.
func New() *Stack {
	return NewWithCap(16)
}

// Creates a new stack with an initial capacity.
func NewWithCap(capacity int) *Stack {
	stack := new(Stack)
	stack.list = arraylist.NewWithCap(capacity)
	return stack
}

// === PUBLIC STACK METHODS ===

// Returns the number of elements in the stack.
func (s *Stack) Len() int {
	return s.list.Len()
}

// Pushes an element to the top of the stack.
func (s *Stack) Push(elem interface{}) {
	s.list.AddLast(elem)
}

// Pops the element off of the top of the stack.
// This means that the element is first removed
// from the stack, and then returned. Will return
// ErrEmptyStack if the stack is empty.
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	elem, _ := s.list.Remove(s.Len() - 1)
	return elem, nil
}

// Peeks at the top element of the stack. This returns
// the top element, but without removing it. Returns
// ErrEmptyStack if the stack is empty.
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	return s.list.GetForce(s.Len() - 1), nil
}

// Returns if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

// implement fmt.Stringer
func (s *Stack) String() string {
	return fmt.Sprint(s.list.AsSlice)
}
