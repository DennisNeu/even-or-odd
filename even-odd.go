package main

import (
	"fmt"
)

func main() {

	var input int = -1

	for {
		fmt.Println("Input a positive number.")
		fmt.Scan(&input)
		fmt.Println(input)

		if input > 0 {
			break
		}

	}

	if input%2 == 0 {
		fmt.Println("The number", input, "is even.")
	} else {
		fmt.Println("The number", input, "is odd.")
	}

}
