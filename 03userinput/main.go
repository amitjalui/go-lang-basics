package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to golang"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err ok (kind of try and catch)
	input, _ := reader.ReadString('\n')
	// input1, _ := reader.ReadByte('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of input is: %T ", input)
	
}