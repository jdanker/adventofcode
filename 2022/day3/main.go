package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")
	prioritySum := 0

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	
	// slice to hold three lines read from the file
	lastThreeLines := make([]string, 3)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// if input string is odd it cannot be a valid input
		if len(line)%2 != 0 {
			fmt.Println("Invalid input. Line is odd number of characters.")
			return
			
		}
		
		// removes the slice if there are 3 lines, adds line to the slice if not
		if len(lastThreeLines) == 3 {
			lastThreeLines = lastThreeLines[:0]
		} else if len(lastThreeLines) < 3 {
			lastThreeLines = append(lastThreeLines, line)
		} 




		half := len(line) / 2
		left := line[:half]
		right := line[half:]
		common := compareString(left, right)
		fmt.Printf("%c\n", common)

		priority := getPriority(common)
		if(priority != 0){
			prioritySum += priority
		}


		
	}
	println(prioritySum)


}

// subtracts unicode value -1 from the character to get the priority
func getPriority(r rune) int {
	if r>='a' && r<='z' {
		return int(r-96)
	} else if r>='A' && r<='Z' {
		return int(r-38)
	}
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

func getBadge(lastThree []string) int {
	
	// TODO: compare 3 strings to find the common character
	return 0
}