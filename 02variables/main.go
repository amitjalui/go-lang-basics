package main

import "fmt"

// jwtToken := 30000 		// ❎ walrus operation only works inside main method
// var jwtToken = 30000 		// ✅
// var jwtToken int = 30000 		// ✅

const LoginToken string = "abc"		// variable name staring with CAP letter is declared as PUBLIC variables.

func main() {
	var username string = "amit"
	fmt.Println(username)
	fmt.Printf("var type is: %T \n", username)

	fmt.Println("--------------")
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("var type is: %T \n", isLoggedIn)

	fmt.Println("--------------")
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("var type is: %T \n", smallVal)

	fmt.Println("--------------")
	
	var smallFloat32 float32 = 255.56546456524
	fmt.Println(smallFloat32)
	fmt.Printf("var type is: %T \n", smallFloat32)

	fmt.Println("--------------")
	
	var smallFloat64 float64 = 255.56546456524
	fmt.Println(smallFloat64)
	fmt.Printf("var type is: %T \n", smallFloat64)

	fmt.Println("--------------")
	
	// DEFAULT VALUES
	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("var type is: %T \n", defaultInt)

	fmt.Println("--------------")
	
	// IMPLICIT TYPE
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("var type is: %T \n", website)

	fmt.Println("--------------")
	
	// NO VAR STYLE (walrus operator)
	numberOfUser := 3000.0
	fmt.Println(numberOfUser)
	fmt.Printf("var type is: %T \n", numberOfUser)

	// Calling PUBLIC variables
	fmt.Println(LoginToken)
	fmt.Printf("var type of Public Variable: %T \n", LoginToken)


}