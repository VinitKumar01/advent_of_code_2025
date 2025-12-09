package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	contentByte, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error while reading the file:\n%b\n", err)
		return
	}

	content := string(contentByte)
	part1(content)
	part2(content)
}

type Point struct {
	x, y, z int64
}

type Edge struct {
	u, v int
	d2   int64
}

type DSU struct {
	parent []int
	size   []int
	count  int
}

func newDSU(n int) *DSU {
	p := make([]int, n)
	sz := make([]int, n)
	for i := range n {
		p[i] = i
		sz[i] = 1
	}
	return &DSU{
		parent: p,
		size:   sz,
		count:  n,
	}
}

func (d *DSU) find(x int) int {
	for x != d.parent[x] {
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
}

func (d *DSU) union(a, b int) bool {
	ra := d.find(a)
	rb := d.find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	d.count--
	return true
}

func part1(content string) {
	points := parsePoints(content)
	n := len(points)
	if n == 0 {
		fmt.Println("0")
		return
	}

	edges := buildEdges(points)
	sort.Slice(edges, func(i, j int) bool { return edges[i].d2 < edges[j].d2 })

	dsu := newDSU(n)
	maxConnections := 1000
	maxConnections = min(len(edges), maxConnections)
	for i := range maxConnections {
		e := edges[i]
		dsu.union(e.u, e.v)
	}

	compSizes := make(map[int]int)
	for i := range n {
		root := dsu.find(i)
		compSizes[root]++
	}

	sizes := make([]int, 0, len(compSizes))
	for _, sz := range compSizes {
		sizes = append(sizes, sz)
	}

	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	if len(sizes) < 3 {
		fmt.Println("0")
		return
	}

	result := sizes[0] * sizes[1] * sizes[2]
	fmt.Println(result)
}

func part2(content string) {
	points := parsePoints(content)
	n := len(points)
	if n == 0 {
		fmt.Println("0")
		return
	}

	edges := buildEdges(points)
	sort.Slice(edges, func(i, j int) bool { return edges[i].d2 < edges[j].d2 })

	dsu := newDSU(n)
	var lastU, lastV int

	for _, e := range edges {
		if dsu.union(e.u, e.v) {
			lastU, lastV = e.u, e.v
			if dsu.count == 1 {
				break
			}
		}
	}

	result := points[lastU].x * points[lastV].x
	fmt.Println(result)
}

func parsePoints(content string) []Point {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	points := make([]Point, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		y, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		z, err3 := strconv.ParseInt(strings.TrimSpace(parts[2]), 10, 64)
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}
		points = append(points, Point{x, y, z})
	}
	return points
}

func buildEdges(points []Point) []Edge {
	n := len(points)
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			d2 := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{u: i, v: j, d2: d2})
		}
	}
	return edges
}
