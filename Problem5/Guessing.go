// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 5 taken from https://data-representation.github.io/problems/go-fundamentals.html
// Code adapted from http://codecry.com/go/factorial-of-a-number

package main

// Import 'fmt' package so we can use Printf AND Scanf(https://golang.org/pkg/fmt/)
import "fmt"

// Import the math/rand function so we can generate a pseudo random number for our guessing game (https://golang.org/pkg/math/rand/)
import "math/rand"

// Import time to use as a seed for our number generator i.e every time the program is run, the time is used as a seed to generate a random number (https://golang.org/pkg/time/)
import "time"

// Main method
func main(){

	// Variables
	//var userGuess int;

	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed)
	
	// Prompt the user for their guess
	fmt.Printf("Please enter your guess: ");

	// Code adapted from (https://stackoverflow.com/questions/14000082/how-do-i-use-fmt-scanf-in-go)
	// incase of Windows line ending problem

	
	// Print out randomly generated number (for testing purposes)
	fmt.Print(randomNumber)
	// If statement to test if the user's guess was right
	//if(userGuess==randomNumber){
	//fmt.Printf("Congratualtions, your guess of %d was correct", userGuess);
	//}
}


