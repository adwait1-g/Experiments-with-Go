// Challenge1: Input from user
//
// Have the user enter a string, then loop through the string to generate a new string in which every character is duplicated, e.g., "hello" => "hheelllloo".
// Test with "世界" as input.
package main

import "fmt"

func main() {
	var input_string string
	var result_string string

	// Get the string
	fmt.Printf("Enter string: ")
	fmt.Scanf("%s", &input_string)

	for _, char := range input_string {
		result_string = result_string + string(char) + string(char)
	}

	fmt.Println("Result string:", result_string)
}
