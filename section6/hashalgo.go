package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
)

type KeyValuePair struct {
	Key   string
	Value any
}

type IHashTable interface {
	Hash(key string) int
	Add(key string, value any)
	Get(key string) (any, bool)
	Print()
}

type HashTable struct {
	size  int
	table [][]KeyValuePair
}

func NewHashTable(size int) IHashTable {
	return &HashTable{
		size:  size,
		table: make([][]KeyValuePair, size),
	}
}

func (ht *HashTable) Hash(key string) int {
	hash := md5.Sum([]byte(key))
	hashHex := hex.EncodeToString(hash[:])

	hashInt := new(big.Int)
	hashInt.SetString(hashHex, 16)

	index := hashInt.Int64() % int64(ht.size)
	if index < 0 {
		index += int64(ht.size)
	}
	return int(index)
}

func (ht *HashTable) Add(key string, value any) {
	index := ht.Hash(key)
	for i, kv := range ht.table[index] {
		if kv.Key == key {
			ht.table[index][i].Value = value
			return
		}
	}
	ht.table[index] = append(ht.table[index], KeyValuePair{Key: key, Value: value})
}

func (ht *HashTable) Print() {
	for i, bucket := range ht.table {
		fmt.Printf("%d", i)
		for _, kv := range bucket {
			fmt.Printf("--> [%s: %v] ", kv.Key, kv.Value)
		}
		fmt.Println()
	}
}

func (ht *HashTable) Get(key string) (any, bool) {
	index := ht.Hash(key)
	for _, kv := range ht.table[index] {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return nil, false
}

func main() {
	var ht IHashTable = NewHashTable(10)
	ht.Add("car", "Tesla")
	ht.Add("car", "Toyota")
	ht.Add("pc", "Mac")
	ht.Add("sns", "Youtube")

	ht.Print()

	if value, found := ht.Get("sns"); found {
		fmt.Println("sns:", value)
	} else {
		fmt.Println("sns not found")
	}
}
