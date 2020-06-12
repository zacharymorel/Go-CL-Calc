package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
* TODO:
	1. Percentage?
	2. Square Root?
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	printHowTo()
	var currentNum int = 0
	var currentCommand string
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.TrimSpace(text)

		// switch here to recognize commands
		switch text {
		case "+":
			currentCommand = text
		case "-":
			currentCommand = text
		case "/":
			currentCommand = text
		case "*":
			currentCommand = text
		case "C":
			currentCommand = ""
			currentNum = 0
		default:
			currentNum = doArithmetic(text, currentNum, currentCommand)
			currentCommand = "" // reset to prevent multiple uses
		}
		fmt.Println("currentNum: ", currentNum)
	}
}

func doArithmetic(text string, curNum int, curCommand string) int {
	var num int
	i, err := strconv.Atoi(text)

	if err != nil { // Exit clause
		fmt.Println(text + " is not a number. GoCalc is expecting a number.")
		return curNum // return current number so there is no override of state
	} else if curNum == 0 {
		num = i // return current number so there is no override of state
		return num
	} else {
		// arithmetic
		switch curCommand {
		case "+":
			num = add(curNum, i)
		case "-":
			num = subtract(curNum, i)
		case "/":
			num = divide(curNum, i)
		case "*":
			num = multiply(curNum, i)
		default:
			return curNum // case no command and curNum is not 0, return value
		}
	}

	return num
}

func printHowTo() {
	fmt.Println("Welcome to GoCalc!")
	fmt.Println("--------------------------------")
	fmt.Println("Here is a list of the following commands: ")
	fmt.Println("+ is addition")
	fmt.Println("- is subtraction")
	fmt.Println("/ is division")
	fmt.Println("* is multiplication")
	fmt.Println("C is CLEAR")
	fmt.Println("--------------------------------")
	fmt.Println("Enter in a number or commander once at a time.")
}

func add(cur, add int) int {
	return cur + add
}

func subtract(crr, sub int) int {
	return crr - sub
}

func divide(crr, div int) int {
	return crr / div
}

func multiply(crr, multi int) int {
	return crr * multi
}
