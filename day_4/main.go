package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content_byte, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error while reading the file:\n%v\n", err)
		return
	}

	content := string(content_byte)

	grids := strings.Split(content, "\n")
	part1(grids)
	part2(grids)
}

func part1(grids []string) {
	total_bundles := 0

	bundle_map := make(map[string]bool)

	for i, grid := range grids {
		for j, c := range grid {
			coordinate := fmt.Sprintf("%v:%v", i, j)
			if c == '@' {
				bundle_map[coordinate] = true
			} else {
				bundle_map[coordinate] = false
			}
		}
	}

	for i, grid := range grids {
		if grid == "" {
			continue
		}
		for j, c := range grid {
			paper_bundles := 0

			if c == '@' {
				var coordinates []string

				coordinates = append(coordinates, fmt.Sprintf("%v:%v", i-1, j-1))
				coordinates = append(coordinates, fmt.Sprintf("%v:%v", i-1, j))
				coordinates = append(coordinates, fmt.Sprintf("%v:%v", i-1, j+1))
				coordinates = append(coordinates, fmt.Sprintf("%v:%v", i, j-1))
				coordinates = append(coordinates, fmt.Sprintf("%v:%v", i, j+1))
				coordinates = append(coordinates, fmt.Sprintf("%v:%v", i+1, j-1))
				coordinates = append(coordinates, fmt.Sprintf("%v:%v", i+1, j))
				coordinates = append(coordinates, fmt.Sprintf("%v:%v", i+1, j+1))

				for _, coordinate := range coordinates {
					if bundle_map[coordinate] {
						paper_bundles += 1
					} else {
						continue
					}
				}
			} else {
				continue
			}

			if paper_bundles < 4 {
				total_bundles += 1
			}
		}
	}

	fmt.Printf("Total bundles: %v\n", total_bundles)
}

func part2(grids []string) {
	total_bundles := 0

	bundle_map := make(map[string]bool)

	for i, grid := range grids {
		for j, c := range grid {
			coordinate := fmt.Sprintf("%v:%v", i, j)
			if c == '@' {
				bundle_map[coordinate] = true
			} else {
				bundle_map[coordinate] = false
			}
		}
	}

	replacements := 1

	for replacements != 0 {
		replacements = 0

		for i, grid := range grids {
			if grid == "" {
				continue
			}
			for j, c := range grid {
				paper_bundles := 0

				if c == '@' {
					var coordinates []string

					coordinates = append(coordinates, fmt.Sprintf("%v:%v", i-1, j-1))
					coordinates = append(coordinates, fmt.Sprintf("%v:%v", i-1, j))
					coordinates = append(coordinates, fmt.Sprintf("%v:%v", i-1, j+1))
					coordinates = append(coordinates, fmt.Sprintf("%v:%v", i, j-1))
					coordinates = append(coordinates, fmt.Sprintf("%v:%v", i, j+1))
					coordinates = append(coordinates, fmt.Sprintf("%v:%v", i+1, j-1))
					coordinates = append(coordinates, fmt.Sprintf("%v:%v", i+1, j))
					coordinates = append(coordinates, fmt.Sprintf("%v:%v", i+1, j+1))

					for _, coordinate := range coordinates {
						if bundle_map[coordinate] {
							paper_bundles += 1
						} else {
							continue
						}
					}
				} else {
					continue
				}

				if paper_bundles < 4 {
					total_bundles += 1
					coordinate := fmt.Sprintf("%v:%v", i, j)
					bundle_map[coordinate] = false
					grids[i] = grids[i][:j] + "." + grids[i][j+1:]
					replacements += 1
				}
			}
		}
	}

	fmt.Printf("Total bundles: %v\n", total_bundles)
}
