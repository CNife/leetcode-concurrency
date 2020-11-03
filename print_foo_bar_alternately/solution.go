package print_foo_bar_alternately

import (
	"github.com/CNife/leetcode-concurrency/run"
)

type FooBar struct {
	run.Runner
	n            int
	fooOk, barOk chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:     n,
		fooOk: make(chan struct{}),
		barOk: make(chan struct{}),
	}
	fb.Register(fb.Foo, fb.Bar)
	return fb
}

func (fb *FooBar) Foo() {
	for i := 0; i < fb.n; i++ {
		<-fb.fooOk
		fb.Print("foo")
		fb.barOk <- struct{}{}
	}
}

func (fb *FooBar) Bar() {
	for i := 0; i < fb.n; i++ {
		fb.fooOk <- struct{}{}
		<-fb.barOk
		fb.Print("bar")
	}
}
