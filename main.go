package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Go calculator!")
	for {
		cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide, [u]pdate: ")

		num1, err := readIntFromUser("Enter number:")
		if err != nil {
			panic(err)
		}
		num2, err := readIntFromUser("Enter second number:")
		if err != nil {
			panic(err)
		}
		var answer int
		if cmd == "a" {
			answer = num2 + num1
		} else if cmd == "s" {
			answer = num2 - num1
		} else if cmd == "m" {
			answer = num2 * num1
		} else if cmd == "d" {
			answer = num2 / num1
		} else {
			fmt.Printf("unknown command %q received", cmd)
		}
		fmt.Printf("answer %v\n", answer)
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
