package main

import (
	"fmt"
)

func main() {	
	foo()
	bar()
	fmt.Println("All Done.")
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}