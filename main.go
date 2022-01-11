package main

import (
	"fmt"

	"github.com/shivamarora1/gomock/example-1/foo"
)

func ProcessFoo(f foo.Foo) {
	f.Do(12)
}

func main() {
	fmt.Printf("main program started...")
}
