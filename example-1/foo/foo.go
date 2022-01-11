package foo

type Foo interface {
	Do(i int) int
}

func ProcessFoo(f Foo) {
	f.Do(12)
}
