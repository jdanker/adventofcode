// DAY 2
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	totalScore := 0
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		if len(line) != 0 {
			roundScore := getScore(line[0:1], line[2:3])
			fmt.Println("Round Score: ", roundScore)
			totalScore += roundScore
		}
	}
	fmt.Println("Total Score: ", totalScore)

}

func getWinner(player1, player2 string) string {
	switch player1 {
		case "A":
			switch player2 {
				case "X":
					return "tie"
				case "Y":
					return "player2"
				case "Z":
					return "player1"
			}
		case "B":
			switch player2 {
				case "X":
					return "player1"
				case "Y":
					return "tie"
				case "Z":
					return "player2"
			}
		case "C":
			switch player2 {
				case "X":
					return "player2"
				case "Y":
					return "player1"
				case "Z":
					return "tie"
			}
	}
	return "error"
}

func getScore(a, b string) int{
	b = getMove(a, b)
	fmt.Println("Player 1: ", a , "Player 2: ", b)
	score := 0
	switch b {
	case "X":
		score++
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	winner := getWinner(a, b)

	switch winner {
		case "player2":
			score += 6
		case "tie":
			score += 3
	}
	return score
}


func getMove(p1, p2 string)string {
	switch p1 {
	case "A":
		switch p2 {
			case "X":
				return "Z"
			case "Y":
				return "X"
			case "Z":
				return "Y"
		}
	case "B":
		switch p2 {
			case "X":
				return "X"
			case "Y":
				return "Y"
			case "Z":
				return "Z"
		}
	case "C":
		switch p2 {
			case "X":
				return "Y"
			case "Y":
				return "Z"
			case "Z":
				return "X"
		}
}
return "error"
}
