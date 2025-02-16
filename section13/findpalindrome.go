package main

// findPalindrome は、文字列 s 内の left, right を中心に拡大しながら
// 回文となる部分文字列をすべて探索して、スライスとして返します。
func findPalindrome(s string, left, right int) []string {
	var result []string
	// 左右のインデックスが文字列の範囲内にある限りループ
	for left >= 0 && right < len(s) {
		// 左右の文字が一致しなければ、回文ではないのでループを終了
		if s[left] != s[right] {
			break
		}
		// 回文である部分文字列（s[left:right+1]）を結果に追加
		result = append(result, s[left:right+1])
		// 左に1つ、右に1つずつ広げる
		left--
		right++
	}
	return result
}

// findAllPalindrome は、文字列 s 内のすべての回文部分文字列を探索して返します。
func findAllPalindrome(s string) []string {
	var result []string
	n := len(s)
	// 文字列が空の場合は空のスライスを返す
	if n == 0 {
		return result
	}
	// 文字列の長さが1なら、その文字列自身が回文
	if n == 1 {
		result = append(result, s)
		return result
	}

	// 中心位置を基準に回文を探索する
	// Python の for i in range(1, len(s)-1) と同等（i は 1 から n-2 まで）
	for i := 1; i <= n-2; i++ {
		// 奇数長の回文：中心が i の場合（左右に 1 ずつ広げる）
		oddPalindromes := findPalindrome(s, i-1, i+1)
		for _, p := range oddPalindromes {
			result = append(result, p)
		}

		// 偶数長の回文：中心が i-1 と i の間の場合
		evenPalindromes := findPalindrome(s, i-1, i)
		for _, p := range evenPalindromes {
			result = append(result, p)
		}
	}

	return result
}

// func main() {
// 	// テスト用の文字列
// 	s := "fdacniadffafadsasasfdadf"
// 	// 文字列内の回文部分文字列をすべて取得
// 	palindromes := findAllPalindrome(s)
// 	// 結果を出力
// 	fmt.Println(palindromes)
// }
