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

func traverse(grid [][]byte, loc Location, word string, path []Location) []string {
	result := make([]string, 0)
	if len(path) == len(word) {
		str := make([]byte, len(path))
		for i, loc := range path {
			str[i] = grid[loc.Row][loc.Column]
		}
		result = append(result, string(str))
		return result
	}
	for _, p1 := range loc.GetNeighbors() {
		if isValid(grid, p1, path...) {
			path = append(path, p1)
			r := traverse(grid, p1, word, path)
			result = append(result, r...)
			path = path[:len(path)-1]
		}
	}
	return result
}

func find(grid [][]byte, word string) bool {
	allResults := make([]string, 0)
	rows := len(grid)
	for i := 0; i < rows; i++ {
		cols := len(grid[i])
		for j := 0; j < cols; j++ {
			result := traverse(grid, Location{i, j}, word, []Location{})
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
		{
			[][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			"DOG",
			false,
		},
		{
			[][]byte{
				{'T', 'E', 'R'},
				{'K', 'G', 'K'},
				{'E', 'E', 'L'},
			},
			"GEEK",
			true,
		},
		{
			[][]byte{
				{'A', 'A', 'A'},
				{'B', 'A', 'C'},
			},
			"AAA",
			true,
		},
		{
			[][]byte{
				{'C', 'A', 'T'},
				{'X', 'Y', 'Z'},
				{'D', 'O', 'G'},
			},
			"TAC",
			true,
		},
		{
			[][]byte{
				{'H', 'E', 'L', 'L', 'O'},
			},
			"HELLO",
			true,
		},
		{
			[][]byte{
				{'H'},
				{'E'},
				{'L'},
				{'L'},
				{'O'},
			},
			"HELLO",
			true,
		},
		{
			[][]byte{
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
			},
			"AAAAA",
			true,
		},
		{
			[][]byte{
				{'C', 'A', 'T'},
				{'D', 'O', 'X'},
				{'G', 'Y', 'Z'},
			},
			"DOG",
			false,
		},
		{
			[][]byte{
				{'C', 'A', 'R'},
				{'D', 'O', 'G'},
				{'H', 'A', 'T'},
			},
			"CAT",
			false,
		},
		{
			[][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			"ABCDE",
			false,
		},
		{
			[][]byte{
				{'S', 'U', 'N'},
				{'S', 'U', 'N'},
				{'S', 'U', 'N'},
			},
			"SUN",
			true,
		},
		{
			[][]byte{
				{'A', 'B'},
				{'A', 'C'},
			},
			"AC",
			true,
		},
		{
			[][]byte{
				{'C', 'A', 'R'},
				{'C', 'A', 'T'},
			},
			"CAT",
			true,
		},
		{
			[][]byte{
				{'A', 'B'},
				{'B', 'A'},
			},
			"ABBA",
			false,
		},
		{
			[][]byte{
				{'A', 'B'},
				{'B', 'A'},
			},
			"ABA",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
			true,
		},
		{
			[][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			"AAAAA",
			false,
		},
		{
			[][]byte{
				{'A', 'B'},
			},
			"ABA",
			false,
		},
		{
			[][]byte{
				{'S', 'E', 'E'},
				{'E', 'S', 'E'},
				{'E', 'E', 'S'},
			},
			"SEE",
			true,
		},
		{
			[][]byte{
				{'A', 'A', 'A'},
			},
			"AAAA",
			false,
		},
		{
			[][]byte{
				{'A', 'B'},
				{'B', 'A'},
			},
			"ABABA",
			false,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCB",
			false,
		},
		{
			[][]byte{
				{'A', 'B', 'A'},
				{'B', 'A', 'B'},
				{'A', 'B', 'A'},
			},
			"ABABA",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C'},
				{'B', 'C', 'D'},
			},
			"ABC",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C'},
				{'A', 'D', 'E'},
			},
			"ADE",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C'},
				{'B', 'A', 'D'},
				{'C', 'D', 'A'},
			},
			"ABABA",
			false,
		},
		{
			[][]byte{
				{'A', 'A', 'A'},
				{'A', 'B', 'A'},
				{'A', 'A', 'A'},
			},
			"ABAA",
			true,
		},
		{
			[][]byte{
				{'A', 'B'},
				{'C', 'A'},
			},
			"ABA",
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
