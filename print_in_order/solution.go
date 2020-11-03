package print_in_order

import "github.com/CNife/leetcode-concurrency/run"

type Foo struct {
	run.Runner
	firstOk, secondOk chan struct{}
}

func NewFoo() *Foo {
	f := &Foo{
		firstOk:  make(chan struct{}, 1),
		secondOk: make(chan struct{}, 1),
	}
	f.Register(f.First, f.Second, f.Third)
	return f
}

func (f *Foo) First() {
	f.Print("first")
	f.firstOk <- struct{}{}
}

func (f *Foo) Second() {
	<-f.firstOk
	f.Print("second")
	f.secondOk <- struct{}{}
}

func (f *Foo) Third() {
	<-f.secondOk
	f.Print("third")
}
