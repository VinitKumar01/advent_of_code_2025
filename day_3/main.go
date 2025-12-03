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
	total := 0

	part1(content, total)
	part2(content)
}

func part1(content string, total int) {
	for bank := range strings.SplitSeq(content, "\n") {
		if bank == "" {
			continue
		}

		maximum := 0

		for i := range len(bank) - 1 {
			for j := len(bank) - 1; i < j; j-- {
				first_digit, err := strconv.Atoi(bank[i : i+1])
				if err != nil {
					fmt.Printf("Error parsing i to int:\n%v\n", err)
					return
				}
				second_digit, err := strconv.Atoi(bank[j : j+1])
				if err != nil {
					fmt.Printf("Error parsing j to int:\n%v\n", err)
					return
				}

				full_digit := (first_digit * 10) + second_digit

				if full_digit > maximum {
					maximum = full_digit
				}
			}
		}

		total += maximum

	}

	fmt.Printf("Total: %v\n", total)
}

func part2(content string) {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	var total int64

	const want = 12

	for _, bank := range lines {
		bank = strings.TrimSpace(bank)
		if bank == "" {
			continue
		}

		if len(bank) < want {
			fmt.Printf("Bank too short (len=%d): %q\n", len(bank), bank)
			continue
		}

		k := len(bank) - want

		reduced := maximizeK(bank, k)

		val, err := strconv.ParseInt(reduced, 10, 64)
		if err != nil {
			fmt.Printf("Error converting %q to int: %v\n", reduced, err)
			continue
		}

		total += val
	}

	fmt.Printf("Total output joltage: %v\n", total)
}

func maximizeK(s string, k int) string {
	stack := make([]rune, 0, len(s))

	for _, ch := range s {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] < ch {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, ch)
	}

	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	return string(stack)
}
