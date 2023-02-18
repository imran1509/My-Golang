package main

import "fmt"

const LoginToken string = "kabckjabc" //Public constant : when first letter of var is capital

func main() {
	var username string = "Imran"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of type: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.16516564651354651
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//defult values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type: %T \n", anothervariable)

	//implicit type
	var website = "https://linkfree.eddiehub.io/imran1509"
	fmt.Println(website)

	//noVarStyle
	NumberOfVisitors := 2500
	fmt.Println(NumberOfVisitors)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
