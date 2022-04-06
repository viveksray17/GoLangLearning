package main

import "fmt"

func main() {
	fmt.Println("Welcome to the video on slices")

	// var fruitList = []string{"Apple", "Banana", "Peach"}
	// fmt.Printf("Type of fruitlist is %T\n", fruitList)
	//
	// fruitList = append(fruitList, "Mango", "Tomato")
	// fmt.Printf("%v\n", fruitList)
	// fruitList = append(fruitList[1:3]) // from index of 1 to index 2 (since 3(the second parameter) is exclusive in ranges)
	// fmt.Println(fruitList)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 235
	highscores[2] = 234
	highscores[3] = 234
}
