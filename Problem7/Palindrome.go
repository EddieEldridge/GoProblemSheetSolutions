// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 7 taken from https://data-representation.github.io/problems/go-fundamentals.html
// Code adapted from https://play.golang.org/p/5FUOzjBa-o

package main

// Import 'fmt' package so we can use Printf AND Scanf(https://golang.org/pkg/fmt/)
import "fmt"

// Function to check if string is a palindrome
// returns a boolean if the inputted string is/isn't a palindrome
func isPalindrome(input string) bool{

	// Divide the string into two parts, starting from the back and front
	// Compare them, if they are not equal to each other then return false
	// else return true
	for i := 0; i < len(input)/2; i++ {
		if (input[i] != input[len(input)-i-1]){
			return false;
		}
	}
	return true;
}

// Main method
func main(){
	
	// Depending on the value put in here, a true or false value will
	// be outputted to the console
	// indicating if the string is a palindrome or not
	fmt.Println(isPalindrome("anna"))
}


