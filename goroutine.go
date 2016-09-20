package main

import (
	"fmt"
)

func main() {
	go foo()
	go bar()
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
