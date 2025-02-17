package main

// import "fmt"

func order_even_first_odd_last_v1(numbers *[]int) {
	evenList := make([]int, 0, len(*numbers))
	oddList := make([]int, 0, len(*numbers))
	for _, num := range *numbers {
		if num%2 == 0 {
			evenList = append(evenList, num)
		} else {
			oddList = append(oddList, num)
		}
	}
	*numbers = append(evenList, oddList...)
}

func order_even_first_odd_last_v2(numbers *[]int) {
	i, j := 0, len(*numbers)-1
	for i < j {
		if (*numbers)[i]%2 == 0 {
			i++
		} else {
			(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
			j--
		}
	}
}

// func main() {
// 	l := []int{0, 1, 3, 4, 2, 4, 5, 1, 6, 9, 8}
// 	order_even_first_odd_last_v2(&l)
// 	fmt.Print(l)
// }
