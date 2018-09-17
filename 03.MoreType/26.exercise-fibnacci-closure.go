package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	fmt.Printf("%d\n%d\n", a, b)
	return func() int {
		a, b = b, a+b
		fmt.Println(b)
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		f()
	}
}
