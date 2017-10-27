// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 4 taken from https://data-representation.github.io/problems/go-fundamentals.html
// Code adapted from http://codecry.com/go/factorial-of-a-number

package main

// Import 'fmt' package so we can use Printf (https://golang.org/pkg/fmt/)
import "fmt"

	

// Main method
func main(){

	// Assign variables
	var i uint;
	var output uint=1;
	var sum uint=0;


	// Change the value of number to choose what number 
	// to get the factorial of
	var number uint = 10;

	// For loop to calculate factorial and sum
	for i=1; i<=number; i++{
		
		output = output * i;
		sum = sum + i;
	}

	// Print result of calcFactorial to the screen
	fmt.Printf("%d! =  %d\n",number,output)

	// Print result of sum of factorial digits to the screen
	fmt.Printf("Sum of numbers in factorial result = %d (not correct)", sum)
	
}


