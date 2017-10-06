// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 7 taken from https://data-representation.github.io/problems/go-fundamentals.html


package main

// Import 'fmt' package so we can use Printf AND Scanf(https://golang.org/pkg/fmt/)
import "fmt"



// Main method
func main(){

	var i int;
	var lowestNum int;
	var highestNum int;

	// Declare our list to loop through
	numbers := [10]int{1,5,4,20,30,10,25,41,99,21}
	
	lowestNum = numbers[0];
	highestNum = numbers[0];

	// Loop through our list using range
	for i= range numbers{
			if(numbers[i] < lowestNum){
				lowestNum = numbers [i];
			}
			if(numbers[i] > highestNum){
				highestNum = numbers [i];
			}

	}

	// Print out our highest and lowest numbers
	fmt.Printf("Highest number: %d\n", highestNum)
	fmt.Printf("Lowest number: %d\n", lowestNum)
	
}


