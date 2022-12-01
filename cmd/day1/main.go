package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input/day1/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	highest, top3tot := solve(scanner)

	fmt.Println("highest cal elf:", highest)
	fmt.Println("top three elf cal total:", top3tot)
	file.Close()
}

func solve(scanner *bufio.Scanner) (int, int) {
	var cals []int
	acc := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			fmt.Println("total for elf:", acc)
			if len(cals) == 0 {
				cals = insertAtIndex(cals, acc, 0)
			}
			for index, cal := range cals {
				if acc > cal {
					cals = insertAtIndex(cals, acc, index)
					break
				}
			}
			acc = 0
			continue
		}
		num, _ := strconv.Atoi(text)
		acc += num
	}

	return cals[0], cals[0] + cals[1] + cals[2]
}

func insertAtIndex(array []int, val, index int) []int {
	cur := val
	for i := index; i < len(array); i++ {
		prev := cur
		cur = array[i]
		array[i] = prev
	}
	return append(array, cur)
}
