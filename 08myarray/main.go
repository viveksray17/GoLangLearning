package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"
	fmt.Printf("Fruit List is %v\n", fruitList)
	fmt.Printf("Fruit List is %v\n", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Printf("Vegetable List is %v\n", vegList)
	fmt.Printf("Vegetable List is %v\n", vegList[0])
}
