package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var elfNum int = 0;
	var calories int =0;
	readFile, err := os.Open("input.txt")

	if err != nil {
        fmt.Println(err)
    }

	fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		numStr := fileScanner.Text();
		if len(numStr) != 0 {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				calories=calories+num;
			}


		} else {
			elfNum++

		}


        fmt.Println(fileScanner.Text())
    }

	readFile.Close()


}


