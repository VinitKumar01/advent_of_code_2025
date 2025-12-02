package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var invalid_ids []int

	content_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading the file:\n%v\n", err)
		return
	}

	content := string(content_bytes)

	ranges := strings.Split(content, ",")

	part1(ranges, invalid_ids)
	part2(ranges, invalid_ids)
}

func part1(ranges []string, invalid_ids []int) {
	for _, rng := range ranges {
		rng = strings.TrimSpace(rng)

		r := strings.Split(rng, "-")
		start := r[0]
		end := r[1]

		start_int, err := strconv.Atoi(start)
		if err != nil {
			fmt.Printf("Error converting start to int:\n%v\n", err)
			return
		}
		end_int, err := strconv.Atoi(end)
		if err != nil {
			fmt.Printf("Error converting end to int:\n%v\n", err)
			return
		}

		for i := start_int; i <= end_int; i++ {
			numStr := strconv.Itoa(i)
			length := len(numStr)
			dup_size := length / 2

			first_half := int(i / int(math.Pow10(dup_size)))
			second_half := int(i % int(math.Pow10(dup_size)))

			if first_half == second_half {
				invalid_ids = append(invalid_ids, i)
			}
		}
	}

	total := 0
	for _, n := range invalid_ids {
		total += n
	}

	fmt.Printf("Total is: %v\n", total)
}

func part2(ranges []string, invalid_ids []int) {
	for _, rng := range ranges {
		rng = strings.TrimSpace(rng)

		r := strings.Split(rng, "-")
		start := r[0]
		end := r[1]

		start_int, err := strconv.Atoi(start)
		if err != nil {
			fmt.Printf("Error converting start to int:\n%v\n", err)
			return
		}
		end_int, err := strconv.Atoi(end)
		if err != nil {
			fmt.Printf("Error converting end to int:\n%v\n", err)
			return
		}

		for i := start_int; i <= end_int; i++ {
			numStr := strconv.Itoa(i)
			n := len(numStr)
			for j := 1; j <= n/2; j++ {
				if n%j == 0 {
					chunk := numStr[:j]
					repeated := strings.Repeat(chunk, n/j)
					if repeated == numStr {
						invalid_ids = append(invalid_ids, i)
						break
					}
				}
			}
		}
	}

	total := 0
	for _, n := range invalid_ids {
		total += n
	}

	fmt.Printf("Total is: %v\n", total)
}
