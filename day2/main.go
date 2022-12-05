package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	//score := solve(scanner)
	scoreCor := solveCorrected(scanner)

	//fmt.Println("score:", score)
	fmt.Println("score corrected:", scoreCor)
	file.Close()
}

func solve(scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		round := strings.Split(text, " ")
		opp := round[0]
		your := round[1]
		score := 0
		if your == "X" {
			score += 1
			if opp == "A" {
				score += 3
			} else if opp == "C" {
				score += 6
			}
		} else if your == "Y" {
			score += 2
			if opp == "B" {
				score += 3
			} else if opp == "A" {
				score += 6
			}
		} else if your == "Z" {
			score += 3
			if opp == "C" {
				score += 3
			} else if opp == "B" {
				score += 6
			}
		}
		acc += score
	}

	return acc
}

func solveCorrected(scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		round := strings.Split(text, " ")
		opp := round[0]
		your := round[1]
		score := 0
		if your == "X" {
			if opp == "A" {
				score += 3
			} else if opp == "B" {
				score += 1
			} else if opp == "C" {
				score += 2
			}
		} else if your == "Y" {
			score += 3
			if opp == "A" {
				score += 1
			} else if opp == "B" {
				score += 2
			} else if opp == "C" {
				score += 3
			}
		} else if your == "Z" {
			score += 6
			if opp == "A" {
				score += 2
			} else if opp == "B" {
				score += 3
			} else if opp == "C" {
				score += 1
			}
		}
		acc += score
	}

	return acc
}
