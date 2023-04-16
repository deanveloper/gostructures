package stack

import "fmt"

type Stack interface {
	Len() int
	Push(elem interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	fmt.Stringer
}
