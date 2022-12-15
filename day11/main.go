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

	monkeyBusiness := solve(lines)
	fmt.Println("monkey business:", monkeyBusiness)
}

type monkey struct {
	id                  int
	items               *[]int
	oper                operation
	testdiv             int
	trueCase, falseCase int
	itemInspCnt         int
}

type operation struct {
	lefth, righth string
	operator      string
}

func solve(lines []string) int {
	var monkeys []*monkey
	monk := &monkey{}
	monkeys = append(monkeys, monk)
	for i := 0; i < len(lines); i++ {
		text := strings.TrimSpace(lines[i])
		if text == "" {
			monk = &monkey{}
			monkeys = append(monkeys, monk)
			continue
		}
		if strings.HasPrefix(text, "Monkey") {
			var id int
			fmt.Sscanf(text, "Monkey %d:", &id)
			monk.id = id
		} else {
			sp := strings.Split(text, ": ")
			if sp[0] == "Starting items" {
				its := strings.Split(sp[1], ", ")
				var items []int
				for _, it := range its {
					num, _ := strconv.Atoi(it)
					items = append(items, num)
				}
				monk.items = &items
			}
			if sp[0] == "Operation" {
				fmt.Sscanf(sp[1], "new = %s %s %s", &monk.oper.lefth, &monk.oper.operator, &monk.oper.righth)
			}
			if sp[0] == "Test" {
				fmt.Sscanf(sp[1], "divisible by %d", &monk.testdiv)
			}
			if sp[0] == "If true" {
				fmt.Sscanf(sp[1], "throw to monkey %d", &monk.trueCase)
			}
			if sp[0] == "If false" {
				fmt.Sscanf(sp[1], "throw to monkey %d", &monk.falseCase)
			}
		}
	}

	modulo := 1
	for _, monk := range monkeys {
		modulo *= monk.testdiv
	}

	for i := 0; i < 10000; i++ {
		for _, monk := range monkeys {
			for _, worry := range *monk.items {
				var lh, rh int
				if monk.oper.lefth == "old" {
					lh = worry
				} else {
					val, _ := strconv.Atoi(monk.oper.lefth)
					lh = val
				}
				if monk.oper.righth == "old" {
					rh = worry
				} else {
					val, _ := strconv.Atoi(monk.oper.righth)
					rh = val
				}
				switch monk.oper.operator {
				case "*":
					worry = lh * rh
					break
				case "+":
					worry = lh + rh
					break
				default:
					break
				}
				//worry /= 3
				worry = worry % modulo

				var newMonk *monkey
				if worry%monk.testdiv == 0 {
					newMonk = monkeys[monk.trueCase]
				} else {
					newMonk = monkeys[monk.falseCase]
				}
				*newMonk.items = append(*newMonk.items, worry)
				monk.itemInspCnt += 1
			}
			*monk.items = make([]int, 0)
		}
	}

	m1, m2 := -1, -1
	for i, monk := range monkeys {
		fmt.Println("monkey", monk.id, "inspected", monk.itemInspCnt, "items")
		if m1 == -1 {
			m1 = i
		} else if m2 == -1 {
			m2 = i
		} else {
			if monk.itemInspCnt > monkeys[m1].itemInspCnt {
				if monkeys[m1].itemInspCnt > monkeys[m2].itemInspCnt {
					m2 = m1
				}
				m1 = i
			} else if monk.itemInspCnt > monkeys[m2].itemInspCnt {
				m2 = i
			}
		}

	}

	return monkeys[m1].itemInspCnt * monkeys[m2].itemInspCnt
}
