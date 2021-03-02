package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Go calculator!!!")
	cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide, [u]pdate: ")
	//var register int
	num1, err := readIntFromUser("Enter number:")
	if err != nil {
		panic(err)
	}
	num2, err := readIntFromUser("Enter second number:")
	if err != nil {
		panic(err)
	}
	if cmd == "a" {
		answer := num2 + num1
		fmt.Printf("answer %v", answer)
	} else if cmd == "s" {
		answer := num2 - num1
		fmt.Printf("answer %v", answer)
	} else if cmd == "m" {
		answer := num2 * num1
		fmt.Printf("answer %v", answer)
	} else if cmd == "d" {
		answer := num2 / num1
		fmt.Printf("answer %v", answer)
	} else {
		fmt.Printf("unknown command %q received", cmd)
	}
}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func readIntFromUser(message string) (int, error) {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return strconv.Atoi(input)
}
