package main

import "fmt"

func goFindPalindrome(s string, left, right int) <-chan string {
	ch := make(chan string)
	go func() {
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}

			ch <- s[left : right+1]
			left--
			right++
		}
		close(ch)
	}()
	return ch
}

func goFindAllPalindrome(s string) <-chan string {
	ch := make(chan string)
	go func() {
		n := len(s)
		if n == 0 {
			close(ch)
			return
		}

		if n == 1 {
			ch <- s
			close(ch)
			return
		}

		for i := 1; i <= n-2; i++ {
			for pal := range goFindPalindrome(s, i-1, i+1) {
				ch <- pal
			}

			for pal := range goFindPalindrome(s, i-1, i) {
				ch <- pal
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	s := "fdacniadffafadsasasfdadf"
	// findAllPalindrome から送られてくる回文部分文字列を順次出力
	for pal := range goFindAllPalindrome(s) {
		fmt.Println(pal)
	}
}
