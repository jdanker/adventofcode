package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	readFile, err := os.Open("input.txt")
	var sum int

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := []int{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				num := int(char - '0') // converting rune to int
				numbers = append(numbers, num)
			}

		}
		if len(numbers) > 0 {
			first := numbers[0]
			last := numbers[len(numbers)-1]
			calValue := num(first, last)
			sum += calValue
		} else {
			fmt.Println("Info: The slice is empty")
		}

	}
	fmt.Println("Calibration Value: ",sum)
	readFile.Close()

}

func num(first int, last int) int {
	var result int

	str := fmt.Sprintf("%d%d", first, last)
	var err error
	result, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: Could not convert string to integer")
		os.Exit(0)
	}

	return result
}
