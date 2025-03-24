package main

import "fmt"

func main() {

	// Variable to store user input.
	var input int

	// Declare a random array.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Ask the user which number to search for.
	fmt.Print("[*] Enter a number to search for (1-9): ")
	fmt.Scan(&input)

	// Start searching.
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == input {
			fmt.Printf("Value: %v is found at index %v", input, i)
			break
		}
	}
}
