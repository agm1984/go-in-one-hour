package main

import "fmt"

// https://www.youtube.com/watch?v=CF9S4QZuV30
// start at 17:17
func main() {
	numSlice := []int{5, 4, 3, 2, 1}
	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice2[0] =", numSlice2[0])
	fmt.Println("numSlice2[:2] =", numSlice2[:2])
	fmt.Println("numSlice2[0:] =", numSlice2[0:]) // returns [2 1]
}
