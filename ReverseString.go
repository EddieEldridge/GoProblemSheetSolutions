// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 10 taken from https://data-representation.github.io/problems/go-fundamentals.html
// Code adapted from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go and https://www.dotnetperls.com/reverse-string-go

package main

// Import 'fmt' package so we can use Printf AND Scanf(https://golang.org/pkg/fmt/)
import "fmt"

// Function to reverse the string entered by the user
func reverseString(value string) string{
	
	// Convert string to rune slice
	data := []rune(value)
	result := []rune{}
	
	// Add runes in reverse order.
	  for i := len(data) - 1; i >= 0; i-- {
        result = append(result, data[i])
    }

	// Return reversed string 
	return string(result)
}

func main() {


	// Prompt user for their string
	//fmt.Print("Please enter the string you would like to reverse: ")
	//fmt.Scanf("%s", userString)

	// Enter the string you would like to reverse
	userString := "eddie"
	reversedString := reverseString(userString)

	// Print out results
	fmt.Printf("Original String: %s", userString)
	fmt.Printf("Reversed String: %s", reversedString)
	
	


}

