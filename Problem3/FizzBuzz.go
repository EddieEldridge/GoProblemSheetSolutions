// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 3 taken from https://data-representation.github.io/problems/go-fundamentals.html
// Code taken from https://golangcode.com/fizz-buzz-test-in-go/

package main

// Import 'fmt' package so we can use Printf (https://golang.org/pkg/fmt/)
import "fmt"

// Main method
func main(){

	// Declare variable for use in the loop
	var i int;

	// For loop to loop through numbers 1-100
	for i= 1; i <= 100; i++ {
		
				if i%3 == 0 {
					// Multiple of 3
					fmt.Printf("fizz")
				}
				if i%5 == 0 {
					// Multiple of 5
					fmt.Printf("buzz")
				}
		
				if i%3 != 0 && i%5 != 0 {
					// Neither, so print the number itself
					fmt.Printf("%d", i)
				}
		
				// A trailing new line (so both fizz + buzz can be printed on the same line)
				fmt.Printf("\n")
		
			}
		}