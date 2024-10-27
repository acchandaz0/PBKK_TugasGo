package main

import "fmt"

func main() {
	var input int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&input)

	if input == 42 {
		fmt.Println("Hello Universe")
	} else {
		fmt.Println(input)
	}
}
