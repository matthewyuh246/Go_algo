package main

import "fmt"

func deleteDuplicateV1(numbers *[]int) {
	tmp := []int{}
	for _, num := range *numbers {
		found := false
		for _, v := range tmp {
			if v == num {
				found = true
				break
			}
		}
		if !found {
			tmp = append(tmp, num)
		}
	}
	*numbers = tmp
}

func deleteDuplicateV3(numbers *[]int) {
	for i := len(*numbers) - 1; i > 0; i-- {
		if (*numbers)[i] == (*numbers)[i-1] {
			*numbers = append((*numbers)[:i], (*numbers)[i+1:]...)
		}
	}
}

func main() {
	numbers := []int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15}
	deleteDuplicateV3(&numbers)
	fmt.Println(numbers)
}
