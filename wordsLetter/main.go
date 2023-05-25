package main

import "fmt"

func main() {
	letter := "12345"
	println(len(letter))

	for i := 0; i < len(letter); i++ {
		fmt.Printf("%c\n", letter[i])
	}
}
