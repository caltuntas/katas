package main

import "fmt"

type Items struct {
	List []int
}

var s = []int{3, 34, 4, 12, 5, 2}
var target = 9

var result []Items

func combination(idx int, data []int, level int) {
	if len(data) == level {
		items := Items{}
		items.List = append(items.List, data...)
		result = append(result, items)
		return
	}

	for i := idx; i < len(s); i++ {
		data = append(data, s[i])
		combination(i+1, data, level)
		data = data[:len(data)-1]
	}
}

func sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func main() {
	result = make([]Items, 0)
	for i := 1; i <= len(s); i++ {
		combination(0, []int{}, i)
	}

	for _, item := range result {
		total := sum(item.List)
		if total == target {
			fmt.Printf("Combination %s, sum = %d\n", fmt.Sprint(item.List), total)
		}
	}
}
