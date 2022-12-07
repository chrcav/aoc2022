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
	var lines []string
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}
	file.Close()

	charsCount := solve(lines, 4)
	fmt.Println("start-of-packet at:", charsCount)
	charsCount = solve(lines, 14)
	fmt.Println("start-of-message at:", charsCount)
}

func solve(lines []string, uniqueSeqLen int) int {
	acc := 0
	for _, text := range lines {
		for i := 0; i < len(text)-uniqueSeqLen; i++ {
			window := []rune(text[i : i+uniqueSeqLen])
			dup := false
		window:
			for j, c1 := range window {
				for k, c2 := range window {
					if j == k {
						continue
					}
					if c1 == c2 {
						dup = true
						break window
					}

				}
			}
			if !dup {
				acc = i + uniqueSeqLen
				break
			}
		}
	}

	return acc
}
