package main

import "fmt"

type age int16

func main() {
	fmt.Println("This is the main function to start program")
}

func sum(a int16, b int16) int16 {
	return a + b
}
