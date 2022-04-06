package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Pointers")
	// var ptr *int
	// fmt.Printf("The value of pointer is %v\n", ptr)

	myNumber := 23
	var ptr = &myNumber                      // ptr is a reference to the direct memory allocation
	fmt.Printf("Value of ptr is %v\n", ptr)  // this returns something like 0xc000014108
	fmt.Printf("Value of ptr is %v\n", *ptr) // this returns the actual value

	*ptr = *ptr + 2
	fmt.Printf("The value of myNumber is %v\n", myNumber)
}
