package main

import (
	"fmt"
	"os"
)

func hash(value string, level int) (hash int) {
	for _, value := range []rune(value) {
		hash += int(value)
	}
	return hash % level
}

type HashTable struct {
	level   int
	storage [][]string
}

func NewHashTable() *HashTable {
	hash := HashTable{}
	hash.level = 14
	return &hash
}

func (table *HashTable) Add(key string, value string) {
	index := hash(key, table.level)

	if len(table.storage)-1 <= index {
		table.storage = make([][]string, 0, 1)
		//table.storage[index] = []string{key,value}
		//table.storage[index] = append(table.storage[index], []string{key,value}...)
	}
}
func main() {
	hash := NewHashTable()
	hash.Add(`eau`, `person`)
	fmt.Println(hash)
	os.Exit(0)
}
