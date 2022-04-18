package main

import "fmt"

func main() {
	result := Status(5)
	if result {
		fmt.Println("Lights ON")
	} else {
		fmt.Println("Lights OFF")
	}
}

func Status(input int) bool {
	result := false

	for i := 0; i <= input; i++ {
		if i%2 == 0 {
			result = true
		} else {
			result = false
		}
	}

	return result
}
