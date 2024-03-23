package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my pizza app")
	fmt.Println("Rate our pizza between 1-5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input) // Trim whitespace and newline characters from input

	fmt.Println("Thanks for Rating!!! :", input)

	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added one to your rating:", numRating+1)
	}

}
