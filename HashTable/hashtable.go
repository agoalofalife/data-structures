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
	storage map[int]map[string]string
}

func NewHashTable() *HashTable {
	hash := HashTable{}
	hash.level = 14
	hash.storage = make(map[int]map[string]string)
	return &hash
}

func (table *HashTable) Add(key string, value string) {
	index := hash(key, table.level)

	if _, exist := table.storage[index]; exist == false {

		table.storage[index] = map[string]string{key: value}
	} else {
		inserted := false

		for keyStorage, _ := range table.storage[index] {
			if table.storage[index][keyStorage] != value {
				table.storage[index][keyStorage] = value
			}
			inserted = true
		}

		if inserted == false {
			table.storage[index] = map[string]string{key: value}
		}
	}
}

func main() {
	hash := NewHashTable()
	hash.Add(`fido`, `dog`)
	hash.Add(`rex`, `dinosour`)
	hash.Add(`rex`, `rex`)
	hash.Add(`rew`, `rew`)
	hash.Add(`reas`, `reas`)

	fmt.Println(hash)
	os.Exit(0)
}
