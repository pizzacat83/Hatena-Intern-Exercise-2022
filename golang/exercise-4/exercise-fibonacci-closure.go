package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var a1, a2 int = 0, 1
	return func() int {
		defer func() { a1, a2 = a2, a1+a2 }()
		return a1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
