package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	ruckSum := solve2(scanner)

	fmt.Println("ruck sum:", ruckSum)
	file.Close()
}

func solve(scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		text := scanner.Text()
		comSize := len(text) / 2
		firstCom := []rune(text[0:comSize])
		secondCom := []rune(text[comSize : comSize*2])
	outer:
		for _, r1 := range firstCom {
			for _, r2 := range secondCom {
				if r1 == r2 {
					if r1 >= 'a' {
						acc += int(r1) - 96
					} else {
						acc += int(r1) - 38
					}
					break outer
				}
			}
		}
	}
	return acc
}

func solve2(scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		sack1 := []rune(scanner.Text())
		scanner.Scan()
		sack2 := []rune(scanner.Text())
		scanner.Scan()
		sack3 := []rune(scanner.Text())
	outer:
		for _, r1 := range sack1 {
			for _, r2 := range sack2 {
				for _, r3 := range sack3 {
					if r1 == r2 && r1 == r3 {
						if r1 >= 'a' {
							acc += int(r1) - 96
						} else {
							acc += int(r1) - 38
						}
						break outer
					}
				}
			}
		}
	}
	return acc
}
