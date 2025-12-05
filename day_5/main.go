package main

import (
	"fmt"
	"os"
	"sort"
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

	// part1(content)
	part2(content)
}

func part1(content string) {
	ingridients := strings.Split(content, "\n\n")

	fresh_ranges := strings.Split(strings.TrimSpace(ingridients[0]), "\n")
	available_ids := strings.Split(strings.TrimSpace(ingridients[1]), "\n")

	total := 0
	count_map := make(map[int]bool)

	for _, fresh_range := range fresh_ranges {
		range_values := strings.Split(fresh_range, "-")
		start, err := strconv.Atoi(range_values[0])
		if err != nil {
			fmt.Printf("Error while parsing start to int:\n%v\n", err)
			return
		}
		end, err := strconv.Atoi(range_values[1])
		if err != nil {
			fmt.Printf("Error while parsing end to int:\n%v\n", err)
			return
		}

		for _, id := range available_ids {
			id_int, err := strconv.Atoi(id)
			if err != nil {
				fmt.Printf("Error while parsing id to int:\n%v\n", err)
				return
			}

			if id_int >= start && id_int <= end && !count_map[id_int] {
				count_map[id_int] = true
				total += 1
			}
		}
	}

	fmt.Printf("Total: %v\n", total)
}

func part2(content string) {
	lines := strings.Split(strings.TrimSpace(content), "\n")

	type rng struct {
		start, end int64
	}

	var ranges []rng

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "-")
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing start: %v\n", err)
			return
		}
		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing end: %v\n", err)
			return
		}

		ranges = append(ranges, rng{start, end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	var merged []rng
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}

		last := &merged[len(merged)-1]

		if r.start <= last.end+1 {
			if r.end > last.end {
				last.end = r.end
			}
		} else {
			merged = append(merged, r)
		}
	}

	var total int64
	for _, r := range merged {
		total += r.end - r.start + 1
	}

	fmt.Printf("Total: %v\n", total)
}
