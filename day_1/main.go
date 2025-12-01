package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contentBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input.txt:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(contentBytes)), "\n")
	var offsets []int

	for _, line := range lines {
		if line == "" {
			continue
		}

		dir := line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}

		offset := val
		if dir == 'L' {
			offset = -val
		}
		offsets = append(offsets, offset)
	}

	fmt.Printf("Part 1: %d\n", part1(offsets))
	fmt.Printf("Part 2: %d (Hex: %X)\n", part2(offsets), part2(offsets))
}

func part1(offsets []int) int {
	pos := 50
	count := 0

	for _, offset := range offsets {
		pos += offset
		if pos%100 == 0 {
			count++
		}
	}

	return count
}

func part2(offsets []int) int {
	zeroCount := 0
	pos := 50

	for _, offset := range offsets {
		if offset >= 0 {
			zeroCount += (pos + offset) / 100
		} else {
			zeroCount -= offset / 100
			offset = offset % 100
			if pos != 0 && (pos+offset <= 0) {
				zeroCount++
			}
		}
		pos = ((pos+offset)%100 + 100) % 100
	}

	return zeroCount
}
