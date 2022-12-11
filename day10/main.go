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
	var lines []string
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}
	file.Close()

	signalSum := solve(lines)
	fmt.Println("signal strength sum:", signalSum)
}

func solve(lines []string) int {
	signalSum := 0
	cycles := 0
	X := 1
	for _, text := range lines {
		com := strings.Split(text, " ")
		cyc := getComCycles(com)
		for i := 0; i < cyc; i++ {
			if X+1 >= cycles%40 && X-1 <= cycles%40 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			cycles += 1
			if cycles%40 == 0 {
				fmt.Println()
			}
			if cycles%40 == 20 {
				signalSum += X * cycles
			}
		}
		if com[0] == "addx" {
			val, _ := strconv.Atoi(com[1])
			X += val
		}
	}

	return signalSum
}

func getComCycles(com []string) int {
	if len(com) == 1 {
		return 1
	} else if len(com) == 2 {
		if com[0] == "addx" {
			return 2
		}
	}
	return 0
}
