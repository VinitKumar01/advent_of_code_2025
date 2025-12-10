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

type Point struct {
	x, y int
}

func part1(content string) {
	coordinates := strings.Split(strings.TrimSpace(content), "\n")

	var points []Point

	for _, coordinate := range coordinates {
		if coordinate == "" {
			continue
		}
		coor := strings.Split(coordinate, ",")
		x, err := strconv.Atoi(coor[0])
		if err != nil {
			fmt.Printf("Error while parsing x to int:\n%v\n", err)
			return
		}
		y, err := strconv.Atoi(coor[1])
		if err != nil {
			fmt.Printf("Error while parsing y to int:\n%v\n", err)
			return
		}
		points = append(points, Point{
			x,
			y,
		})
	}

	maximum := 0

	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			length := max(points[i].x-points[j].x, points[j].x-points[i].x) + 1
			height := max(points[i].y-points[j].y, points[j].y-points[i].y) + 1
			area := length * height
			maximum = max(maximum, area)
		}
	}

	fmt.Printf("Area: %v\n", maximum)
}

func part2(content string) {
	coordinates := strings.Split(strings.TrimSpace(content), "\n")

	var points []Point

	for _, coordinate := range coordinates {
		if coordinate == "" {
			continue
		}
		coor := strings.Split(coordinate, ",")
		x, err := strconv.Atoi(coor[0])
		if err != nil {
			fmt.Printf("Error while parsing x to int:\n%v\n", err)
			return
		}
		y, err := strconv.Atoi(coor[1])
		if err != nil {
			fmt.Printf("Error while parsing y to int:\n%v\n", err)
			return
		}
		points = append(points, Point{
			x,
			y,
		})
	}

	if len(points) == 0 {
		fmt.Printf("Area: %v\n", 0)
		return
	}

	minX, maxX := points[0].x, points[0].x
	minY, maxY := points[0].y, points[0].y
	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	type edgeV struct{ x, y1, y2 int }
	var vEdges []edgeV
	for i := range len(points) {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		if p1.x == p2.x {
			yy1 := p1.y
			yy2 := p2.y
			if yy1 > yy2 {
				yy1, yy2 = yy2, yy1
			}
			vEdges = append(vEdges, edgeV{p1.x, yy1, yy2})
		}
	}

	pointMap := make(map[[2]int]bool)
	redsByRow := make(map[int][]int)
	for _, p := range points {
		pointMap[[2]int{p.x, p.y}] = true
		redsByRow[p.y] = append(redsByRow[p.y], p.x)
	}
	for y := range redsByRow {
		a := redsByRow[y]
		var qs func(int, int)
		qs = func(l, r int) {
			if l >= r {
				return
			}
			piv := a[(l+r)/2]
			i, j := l, r
			for i <= j {
				for a[i] < piv {
					i++
				}
				for a[j] > piv {
					j--
				}
				if i <= j {
					a[i], a[j] = a[j], a[i]
					i++
					j--
				}
			}
			if l < j {
				qs(l, j)
			}
			if i < r {
				qs(i, r)
			}
		}
		qs(0, len(a)-1)
		redsByRow[y] = a
	}

	type interval struct{ a, b int }
	intervalCache := make(map[int][]interval)

	sortInts := func(a []int) {
		if len(a) < 2 {
			return
		}
		var qs func(int, int)
		qs = func(l, r int) {
			if l >= r {
				return
			}
			piv := a[(l+r)/2]
			i, j := l, r
			for i <= j {
				for a[i] < piv {
					i++
				}
				for a[j] > piv {
					j--
				}
				if i <= j {
					a[i], a[j] = a[j], a[i]
					i++
					j--
				}
			}
			if l < j {
				qs(l, j)
			}
			if i < r {
				qs(i, r)
			}
		}
		qs(0, len(a)-1)
	}

	getIntervals := func(y int) []interval {
		if itv, ok := intervalCache[y]; ok {
			return itv
		}
		var xs []int
		for _, e := range vEdges {
			if e.y1 <= y && y < e.y2 {
				xs = append(xs, e.x)
			}
		}
		sortInts(xs)
		var its []interval
		for i := 0; i+1 < len(xs); i += 2 {
			l := xs[i]
			r := xs[i+1] - 1
			if l <= r {
				its = append(its, interval{l, r})
			}
		}
		intervalCache[y] = its
		return its
	}

	countInRange := func(arr []int, l, r int) int {
		if len(arr) == 0 || l > r {
			return 0
		}
		lo := 0
		hi := len(arr)
		for lo < hi {
			m := (lo + hi) / 2
			if arr[m] < l {
				lo = m + 1
			} else {
				hi = m
			}
		}
		start := lo
		lo = 0
		hi = len(arr)
		for lo < hi {
			m := (lo + hi) / 2
			if arr[m] <= r {
				lo = m + 1
			} else {
				hi = m
			}
		}
		end := lo
		return end - start
	}

	minf := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	maxf := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maximum := 0

	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			x1, y1 := points[i].x, points[i].y
			x2, y2 := points[j].x, points[j].y
			lx := minf(x1, x2)
			rx := maxf(x1, x2)
			ly := minf(y1, y2)
			ry := maxf(y1, y2)
			valid := true
			for y := ly; y <= ry; y++ {
				its := getIntervals(y)
				curr := lx
				okRow := true
				for _, it := range its {
					if it.b < curr {
						continue
					}
					if it.a > rx {
						break
					}
					if it.a > curr {
						gapL := curr
						gapR := minf(it.a-1, rx)
						gapLen := gapR - gapL + 1
						if gapLen > 0 {
							reds := countInRange(redsByRow[y], gapL, gapR)
							if reds != gapLen {
								okRow = false
								break
							}
						}
					}
					if it.b+1 > curr {
						curr = it.b + 1
					}
					if curr > rx {
						break
					}
				}
				if okRow && curr <= rx {
					gapL := curr
					gapR := rx
					gapLen := gapR - gapL + 1
					if gapLen > 0 {
						reds := countInRange(redsByRow[y], gapL, gapR)
						if reds != gapLen {
							okRow = false
						}
					}
				}
				if !okRow {
					valid = false
					break
				}
			}
			if valid {
				length := maxf(x1-x2, x2-x1) + 1
				height := maxf(y1-y2, y2-y1) + 1
				area := length * height
				if area > maximum {
					maximum = area
				}
			}
		}
	}

	fmt.Printf("Area: %v\n", maximum)
}
