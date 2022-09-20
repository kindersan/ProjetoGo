// Golang program to demonstrate the
// example of append() function

package main

import (
	"fmt"
)

func main() {
	// Creating int and string slices
	s1 := []int{10, 20, 30}
	s2 := []string{"Hello", "World"}
	fmt.Println("S1 :",cap(s1))
	// Creating slices to append
	slice1 := []int{40, 50}
	slice2 := []string{"How're you?", "Boys"}

	// Printing types and values of slices
	fmt.Printf("%T, %v\n", s1, s1)
	fmt.Printf("%T, %q\n", s2, s2)

	// Appending the slices
	s1 = append(s1, slice1...)
	s2 = append(s2, slice2...)
	fmt.Println("S1:",cap(s1))

	// After appending,
	// Printing types and values of slices
	fmt.Println("After appending...")
	fmt.Printf("%T, %v\n", s1, s1)
	fmt.Printf("%T, %q\n", s2, s2)
}

