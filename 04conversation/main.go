package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")

	fmt.Print("Please rate our Pizza between 1 and 5: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		if (numRating < 1 || numRating > 5) {
			fmt.Println("Invalid number, Please enter between 1 to 5 again!")
		} else {
			fmt.Println("Adding 1 to your rating:", numRating + 1)
			fmt.Printf("Type of input is converted: %T ", numRating)
		}
	}

}