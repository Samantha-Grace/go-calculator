package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Go calculator!")
	cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide, [u]pdate: ")

	//var register int := readRegister("0")

	num1, err := readIntFromUser("Enter number:")
	if err != nil {
		panic(err)
	}
	num2, err := readIntFromUser("Enter second number:")
	if err != nil {
		panic(err)
	}

	num1, num2, err := readIntFromUser("0")
	if answer = int {
		fmt.Print(message)
	}



	//output := readIntFromUser("0")

	if cmd == "a" {
		answer := num2 + num1
		fmt.Printf("answer %v", answer)

		//output := readRegister("0")
		//fmt.Printf("output %v", output)

	} else if cmd == "s" {
		answer := num2 - num1
		//output := "0"
		fmt.Printf("answer %v", answer)
		//fmt.Printf("output %v", output)
	} else if cmd == "m" {
		answer := num2 * num1
		fmt.Printf("answer %v", answer)
	} else if cmd == "d" {
		answer := num2 / num1
		fmt.Printf("answer %v", answer)
		if num2 == 0 {
			panic("you can not divide to 0")
		}
	} else {
		fmt.Printf("unknown command %q received", cmd)

	} 
	
	  switch answer := "0"; answer{

		}
	}

/*
func output() {
	fmt.Println("0")
	//cmd := readLine("0")
}
*/
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

/*
func readRegister(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input



		switch == "a" {
		"0" := answer
		fmt.Printf("answer %v", answer)
}
*/



