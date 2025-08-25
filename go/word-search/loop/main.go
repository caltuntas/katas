package main

import (
	"fmt"
	"slices"
)

type Location struct {
	Row    int
	Column int
}

func (loc Location) GetNeighbors() []Location {
	return []Location{{loc.Row - 1, loc.Column}, {loc.Row + 1, loc.Column}, {loc.Row, loc.Column - 1}, {loc.Row, loc.Column + 1}}
}

func isValid(grid [][]byte, currentLoc Location, locs ...Location) bool {
	rowCount := len(grid)
	colCount := len(grid[0])
	for _, loc := range locs {
		if currentLoc.Row == loc.Row && currentLoc.Column == loc.Column {
			return false
		}
	}
	return currentLoc.Row >= 0 && currentLoc.Column >= 0 && currentLoc.Row < rowCount && currentLoc.Column < colCount
}

func traverse(grid [][]byte, loc Location) []string {
	result := make([]string,0)
	p0 := loc
	for _, p1 := range p0.GetNeighbors() {
		if isValid(grid, p1, p0) {
			for _, p2 := range p1.GetNeighbors() {
				if isValid(grid, p2, p0, p1) {
					for _, p3 := range p2.GetNeighbors() {
						if isValid(grid, p3, p0, p1, p2) {
							fmt.Print(p0)
							fmt.Print(p1)
							fmt.Print(p2)
							fmt.Print(p3)
							str := []byte{
								grid[p0.Row][p0.Column],
								grid[p1.Row][p1.Column],
								grid[p2.Row][p2.Column],
								grid[p3.Row][p3.Column],
							}
							fmt.Println(string(str))
							result = append(result, string(str))
						}
					}
				}
			}
		}
	}
	return result
}

func find(grid [][]byte, word string) bool {
	allResults :=make([]string,0)
	rows := len(grid)
	for i := 0; i < rows; i++ {
		cols := len(grid[i])
		for j := 0; j < cols; j++ {
			result := traverse(grid, Location{i, j})
			allResults = append(allResults, result...)
		}
	}
	return slices.Contains(allResults, word)
}

func main() {
	tests := []struct {
		grid   [][]byte
		word   string
		result bool
	}{
		{
			[][]byte{
				{'T', 'E', 'E'},
				{'S', 'G', 'K'},
				{'T', 'E', 'L'},
			},
			"GEEK",
			true,
		},
	}

	for i, c := range tests {
		res := find(c.grid, c.word)
		if res != c.result {
			fmt.Printf("test case failed %d\n", i)
		}
	}
}
