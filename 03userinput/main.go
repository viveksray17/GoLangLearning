package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var welcome string = "Welcome to the calculator app"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the first number: ")
	firstnum, _ := reader.ReadString('\n')
	fmt.Print("Enter the second number: ")
	secondnum, _ := reader.ReadString('\n')

	intfirstnum, _ := strconv.ParseFloat(strings.TrimSpace(firstnum), 64)
	intsecondnum, _ := strconv.ParseFloat(strings.TrimSpace(secondnum), 64)
	fmt.Printf("The sum of both the numbers is %v\n", intfirstnum+intsecondnum)
}
