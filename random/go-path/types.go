package types

import "fmt"

type Foo struct {
	Counter uint64
}

func NewFoo(i uint64) Foo {
	return Foo{i}
}

func (f Foo) String() string {
	return fmt.Sprintf("%v", f.Counter)

}
