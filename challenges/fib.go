//2. Write a function to generate 'n' numbers from fibonacci series
// Print the first n numbers from the fibonacci series
package main

import "fmt"

func main() {
	var n int
	n1 := 0
	n2 := 1

	// Get the number
	fmt.Printf("Enter n: ")
	fmt.Scanf("%d", &n)
	if n <= 0 {
		fmt.Println("Invalid n. n must be positive")
		return
	}

	if n == 1 {
		fmt.Println(n1)
		return
	} else if n == 2 {
		fmt.Println(n1)
		fmt.Println(n2)
		return
	} else {
		fmt.Println(n1)
		fmt.Println(n2)
		for i := 3; i <= n; i++ {
			next := n1 + n2
			n2 = n1
			n1 = next
			fmt.Println(n1 + n2)
		}
	}
}
