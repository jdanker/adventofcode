package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var elfNum int = 0
	var calories int = 0
	elfList := make([]int, 0) // creates an empty slice
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		numStr := fileScanner.Text()
		if len(numStr) != 0 {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err)

			} else {
				calories = calories + num
				elfList = append(elfList, calories)
			}
		} else {
			fmt.Println("Elf Number: ", elfNum, "Calories: ", calories)
			calories = 0
			elfNum++
		}

	}

	max := elfList[0]

	for _, num := range elfList {
		if num > max {
			max = num
		}
	}

	fmt.Println("Most Calories:", max)
	readFile.Close()

}
