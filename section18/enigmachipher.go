package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const ALPHABET1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func indexOf(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

type PlugBoard struct {
	alphabet    string
	forwardMap  map[rune]rune
	backwardMap map[rune]rune
}

func NewPlugBoard(mapAlphabet string) *PlugBoard {
	pb := &PlugBoard{
		alphabet:    ALPHABET1,
		forwardMap:  make(map[rune]rune),
		backwardMap: make(map[rune]rune),
	}
	pb.mapping(mapAlphabet)
	return pb
}

func (pb *PlugBoard) mapping(mapAlphabet string) {
	for i, c := range ALPHABET1 {
		mapped := rune(mapAlphabet[i])
		pb.forwardMap[c] = mapped
		pb.backwardMap[mapped] = c
	}
}

func (pb *PlugBoard) Forward(index int) int {
	char := rune(pb.alphabet[index])
	mapped := pb.forwardMap[char]
	return indexOf(pb.alphabet, mapped)
}

func (pb *PlugBoard) Backward(index int) int {
	char := rune(pb.alphabet[index])
	mapped := pb.backwardMap[char]
	return indexOf(pb.alphabet, mapped)
}

type Rotor struct {
	*PlugBoard
	offset    int
	rotations int
}

func NewRotor(mapAlphabet string, offset int) *Rotor {
	return &Rotor{
		PlugBoard: NewPlugBoard(mapAlphabet),
		offset:    offset,
		rotations: 0,
	}
}

func (r *Rotor) Rotate() int {
	shift := r.offset
	r.alphabet = r.alphabet[shift:] + r.alphabet[:shift]
	r.rotations += shift
	return r.rotations
}

func (r *Rotor) Reset() {
	r.rotations = 0
	r.alphabet = ALPHABET1
}

type Reflector struct {
	mapping map[rune]rune
}

func NewReflector(mapAlphabet string) *Reflector {
	m := make(map[rune]rune)
	for i, c := range ALPHABET1 {
		m[c] = rune(mapAlphabet[i])
	}

	for k, v := range m {
		if m[v] != k {
			panic(fmt.Sprintf("Invalid reflector mapping: %c -> %c, but %c -> %c", k, v, v, m[v]))
		}
	}
	return &Reflector{mapping: m}
}

func (ref *Reflector) Reflector(index int) int {
	letter := rune(ALPHABET1[index])
	reflected := ref.mapping[letter]
	return indexOf(ALPHABET1, reflected)
}

type EnigmaMachine struct {
	plugBoard *PlugBoard
	rotors    []*Rotor
	reflector *Reflector
}

func NewEnigmaMachine(pb *PlugBoard, rotors []*Rotor, ref *Reflector) *EnigmaMachine {
	return &EnigmaMachine{
		plugBoard: pb,
		rotors:    rotors,
		reflector: ref,
	}
}

func (em *EnigmaMachine) goThrough(char rune) rune {
	upperChar := char
	if upperChar >= 'a' && upperChar <= 'z' {
		upperChar = upperChar - 'a' + 'A'
	}
	if indexOf(ALPHABET1, upperChar) == -1 {
		return char
	}

	indexNum := indexOf(ALPHABET1, upperChar)
	indexNum = em.plugBoard.Forward(indexNum)

	for _, rotor := range em.rotors {
		indexNum = rotor.Forward(indexNum)
	}

	indexNum = em.reflector.Reflector(indexNum)

	for i := len(em.rotors) - 1; i >= 0; i-- {
		indexNum = em.rotors[i].Backward(indexNum)
	}

	indexNum = em.plugBoard.Backward(indexNum)
	resultChar := rune(ALPHABET1[indexNum])

	for i := len(em.rotors) - 1; i >= 0; i-- {
		if em.rotors[i].Rotate()%len(ALPHABET1) != 0 {
			break
		}
	}

	return resultChar
}

func (em *EnigmaMachine) Encrypt(text string) string {
	var sb strings.Builder
	for _, c := range text {
		sb.WriteRune(em.goThrough(c))
	}
	return sb.String()
}

func (em *EnigmaMachine) Decrypt(text string) string {
	for _, rotor := range em.rotors {
		rotor.Reset()
	}
	var sb strings.Builder
	for _, c := range text {
		sb.WriteRune(em.goThrough(c))
	}
	return sb.String()
}

func getRandomAlphabet() string {
	runes := []rune(ALPHABET1)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

func getReflectorMapping() string {
	runes := []rune(ALPHABET1)
	indexes := make([]int, len(runes))
	for i := 0; i < len(runes); i++ {
		indexes[i] = i
	}

	for len(indexes) >= 2 {
		i := rand.Intn(len(indexes))
		x := indexes[i]

		indexes = append(indexes[:i], indexes[i+1:]...)

		j := rand.Intn(len(indexes))
		y := indexes[j]

		indexes = append(indexes[:j], indexes[j+1:]...)
		runes[x], runes[y] = runes[y], runes[x]
	}
	return string(runes)
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	pb := NewPlugBoard(getRandomAlphabet())

	r1 := NewRotor(getRandomAlphabet(), 3)
	r2 := NewRotor(getRandomAlphabet(), 2)
	r3 := NewRotor(getRandomAlphabet(), 1)

	reflectorMapping := getReflectorMapping()
	reflector := NewReflector(reflectorMapping)

	machine := NewEnigmaMachine(pb, []*Rotor{r1, r2, r3}, reflector)

	message := "ATTACK SILICON VALLEY"
	encrypted := machine.Encrypt(message)
	fmt.Println("Encrypted:", encrypted)
	decrypted := machine.Decrypt(encrypted)
	fmt.Println("Decrypted:", decrypted)
}
