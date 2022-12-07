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

	dirSizeSum, dirToDeleteSize := solve(lines)
	fmt.Println("sum of large folder sizes:", dirSizeSum)
	fmt.Println("folder size to delete:", dirToDeleteSize)
}

type dir struct {
	name        string
	parent      *dir
	fileSizeSum int
	dirs        []*dir
}

func solve(lines []string) (int, int) {
	rootDir := dir{name: "/", parent: nil}
	curDir := &rootDir
	for _, text := range lines {
		tokens := strings.Split(text, " ")
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					if curDir.parent != nil {
						curDir = curDir.parent
					}
				} else {
					for _, d := range curDir.dirs {
						if tokens[2] == d.name {
							curDir = d
							break
						}
					}
				}
			}
		} else if tokens[0] == "dir" {
			fmt.Println("adding dir", tokens[1], "to cur dir", curDir.name)
			curDir.dirs = append(curDir.dirs, &dir{name: tokens[1], parent: curDir})
		} else {
			size, parsed := strconv.Atoi(tokens[0])
			if parsed == nil {
				curDir.fileSizeSum += size
			}
		}
	}

	updateDirSizes(&rootDir)

	printDirTree(&rootDir, 0)

	return getSmallDirSizeSum(&rootDir), findDirThatFreeEnoughSpace(&rootDir, 30000000-(70000000-rootDir.fileSizeSum), rootDir.fileSizeSum)
}

func updateDirSizes(curDir *dir) {
	dirSizeSum := 0
	for _, d := range curDir.dirs {
		updateDirSizes(d)
		dirSizeSum += d.fileSizeSum
	}
	curDir.fileSizeSum += dirSizeSum
}

func getSmallDirSizeSum(curDir *dir) int {
	acc := 0
	for _, d := range curDir.dirs {
		if d.fileSizeSum < 100000 {
			acc += d.fileSizeSum
		}
		acc += getSmallDirSizeSum(d)
	}
	return acc
}

func findDirThatFreeEnoughSpace(curDir *dir, sizeToFree int, cur int) int {
	newcur := cur
	for _, d := range curDir.dirs {
		if d.fileSizeSum < cur && d.fileSizeSum > sizeToFree {
			newcur = d.fileSizeSum
			recur := findDirThatFreeEnoughSpace(d, sizeToFree, newcur)
			if recur < newcur {
				newcur = recur
			}
		}
	}
	return newcur
}

func printDirTree(curDir *dir, depth int) {
	indent := strings.Repeat(" ", depth) + "-"
	fmt.Println(indent, curDir.name)
	fmt.Println(indent, "| size", curDir.fileSizeSum)
	for _, dir := range curDir.dirs {
		printDirTree(dir, depth+1)
	}
}
