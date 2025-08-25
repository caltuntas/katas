package main

import (
	"fmt"
	"slices"
)

type Location struct {
	Row    int
	Column int
}

func findLocations(grid [][]byte, chr byte) []Location {
	locations := make([]Location, 0)
	rows := len(grid)
	for i := 0; i < rows; i++ {
		cols := len(grid[i])
		for j := 0; j < cols; j++ {
			if chr == grid[i][j] {
				locations = append(locations, Location{Row: i, Column: j})
			}
		}
	}
	return locations
}

func find(grid [][]byte, word string) bool {
	findNext := func(curRow, curCol int, nextChar byte) []Location {
		rowCount := len(grid)
		colCount := len(grid[0])
		result := make([]Location, 0)
		for _, p := range [4][2]int{{curRow - 1, curCol}, {curRow + 1, curCol}, {curRow, curCol - 1}, {curRow, curCol + 1}} {
			row, col := p[0], p[1]
			if row >= 0 && col >= 0 && row < rowCount && col < colCount {
				if grid[row][col] == nextChar {
					result = append(result, Location{Row: row, Column: col})
				}
			}
		}
		return result
	}

	curChar := 0
	stack := make([]Location, 0)
	visited := make([]Location, 0)
	startLocations := findLocations(grid, word[curChar])

	for _, startLoc := range startLocations {
		curChar++
		if curChar >= len(word) {
			break
		}
		visited = append(visited, startLoc)
		locations := findNext(startLoc.Row, startLoc.Column, word[curChar])
		stack = append(stack, locations...)
		for len(stack) > 0 {
			curChar++
			if curChar >= len(word) {
				return true
			}
			l := len(stack)
			loc := stack[l-1]
			stack = stack[:l-1]
			locations = findNext(loc.Row, loc.Column, word[curChar])
			visited = append(visited, loc)
			if len(locations) <= 0 {
				curChar--
			}
			for _, ll := range locations {
				if !slices.Contains(visited, ll) && !slices.Contains(stack, ll) {
					stack = append(stack, ll)
				}
			}
		}
		curChar = 0
		visited = nil
	}

	return false
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
