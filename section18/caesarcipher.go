package main

// import "fmt"

const lenAlphabet = 26

func caesarCipher(text string, shift int) string {
	result := []rune{}
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			newVal := (int(char-'A') + shift) % lenAlphabet
			if newVal < 0 {
				newVal += lenAlphabet
			}
			result = append(result, rune(newVal)+'A')
		} else if char >= 'a' && char <= 'z' {
			newVal := (int(char-'a') + shift) % lenAlphabet
			if newVal < 0 {
				newVal += lenAlphabet
			}
			result = append(result, rune(newVal)+'a')
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

type Candidate struct {
	Shift     int
	Decrypted string
}

func caesarCipherHack(text string) []Candidate {
	var candidates []Candidate
	for candidateShift := 1; candidateShift <= lenAlphabet; candidateShift++ {
		var reverted []rune
		for _, char := range text {
			if char >= 'A' && char <= 'Z' {
				index := int(char) - candidateShift
				if index < int('A') {
					index += lenAlphabet
				}
				reverted = append(reverted, rune(index))
			} else if char >= 'a' && char <= 'z' {
				index := int(char) - candidateShift
				if index < int('a') {
					index += lenAlphabet
				}
				reverted = append(reverted, rune(index))
			} else {
				reverted = append(reverted, char)
			}
		}
		candidates = append(candidates, Candidate{Shift: candidateShift, Decrypted: string(reverted)})
	}
	return candidates
}

// func main() {
// 	encrypted := caesarCipher("ATTACK SILICON VALLY by enginner", 3)
// 	fmt.Println(encrypted)

// 	decrypted := caesarCipher(encrypted, -3)
// 	fmt.Println(decrypted)

// 	candidates := caesarCipherHack(encrypted)
// 	for _, candidate := range candidates {
// 		fmt.Printf("%2d %s\n", candidate.Shift, candidate.Decrypted)
// 	}
// }
