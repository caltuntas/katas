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

func traverse(grid [][]byte, loc Location, word string, path []Location, result *[][]Location) {
	if word[len(path)-1] != grid[loc.Row][loc.Column] {
		return
	}
	if len(path) == len(word) {
		c := make([]Location, len(path))
		copy(c, path)
		*result = append(*result, c)
		return
	}
	for _, p1 := range loc.GetNeighbors() {
		if isValid(grid, p1, path...) {
			printLocation(path, grid)
			path = append(path, p1)
			traverse(grid, p1, word, path, result)
			path = path[:len(path)-1]
		}
	}
}

func pathToString(grid [][]byte, paths [][]Location) []string {
	result := make([]string, len(paths))
	for j, p := range paths {
		chars := make([]byte, len(p))
		for i, l := range p {
			chars[i] = grid[l.Row][l.Column]
		}
		result[j] = string(chars)
	}
	return result
}

func printLocations(grid [][]byte, paths [][]Location) {
	for _, path := range paths {
		printLocation(path, grid)
		fmt.Println()
	}
}

func printLocation(path []Location, grid [][]byte) {
	for _, loc := range path {
		fmt.Print(loc)
	}
	for _, loc := range path {
		fmt.Printf("%c", grid[loc.Row][loc.Column])
	}
}

func printDot(grid [][]byte, paths [][]Location) {
	edges := make(map[string]byte)
	nodes := make(map[string]string)
	labels := make(map[string]string)
	counter := make(map[string]int)
	getName := func(current Location, parent string) string {
		chr := string(grid[current.Row][current.Column])
		str := fmt.Sprintf("%s_%d%d", chr, current.Row, current.Column)
		oldStr := str
		nodeKey := str + "-" + parent
		_, ok := nodes[nodeKey]
		if !ok {
			if counter[str] > 0 {
				str = fmt.Sprintf("%s_%d", str, counter[str])
			}
			nodes[nodeKey] = str
			counter[oldStr] += 1
			labels[str] = chr
		}
		return nodes[nodeKey]
	}
	for _, p := range paths {
		var parent string
		for _, l := range p {
			name := getName(l, parent)
			if parent != "" {
				edge := fmt.Sprintf("  %s->%s", parent, name)
				edges[edge] = 1
			}
			parent = name
		}
	}

	fmt.Println("digraph G {")
	fmt.Println("  node [shape=circle];")
	for e := range edges {
		fmt.Println(e)
	}
	for l := range labels {
		fmt.Printf("  %s [label=\"%s\"];\n", l, labels[l])
	}
	fmt.Println("}")
}

func find(grid [][]byte, word string) bool {
	allResults := make([]string, 0)
	rows := len(grid)
	for i := 0; i < rows; i++ {
		cols := len(grid[i])
		for j := 0; j < cols; j++ {
			if grid[i][j] == word[0] {
				result := make([][]Location, 0)
				fmt.Printf("Starting point is {%d,%d}=%c\n",i,j,grid[i][j])
				traverse(grid, Location{i, j}, word, []Location{{i, j}}, &result)
				str := pathToString(grid, result)
				printLocations(grid, result)
				printDot(grid, result)
				allResults = append(allResults, str...)
			}
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
				{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'},
				{'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K'},
				{'C', 'D', 'E', 'F', 'A', 'H', 'I', 'J', 'K', 'L'},
				{'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M'},
				{'E', 'F', 'G', 'H', 'I', 'A', 'K', 'L', 'M', 'N'},
				{'F', 'G', 'H', 'I', 'J', 'K', 'A', 'M', 'N', 'O'},
				{'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'A', 'A'},
				{'H', 'I', 'J', 'K', 'A', 'M', 'N', 'O', 'P', 'Q'},
				{'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R'},
				{'J', 'K', 'L', 'M', 'N', 'O', 'A', 'Q', 'R', 'S'},
				{'A', 'B', 'C', 'D', 'E', 'A', 'G', 'H', 'I', 'J'},
				{'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K'},
				{'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L'},
				{'D', 'E', 'F', 'A', 'H', 'I', 'J', 'K', 'A', 'M'},
				{'E', 'F', 'A', 'H', 'I', 'J', 'A', 'L', 'M', 'N'},
				{'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O'},
				{'G', 'H', 'I', 'J', 'A', 'L', 'M', 'N', 'O', 'P'},
				{'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q'},
				{'I', 'J', 'K', 'A', 'M', 'N', 'O', 'P', 'Q', 'R'},
				{'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S'},
			},
			"ABDEFGHMN",
			false,
		},
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

