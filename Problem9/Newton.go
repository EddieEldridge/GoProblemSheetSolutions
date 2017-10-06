// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 9 taken from https://data-representation.github.io/problems/go-fundamentals.html
// Code adapted from https://tour.golang.org/flowcontrol/8

package main

// Import 'fmt' package so we can use Printf AND Scanf(https://golang.org/pkg/fmt/)
import "fmt"

// Import 'math' package so we can use sqrt
import "math"

func newtonFunction(testNumber float64,z float64) float64 {

	// Run newtons method with our testNumber inputted by the user
	return z - (((z*z) - testNumber) / (2*z))
	
}

func main() {

	// Variables
	var counter int = 0;

	// Assign a value to testnumber
	var testNumber float64=5;

	
	// initial guess
	z := float64(1);
	
	// Loop using newtons method
	for z = 1.0; z != newtonFunction(testNumber,z); z = newtonFunction(testNumber,z){
		

			counter++
			fmt.Printf("Guess %d :%12.10f \n",counter,z)
		}

	// Output results
	fmt.Printf("The square of %.2f is %12.10f according to Newtons Theory\n",testNumber,z)
	fmt.Printf("The square of %.2f is %.8f according to 'math.Sqrt(testNumber)'",testNumber,math.Sqrt(testNumber))

}




