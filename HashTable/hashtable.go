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

func capUp(slice [][]string, number int) [][]string {
	newSlice := make([][]string, number+1)
	for i := range slice {
		newSlice[i] = slice[i]
	}
	slice = newSlice
	return slice
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
	fmt.Println(index)
	// feature language golang equals isset or exist
	if len(table.storage)-1 <= index {
		// size up
		table.storage = capUp(table.storage, index)

		//table.storage = make([][]string, index+1)
		table.storage[index] = []string{key, value}
	} else {

		inserted := false
		for keyStorage, _ := range table.storage {
			if table.storage[index][keyStorage] == key {
				table.storage[index][1] = value
			}
			inserted = true
		}

		if inserted == false {
			table.storage[index] = append(table.storage[index], key, value)
		}
	}
}

func main() {
	hash := NewHashTable()
	hash.Add(`fido`, `dog`)
	hash.Add(`rex`, `dinosour`)
	hash.Add(`rex`, `dinosour`)

	fmt.Println(hash)
	os.Exit(0)
}
