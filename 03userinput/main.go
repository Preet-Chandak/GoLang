package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "WELCOME TO USER INPUT PROGRAMMER"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	//comma ok //err ok

	// input, err := reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating", input)

	fmt.Printf("Type of input %T", input)

}
