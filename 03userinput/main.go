package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcome string = "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our program: ")

	// comma ok // error ok
	input, _ := reader.ReadString('\n') // _ is for catching the error
	fmt.Printf("Thanks for rating, %v", input)
	fmt.Printf("Type of this rating is %T\n", input)
}
