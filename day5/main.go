package main

import (
	"bufio"
	"container/list"
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

	stackTops := solve(lines)
	stackTopsCorrect := solve2(lines)

	fmt.Println("stack top sequence:", stackTops)
	fmt.Println("stack top sequence correct:", stackTopsCorrect)
}

func solve(lines []string) string {
	acc := ""
	var stacks []*list.List
	instrLine := 0
	for lineInd, text := range lines {
		if strings.Contains(text, "1") {
			instrLine = lineInd + 2
			break
		}
		for i := 0; i < len(text); i += 4 {
			stackCol := text[i+1 : i+2]
			if i/4 >= len(stacks) {
				stacks = append(stacks, list.New())
			}
			if strings.TrimSpace(stackCol) != "" {
				stacks[i/4].PushBack(strings.TrimSpace(stackCol))
			}
		}
	}

	for lineInd := instrLine; lineInd < len(lines); lineInd++ {
		text := lines[lineInd]
		moves := strings.Split(text, " ")
		numCrates, _ := strconv.Atoi(moves[1])
		fromStack, _ := strconv.Atoi(moves[3])
		toStack, _ := strconv.Atoi(moves[5])
		for i := 0; i < numCrates; i++ {
			stacks[toStack-1].PushFront(stacks[fromStack-1].Remove(stacks[fromStack-1].Front()))
		}
	}

	for _, stack := range stacks {
		acc += stack.Front().Value.(string)
	}

	return acc
}

func solve2(lines []string) string {
	acc := ""
	var stacks []*list.List
	instrLine := 0
	for lineInd, text := range lines {
		if strings.Contains(text, "1") {
			instrLine = lineInd + 2
			break
		}
		for i := 0; i < len(text); i += 4 {
			stackCol := text[i+1 : i+2]
			if i/4 >= len(stacks) {
				stacks = append(stacks, list.New())
			}
			if strings.TrimSpace(stackCol) != "" {
				stacks[i/4].PushBack(strings.TrimSpace(stackCol))
			}
		}
	}

	for lineInd := instrLine; lineInd < len(lines); lineInd++ {
		text := lines[lineInd]
		moves := strings.Split(text, " ")
		numCrates, _ := strconv.Atoi(moves[1])
		fromStack, _ := strconv.Atoi(moves[3])
		toStack, _ := strconv.Atoi(moves[5])
		toStackFront := stacks[toStack-1].Front()
		for i := 0; i < numCrates; i++ {
			if stacks[toStack-1].Len() != i {
				stacks[toStack-1].InsertBefore(stacks[fromStack-1].Remove(stacks[fromStack-1].Front()), toStackFront)
			} else {
				stacks[toStack-1].PushBack(stacks[fromStack-1].Remove(stacks[fromStack-1].Front()))
			}
		}
	}

	for _, stack := range stacks {
		acc += stack.Front().Value.(string)
	}

	return acc
}
