package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	contentBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error while reading the file:\n%v\n", err)
		return
	}
	content := string(contentBytes)
	part1(content)
	part2(content)
}

func part1(content string) {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	totalPresses := 0
	reBrackets := regexp.MustCompile(`\[(.*?)\]`)
	reParens := regexp.MustCompile(`\((.*?)\)`)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		bracketMatch := reBrackets.FindStringSubmatch(line)
		if len(bracketMatch) < 2 {
			continue
		}
		diagram := bracketMatch[1]
		targetMask := 0
		for i, c := range diagram {
			if c == '#' {
				targetMask |= (1 << i)
			}
		}

		buttonMatches := reParens.FindAllStringSubmatch(line, -1)
		buttons := []int{}
		for _, bm := range buttonMatches {
			if len(bm) < 2 {
				continue
			}
			nums := strings.Split(bm[1], ",")
			mask := 0
			for _, numStr := range nums {
				numStr = strings.TrimSpace(numStr)
				if numStr == "" {
					continue
				}
				idx, err := strconv.Atoi(numStr)
				if err != nil {
					continue
				}
				mask |= (1 << idx)
			}
			buttons = append(buttons, mask)
		}

		minPresses := findMinPresses(targetMask, buttons)
		totalPresses += minPresses
	}
	fmt.Println("Part 1:", totalPresses)
}

func findMinPresses(target int, buttons []int) int {
	numButtons := len(buttons)
	minPresses := -1
	for combo := range 1 << numButtons {
		state := 0
		count := 0
		for i := range numButtons {
			if (combo>>i)&1 == 1 {
				state ^= buttons[i]
				count++
			}
		}
		if state == target {
			if minPresses == -1 || count < minPresses {
				minPresses = count
			}
		}
	}
	if minPresses == -1 {
		return 0
	}
	return minPresses
}

type Machine struct {
	matrix  [][]float64
	targets []float64
}

func part2(content string) {
	machines := parseMachines(content)
	totalPresses := 0

	for i, m := range machines {
		res := solveMachineRobust(m)
		if res == -1 {
			fmt.Printf("Machine %d: No solution found\n", i+1)
		} else {
			totalPresses += res
		}
	}
	fmt.Println("Part 2 Answer:", totalPresses)
}

func parseMachines(content string) []Machine {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	var machines []Machine

	reParens := regexp.MustCompile(`\((.*?)\)`)
	reCurly := regexp.MustCompile(`\{(.*?)\}`)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		curlyMatch := reCurly.FindStringSubmatch(line)
		if len(curlyMatch) < 2 {
			continue
		}
		targetStrs := strings.Split(curlyMatch[1], ",")
		var targets []float64
		for _, s := range targetStrs {
			val, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
			targets = append(targets, val)
		}
		numCounters := len(targets)

		buttonMatches := reParens.FindAllStringSubmatch(line, -1)
		var cols [][]float64
		for _, bm := range buttonMatches {
			vec := make([]float64, numCounters)
			for idxStr := range strings.SplitSeq(bm[1], ",") {
				idx, err := strconv.Atoi(strings.TrimSpace(idxStr))
				if err == nil && idx < numCounters {
					vec[idx] = 1.0
				}
			}
			cols = append(cols, vec)
		}

		numButtons := len(cols)
		matrix := make([][]float64, numCounters)
		for r := range numCounters {
			matrix[r] = make([]float64, numButtons)
			for c := range numButtons {
				matrix[r][c] = cols[c][r]
			}
		}

		machines = append(machines, Machine{matrix: matrix, targets: targets})
	}
	return machines
}

func solveMachineRobust(m Machine) int {
	rows := len(m.matrix)
	if rows == 0 {
		return 0
	}
	cols := len(m.matrix[0])
	aug := make([][]float64, rows)
	for r := range rows {
		aug[r] = make([]float64, cols+1)
		copy(aug[r], m.matrix[r])
		aug[r][cols] = m.targets[r]
	}

	pivots := rref(aug)

	pivotColToRow := make(map[int]int)
	isPivot := make(map[int]bool)
	for r, c := range pivots {
		pivotColToRow[c] = r
		isPivot[c] = true
	}

	var freeCols []int
	for c := range cols {
		if !isPivot[c] {
			freeCols = append(freeCols, c)
		}
	}

	for r := range rows {
		allZero := true
		for c := range cols {
			if math.Abs(aug[r][c]) > 1e-9 {
				allZero = false
				break
			}
		}
		if allZero && math.Abs(aug[r][cols]) > 1e-9 {
			return -1
		}
	}

	minTotal := -1.0

	calcTotal := func(freeVals []int) float64 {
		total := 0.0
		for _, v := range freeVals {
			total += float64(v)
		}

		for c := range cols {
			if isPivot[c] {
				r := pivotColToRow[c]
				val := aug[r][cols]
				for i, fCol := range freeCols {
					coeff := aug[r][fCol]
					val -= coeff * float64(freeVals[i])
				}

				if val < -1e-9 {
					return -1
				}
				rounded := math.Round(val)
				if math.Abs(val-rounded) > 1e-4 {
					return -1
				}
				total += rounded
			}
		}
		return total
	}

	var recurse func(idx int, currentFree []int)
	recurse = func(idx int, currentFree []int) {
		if idx == len(freeCols) {
			res := calcTotal(currentFree)
			if res != -1 {
				if minTotal == -1 || res < minTotal {
					minTotal = res
				}
			}
			return
		}

		for val := range 200 {
			recurse(idx+1, append(currentFree, val))
		}
	}

	recurse(0, []int{})

	if minTotal == -1 {
		return -1
	}
	return int(minTotal)
}

func rref(aug [][]float64) map[int]int {
	rows := len(aug)
	cols := len(aug[0])
	pivotRow := 0
	pivots := make(map[int]int)

	for c := 0; c < cols-1 && pivotRow < rows; c++ {
		sel := pivotRow
		for i := pivotRow + 1; i < rows; i++ {
			if math.Abs(aug[i][c]) > math.Abs(aug[sel][c]) {
				sel = i
			}
		}
		if math.Abs(aug[sel][c]) < 1e-9 {
			continue
		}

		aug[pivotRow], aug[sel] = aug[sel], aug[pivotRow]
		pivots[pivotRow] = c

		div := aug[pivotRow][c]
		for j := c; j < cols; j++ {
			aug[pivotRow][j] /= div
		}

		for i := range rows {
			if i != pivotRow {
				factor := aug[i][c]
				for j := c; j < cols; j++ {
					aug[i][j] -= factor * aug[pivotRow][j]
				}
			}
		}
		pivotRow++
	}
	return pivots
}
