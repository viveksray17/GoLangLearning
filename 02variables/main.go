package main

import "fmt"

const LoginToken string = "thisisalogintoken" // public variable with first letter capital

func main() {
	var username string = "vivek"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float32 = 255.45555554352 // float 32 has less precision as compared to float64
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable) // go doesn't put garbage value to the variable
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	var anotherVariableString string
	fmt.Println(anotherVariableString) // same for string
	fmt.Printf("Variable is of type: %T\n", anotherVariableString)

	// no var style declaration
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
