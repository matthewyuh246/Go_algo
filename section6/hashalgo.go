package hash

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

type KeyValuePair struct {
	Key   string
	Value any
}

type IHashTable interface {
	hash(key string) int
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

	return int(hashInt.Int64() % int64(ht.size))
}

func (ht *HashTable) Add(key string, value any) {
	index := ht.hash(key)
	for i, kv := range ht.table[index] {
		if kv.Key == key {
			ht.table[index][i].Value = value
			return
		}
	}
	ht.table[index] = append(ht.table[index], KeyValuePair{Key: key, Value: value})
}

func (ht *HashTable) Print() {
	for i, kv := range ht.table {
		print(i, "-->")
	}
}

func (ht *HashTable) Get(key string) (any, bool) {

}

func main() {

}
