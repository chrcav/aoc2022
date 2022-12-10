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

	uniquetaillen1 := solve(lines, 2)
	fmt.Println("unique spots visited:", uniquetaillen1)
	uniquetaillen10 := solve(lines, 10)
	fmt.Println("unique spots for len 10:", uniquetaillen10)
}

func solve(lines []string, len int) int {
	visitedCount := 1
	var ropes [][2]int
	for i := 0; i < len; i++ {
		ropes = append(ropes, [2]int{0, 0})
	}
	var visited [][]int
	visited = append(visited, []int{0, 0})
	for _, text := range lines {
		dis, dir := parseInstruction(text)
		for i := 0; i < dis; i++ {
			if dir == "R" {
				ropes[0][0] += 1
			}
			if dir == "L" {
				ropes[0][0] -= 1
			}
			if dir == "U" {
				ropes[0][1] += 1
			}
			if dir == "D" {
				ropes[0][1] -= 1
			}
			for j := 1; j < len; j++ {
				xdiff := ropes[j-1][0] - ropes[j][0]
				ydiff := ropes[j-1][1] - ropes[j][1]
				if abs(xdiff) > 1 || abs(ydiff) > 1 {
					//fmt.Println("x diff:", xdiff)
					if abs(xdiff) > 0 {
						ropes[j][0] += xdiff / abs(xdiff)
					}
					//fmt.Println("y diff:", ydiff)
					if abs(ydiff) > 0 {
						ropes[j][1] += ydiff / abs(ydiff)
					}
				}
			}
			beenHereBefore := false
			for _, loc := range visited {
				if ropes[len-1][0] == loc[0] && ropes[len-1][1] == loc[1] {
					beenHereBefore = true
					break
				}
			}
			if !beenHereBefore {
				visitedCount += 1
				visited = append(visited, []int{ropes[len-1][0], ropes[len-1][1]})
			}
		}
	}

	return visitedCount
}

func parseInstruction(instr string) (int, string) {
	sp := strings.Split(instr, " ")
	dis, _ := strconv.Atoi(sp[1])
	return dis, sp[0]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
