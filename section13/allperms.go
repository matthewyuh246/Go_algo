package main

import "fmt"

func allPerms(elements []int) [][]int {
	if len(elements) <= 1 {
		cpy := make([]int, len(elements))
		copy(cpy, elements)
		return [][]int{cpy}
	}

	result := [][]int{}
	firstElem := elements[0]
	tail := elements[1:]
	permsTail := allPerms(tail)

	for _, perm := range permsTail {
		for i := 0; i <= len(perm); i++ {
			newPerm := make([]int, 0, len(perm)+1)
			newPerm = append(newPerm, perm[:i]...)
			newPerm = append(newPerm, firstElem)
			newPerm = append(newPerm, perm[i:]...)
			result = append(result, newPerm)
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	perms := allPerms(nums)
	for _, p := range perms {
		fmt.Println(p)
	}
}
