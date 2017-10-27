// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 5 taken from https://data-representation.github.io/problems/go-fundamentals.html
// Code adapted from http://golangcookbook.blogspot.ie/2012/11/generate-random-number-in-given-range.html
package main

// Import 'fmt' package so we can use Printf AND Scanf(https://golang.org/pkg/fmt/)
import "fmt"

// Import the math/rand function so we can generate a pseudo random number for our guessing game (https://golang.org/pkg/math/rand/)
import "math/rand"

// Import time to use as a seed for our number generator i.e every time the program is run, the time is used as a seed to generate a random number (https://golang.org/pkg/time/)
import "time"


// Function to generate random number	
func random(min, max int) int{

				// Use the time of our computer in nanoseconds, as a seed for our RNG each time it's run
				rand.Seed(time.Now().UnixNano());
				return rand.Intn(max-min)+min;
				
			}
			

// Main method
func main(){

	// Set range for RNG
	// Pass thes values into random function above
	myRandom := random(1,100)

	// Variables
	var userGuess int;
	var ifCorrect bool=true;
	var counter int=0;
	var previousNumber int;
	
	
	// Print out randomly generated number (for testing purposes)
	//fmt.Println(myRandom)


	// If statement to let the user keep guessing until they are correct
	for(ifCorrect==true){
				
				// Prompt the user for their guess
				fmt.Print("Please enter your guess(1-100): ");
				fmt.Scan(&userGuess);
				fmt.Println();

				if(userGuess==previousNumber) {
					
					// Don't count the guess if they have already tried that number
					counter --;
					fmt.Println("Number already chosen!\n");
				}

				// Increase counter to count number of guesses
				counter++;

						// If statement to test if the user's guess was right
						if(userGuess==myRandom){
						fmt.Printf("Congratualtions, your guess of %d was correct!\n", userGuess);
						fmt.Printf("It only took you %d tries!\n", counter);

						// Exit the loop if the user is correct
						ifCorrect=false;

					}	else if(userGuess>myRandom){
						fmt.Printf("Your guess was too high\n");
					}   else if(userGuess<myRandom){
						fmt.Printf("Your guess was too low\n");
					}
					
					// If the user guesses the same number again, tell them and don't count it as a guess
					previousNumber = userGuess;
			
	}
	

	
}


