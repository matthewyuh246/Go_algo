package main

import "fmt"

type Move struct {
	Disk int
	Src  string
	Dest string
}

func hanoi(disk int, src, dest, support string) {
	if disk < 1 {
		return
	}

	hanoi(disk-1, src, support, dest)
	fmt.Printf("move %d from %s to %s\n", disk, src, dest)
	hanoi(disk-1, support, dest, src)
}

func getHanoiMovement(disk int, src, dest, support string) []Move {
	var result []Move

	var hanoi func(disk int, src, dest, support string)
	hanoi = func(disk int, src, dest, support string) {
		if disk < 1 {
			return
		}

		hanoi(disk-1, src, support, dest)
		result = append(result, Move{Disk: disk, Src: src, Dest: dest})
		hanoi(disk-1, support, dest, src)
	}

	hanoi(disk, src, dest, support)
	return result
}

// func main() {
// 	hanoi(3, "A", "C", "B")
// 	for _, r := range getHanoiMovement(4, "A", "C", "B") {
// 		fmt.Println(r)
// 	}
// }
