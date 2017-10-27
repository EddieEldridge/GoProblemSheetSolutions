// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 8 taken from https://data-representation.github.io/problems/go-fundamentals.html
// Code adapted from http://austingwalters.com/merge-sort-in-go-golang/

package main

// Import 'fmt' package so we can use Printf AND Scanf(https://golang.org/pkg/fmt/)
import "fmt"

// Import 'sort' to sort our lists (https://golang.org/pkg/sort/)
import "sort"



func splitSlices() {

	// Declare two lists, odd and even so it's easy to test for us to test if we merged and sorted them correctly
	evenList := []int{2,4,6,8,10}
	oddList := []int{1,3,5,7,9}
	
	// Print out the unsorted evenList
	fmt.Println("\n");	
	fmt.Println("evenList (Unsorted)\n")
	fmt.Println(evenList)
	fmt.Println("\n");

	// Print out the unsorted oddList
	fmt.Println("oddList (Unsorted)\n")
	fmt.Println(oddList)
	fmt.Println("\n");	

	// Append the oddList to the evenList
	evenList = append(evenList, oddList...)

	// Print out the appended slice
	fmt.Println("Appeneded Slice (Unsorted)\n")
	fmt.Println(evenList)
	fmt.Println("\n");	

	// Sort the evenList after you have added the oddList to it
	sort.Ints(evenList)

	// Print out the appended slice
	fmt.Println("Appended Slice (Sorted)\n")
	fmt.Println(evenList)
	fmt.Println("\n");	
}
func main() {
	
	splitSlices();
}
	



