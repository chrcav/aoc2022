package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	visibleTrees, scenicScore := solve(lines)
	fmt.Println("trees visible:", visibleTrees)
	fmt.Println("trees visible from house:", scenicScore)
}

func solve(lines []string) (int, int) {
	visibleTrees := 0
	scenicScore := 0
	for j, text := range lines {
		for i := 0; i < len(text); i++ {
			h0, _ := strconv.Atoi(text[i : i+1])
			visible, vfh := calcVisibleTrees(lines, h0, i, j)
			if visible {
				visibleTrees += 1
			}
			if vfh > scenicScore {
				scenicScore = vfh
			}
		}
	}

	return visibleTrees, scenicScore
}

func calcVisibleTrees(grid []string, h0, tx, ty int) (bool, int) {
	visible := true
	vt := 0
	y := ty
	westvfh := 0
	for x := tx - 1; x >= 0; x-- {
		h1, _ := strconv.Atoi(grid[y][x : x+1])
		westvfh += 1
		if h1 >= h0 {
			visible = false
			break
		}
	}
	if visible {
		vt += 1
	}
	visible = true
	eastvfh := 0
	for x := tx + 1; x < len(grid[0]); x++ {
		h1, _ := strconv.Atoi(grid[y][x : x+1])
		eastvfh += 1
		if h1 >= h0 {
			visible = false
			break
		}
	}
	if visible {
		vt += 1
	}
	visible = true
	x := tx
	northvfh := 0
	for y := ty - 1; y >= 0; y-- {
		h1, _ := strconv.Atoi(grid[y][x : x+1])
		northvfh += 1
		if h1 >= h0 {
			visible = false
			break
		}
	}
	if visible {
		vt += 1
	}
	visible = true
	southvfh := 0
	for y := ty + 1; y < len(grid); y++ {
		h1, _ := strconv.Atoi(grid[y][x : x+1])
		southvfh += 1
		if h1 >= h0 {
			visible = false
			break
		}
	}
	if visible {
		vt += 1
	}
	if vt > 0 {
		visible = true
	}

	return visible, westvfh * eastvfh * northvfh * southvfh
}
