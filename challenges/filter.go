// Implement your own filter function
// A filter function will take a slice & a predicate which evaluates to bool. The filter fn should return a slice back of same type.
//
// Eg. Take a slice of string: []string{"Iron Man", "Batman", "Superman", "Spider-man", "Wonder Woman", "Iron Fist", "Daredevil", "Supergirl", "Flash"}.
// func filter(s []string, predicate func(string) bool) []string {
// }

package main

import (
	"fmt"
	"strings"
)

func filter(s []string, predicate func(string) bool) []string {

	var result []string

	for _, str := range s {
		if predicate(str) == true {
			result = append(result, str)
		}
	}

	return result
}

func main() {
	slice := []string{"Iron Man", "Batman", "Superman", "Spider-man", "Wonder Woman", "Iron Fist", "Daredevil", "Supergirl", "Flash"}

	manNotPresent := func(str string) bool {
		return !strings.Contains(str, "man") && !strings.Contains(str, "Man")
	}

	secondCharNotIsVowel := func(str string) bool {
		if str[1] == 'a' || str[1] == 'e' || str[1] == 'i' || str[1] == 'o' || str[1] == 'u' {
			return true
		}

		return false
	}

	fmt.Println("All strings with no man/Man in it: ")
	filter_strs := filter(slice, manNotPresent)
	for _, str := range filter_strs {
		fmt.Println(str)
	}

	fmt.Println("All strings with vowel as second character: ")
	filter_strs = filter(slice, secondCharNotIsVowel)
	for _, str := range filter_strs {
		fmt.Println(str)
	}
}
