package main

// import "fmt"

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateKey(message, keyword string) string {
	key := keyword
	remainLength := len(message) - len(keyword)
	for i := range remainLength {
		key += string(key[i])
	}
	return key
}

func encrypt(message, key string) string {
	result := ""
	for i, char := range message {
		if char < 'A' || char > 'Z' {
			result += string(char)
			continue
		}

		newVal := (int(char)+int(key[i]))%len(ALPHABET) + int('A')
		result += string(rune(newVal))
	}
	return result
}

func decrypt(cipherText, key string) string {
	result := ""
	for i, char := range cipherText {
		if char < 'A' || char > 'Z' {
			result += string(char)
			continue
		}

		newVal := (int(cipherText[i])-int(key[i])+len(ALPHABET))%len(ALPHABET) + int('A')
		result += string(rune(newVal))
	}
	return result
}

// func main() {
// 	t := "ATTACK SILICON VALLEY"
// 	k := generateKey(t, "HELLO")
// 	e := encrypt(t, k)
// 	fmt.Println(e)
// 	d := decrypt(e, k)
// 	fmt.Println(d)
// }
