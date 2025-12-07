package main

import (
	"fmt"
	"os"
	"strings"
)

type Beam struct {
	x, y int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")

	fmt.Println("Part 1 - Total splits:", part1(lines))
	fmt.Println("Part 2 - Total timelines:", part2(lines))
}

func part1(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	var startX int
	for i, c := range grid[0] {
		if c == 'S' {
			startX = i
			break
		}
	}

	queue := []Beam{{startX, 1}}
	splits := 0
	visited := make(map[[2]int]bool)

	for len(queue) > 0 {
		beam := queue[0]
		queue = queue[1:]

		if beam.y >= height {
			continue
		}
		if beam.x < 0 || beam.x >= width {
			continue
		}

		pos := [2]int{beam.x, beam.y}
		if visited[pos] {
			continue
		}
		visited[pos] = true

		switch grid[beam.y][beam.x] {
		case '.':
			queue = append(queue, Beam{beam.x, beam.y + 1})

		case '^':
			splits++
			queue = append(queue, Beam{beam.x - 1, beam.y + 1})
			queue = append(queue, Beam{beam.x + 1, beam.y + 1})
		default:
			queue = append(queue, Beam{beam.x, beam.y + 1})
		}
	}

	return splits
}

func part2(grid []string) int {
	height := len(grid)
	if height == 0 {
		return 0
	}
	sx, sy := -1, -1
	for y := range height {
		for x := range grid[y] {
			if grid[y][x] == 'S' || grid[y][x] == 's' {
				sx, sy = x, y
				break
			}
		}
		if sx != -1 {
			break
		}
	}
	if sx == -1 {
		return 0
	}
	memo := make(map[[2]int]int)
	var paths func(x, y int) int
	paths = func(x, y int) int {
		if x < 0 {
			return 1
		}
		if y >= height {
			return 1
		}
		if x >= len(grid[y]) {
			return 1
		}
		key := [2]int{x, y}
		if v, ok := memo[key]; ok {
			return v
		}
		c := grid[y][x]
		var res int
		if c == '^' {
			res = paths(x-1, y+1) + paths(x+1, y+1)
		} else {
			res = paths(x, y+1)
		}
		memo[key] = res
		return res
	}
	return paths(sx, sy+1)
}
