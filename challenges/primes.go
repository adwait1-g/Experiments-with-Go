// 4. Prime Numbers

package main

import "fmt"

func main() {

	// Loop through
	for i := 2; i <= 25; i++ {
		fmt.Printf("%d: ", i)
		// Check if i is prime or not.
		// Method: Start from 2 and go till i/2.
		// If it has more than one factor, then it is
		// NOT prime.
		factor := 0
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				factor = j
				break
			}
		}

		// Check how many factors are there.
		if factor == 0 {
			fmt.Println("Prime Number")
		} else {
			fmt.Printf("%d X %d\n", factor, i/factor)
		}
	}
}
