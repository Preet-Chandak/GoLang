package main

import "fmt"

func main() {
	var username string = "preet"
	fmt.Println(username)
	fmt.Printf("variable is the type of: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is the type of: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is the type of: %T \n", smallVal)

	var smallFloat float32 = 255.54344
	fmt.Println(smallFloat)
	fmt.Printf("variable is the type of: %T \n", smallFloat)

	//default value
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("type: %T \n", anothervariable)

	// no var sytle

	numberOfUser := 30000.00 //walrus operator is only allowed inside any methods not outside
	fmt.Println(numberOfUser)
}
