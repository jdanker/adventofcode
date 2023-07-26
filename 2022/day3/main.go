package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// if input string is odd it cannot be a valid input
		if len(line)%2 != 0 {
			fmt.Println("Invalid input. Line is odd number of characters.")
			return
			
		}

		half := len(line) / 2
		left := line[:half]
		right := line[half:]





		
	}


}

func getPriority(r rune) int {
	return 0
}

func compareString(s1, s2 string) rune {
	return 'a'
}