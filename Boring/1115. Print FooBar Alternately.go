package boring

type FooBar struct {
	n       int
	fooChan chan struct{}
	barChan chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{n: n,
		fooChan: make(chan struct{}, 1),
		barChan: make(chan struct{}),
	}
	fb.fooChan <- struct{}{}
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooChan
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.barChan <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barChan
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.fooChan <- struct{}{}
	}
}
