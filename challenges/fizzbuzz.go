// Challenge 0: FizzBuzz challenge
// Print numbers 1 to 100:
//
//     for every number divisible by 3 print "Fizz" instead
//     for every number divisible by 5 print "Buzz" instead
//     for numbers divisible by 3 & 5 print "FizzBuzz"
//     for other numbers print them as is

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d", i)
		if i%15 == 0 {
			fmt.Printf(": FizzBuzz")
		} else if i%3 == 0 {
			fmt.Printf(": Fizz")
		} else if i%5 == 0 {
			fmt.Printf(": Buzz")
		}
		fmt.Printf("\n")
	}
}
