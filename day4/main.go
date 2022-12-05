package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	overlapping := solve2(scanner)

	fmt.Println("overlapping ranges:", overlapping)
	file.Close()
}

func solve(scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		text := scanner.Text()
		ranges := strings.Split(text, ",")
		var minMax1 []int
		for _, n := range strings.Split(ranges[0], "-") {
			num, _ := strconv.Atoi(n)
			minMax1 = append(minMax1, num)
		}
		var minMax2 []int
		for _, n := range strings.Split(ranges[1], "-") {
			num, _ := strconv.Atoi(n)
			minMax2 = append(minMax2, num)
		}
		fmt.Println(minMax1, minMax2)
		if minMax1[0] <= minMax2[0] && minMax1[1] >= minMax2[1] {
			acc += 1
		} else if minMax1[0] >= minMax2[0] && minMax1[1] <= minMax2[1] {
			acc += 1
		}
	}

	return acc
}

func solve2(scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		text := scanner.Text()
		ranges := strings.Split(text, ",")
		var minMax1 []int
		for _, n := range strings.Split(ranges[0], "-") {
			num, _ := strconv.Atoi(n)
			minMax1 = append(minMax1, num)
		}
		var minMax2 []int
		for _, n := range strings.Split(ranges[1], "-") {
			num, _ := strconv.Atoi(n)
			minMax2 = append(minMax2, num)
		}
		fmt.Println(minMax1, minMax2)
		if minMax1[0] <= minMax2[0] && minMax1[1] >= minMax2[0] {
			acc += 1
		} else if minMax1[0] <= minMax2[1] && minMax1[1] >= minMax2[1] {
			acc += 1
		} else if minMax1[0] >= minMax2[0] && minMax1[0] <= minMax2[1] {
			acc += 1
		} else if minMax1[1] >= minMax2[0] && minMax1[1] <= minMax2[1] {
			acc += 1
		}
	}

	return acc
}
