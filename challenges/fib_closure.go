// Fibonacci closure

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	next := 0
	n1 := 0
	n2 := 1
	callCount := 0

	return func() int {
		var result int
		if callCount == 0 {
			result = n1
		} else if callCount == 1 {
			result = n2
		} else {
			next = n1 + n2
			n1 = n2
			n2 = next
			result = next
		}

		callCount++
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
