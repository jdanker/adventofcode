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
		common := compareString(left, right)
		fmt.Printf("%c\n", common)
		
		
	}


}

func getPriority(r rune) int {
	return 0
}

func compareString(s1, s2 string) rune {
	s1Chars := make(map[rune]bool)
	s2Chars := make(map[rune]bool)

	// add the characters of each string to a map
	for _, c := range s1 {
		s1Chars[c] = true
	}
	for _, c := range s2 {
		s2Chars[c] = true
	}
	// compare the maps to find the first character that is in both strings
	for c := range s1Chars {
		if s2Chars[c] {
			return c
		}
	}

	return 0
}