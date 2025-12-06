package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content_byte, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error while reading the file:\n%v\n", err)
		return
	}

	content := string(content_byte)
	part1(content)
	part2(content)
}

func part1(content string) {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	operations := strings.Split(lines[len(lines)-1], " ")
	var reduced_operations []string
	for _, operation := range operations {
		if strings.TrimSpace(operation) != "" {
			reduced_operations = append(reduced_operations, strings.TrimSpace(operation))
		}
	}

	var numbers [][]int
	for i, line := range lines {
		if i == len(lines)-1 {
			continue
		}
		nums := strings.Split(line, " ")
		var ns []int
		for _, num := range nums {
			if strings.TrimSpace(num) != "" {
				n, err := strconv.Atoi(strings.TrimSpace(num))
				if err != nil {
					fmt.Printf("Error while parsing n to int:\n%v\n", err)
					return
				}
				ns = append(ns, n)
			}
		}
		numbers = append(numbers, ns)
	}

	total := 0

	for i, operator := range reduced_operations {
		result_add := 0
		result_mul := 1
		for _, num := range numbers {
			if operator != "*" {
				result_mul = 0
			}

			switch operator {
			case "+":
				result_add += num[i]
			case "*":
				result_mul *= num[i]
			}
		}
		total = total + result_mul + result_add
	}

	fmt.Printf("total:%v\n", total)
}

func part2(content string) {
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	if len(lines) < 2 {
		fmt.Printf("total:0\n")
		return
	}

	opLine := lines[len(lines)-1]
	digitLines := lines[:len(lines)-1]

	width := len(opLine)
	for _, l := range digitLines {
		if len(l) > width {
			width = len(l)
		}
	}

	for i, l := range digitLines {
		if len(l) < width {
			digitLines[i] = l + strings.Repeat(" ", width-len(l))
		}
	}
	if len(opLine) < width {
		opLine = opLine + strings.Repeat(" ", width-len(opLine))
	}

	height := len(digitLines)
	blankCols := make([]bool, width)

	for c := range width {
		isBlank := true
		for r := range height {
			if digitLines[r][c] != ' ' {
				isBlank = false
				break
			}
		}
		if isBlank && opLine[c] != ' ' {
			isBlank = false
		}
		blankCols[c] = isBlank
	}

	total := 0
	c := 0

	for c < width {
		if blankCols[c] {
			c++
			continue
		}

		start := c
		for c < width && !blankCols[c] {
			c++
		}
		end := c - 1

		var op byte
		for cc := start; cc <= end; cc++ {
			if opLine[cc] == '+' || opLine[cc] == '*' {
				op = opLine[cc]
				break
			}
		}
		if op == 0 {
			fmt.Println("Error: no operator found in problem block")
			return
		}

		var nums []int
		for cc := end; cc >= start; cc-- {
			var sb strings.Builder
			for r := range height {
				ch := digitLines[r][cc]
				if ch != ' ' {
					sb.WriteByte(ch)
				}
			}
			s := sb.String()
			if strings.TrimSpace(s) == "" {
				continue
			}
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Error parsing number %q: %v\n", s, err)
				return
			}
			nums = append(nums, n)
		}

		if len(nums) == 0 {
			continue
		}

		res := 0
		if op == '+' {
			for _, n := range nums {
				res += n
			}
		} else {
			res = 1
			for _, n := range nums {
				res *= n
			}
		}

		total += res
	}

	fmt.Printf("total:%v\n", total)
}
