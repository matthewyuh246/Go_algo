package main

import (
	"fmt"
	"strings"
)

func order_change_by_indexes_v1(chars []string, indexes []int) string {
	tmp := make([]string, len(chars))
	for i, index := range indexes {
		tmp[index] = chars[i]
	}
	return strings.Join(tmp, "")
}

func order_change_by_indexes_v2(chars []string, indexes []int) string {
	i, lenIndexes := 0, len(indexes)-1
	for i < lenIndexes {
		for i != indexes[i] {
			index := indexes[i]
			chars[index], chars[i] = chars[i], chars[index]
			indexes[index], indexes[i] = indexes[i], indexes[index]
		}
		i++
	}
	return strings.Join(chars, "")
}

func main() {
	w := []string{"h", "y", "n", "p", "t", "o"}
	i := []int{3, 1, 5, 0, 2, 4}

	result := order_change_by_indexes_v2(w, i)
	fmt.Println(result)
}
