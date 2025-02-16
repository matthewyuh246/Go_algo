package main

// import "fmt"

func isPalindrome(strings string) bool {
	lenStrings := len(strings)
	if lenStrings == 0 {
		return false
	}
	if lenStrings == 1 {
		return true
	}

	start, end := 0, lenStrings-1
	for start < end {
		if strings[start] != strings[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// func main() {
// 	fmt.Println(isPalindrome("racecar"))
// 	fmt.Println(isPalindrome("test"))
// }
