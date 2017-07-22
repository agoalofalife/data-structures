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
	storage map[int]map[int]map[string]string
}

func NewHashTable() *HashTable {
	hash := HashTable{}
	hash.level = 14
	hash.storage = make(map[int]map[int]map[string]string)
	return &hash
}

func (table *HashTable) Add(key string, value string) {
	index := hash(key, table.level)

	if _, exist := table.storage[index]; exist == false {
		table.storage[index] = make(map[int]map[string]string)
		table.storage[index][len(table.storage[index])] = map[string]string{key: value}
	} else {
		inserted := false

		for keyStorage, _ := range table.storage[index] {
			if _, exist := table.storage[index][keyStorage][key] ;exist{
				table.storage[index][keyStorage][key] = value
				inserted = true
			}
		}

		if inserted == false {
			table.storage[index][len(table.storage[index])] = map[string]string{key: value}
		}
	}
}


func (table *HashTable) Search(key string) (value string){
	index := hash(key, table.level)

	if _, exist := table.storage[index]; exist == false {
		return
	} else {
		for keyStorage,_ := range table.storage[index]{
			if value, exist := table.storage[index][keyStorage][key];exist {
				return value
			}
		}
	}
	return
}

func (table *HashTable) Remove(key string) bool{
	index := hash(key, table.level)
	_, exist := table.storage[index][0][key]
	if len(table.storage[index]) == 1 && exist {
		delete(table.storage, index)
		return  true
	} else {
		for keyStorage,_ := range table.storage[index] {
			if _, exist := table.storage[index][keyStorage][key] ;exist{
				delete(table.storage[index], keyStorage)
				return true
			}
		}
	}
	return false
}
func main() {
	hash := NewHashTable()
	hash.Add(`fido`, `dog`)
	hash.Add(`rex`, `dinosour`)
	hash.Add(`rex`, `rex`)
	hash.Add(`rew`, `rew`)
	hash.Add(`reas`, `reasValue`)

	fmt.Println(	hash.Remove(`fido`))
	os.Exit(0)
}
