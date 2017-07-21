package main

import (
	"fmt"
	"os"
)

type Set struct {
	collection []int
}

func (set *Set) Add(value int) {
	if set.Exist(value) == false {
		set.collection = append(set.collection, value)
	} else {
		panic(`Value already exists in Set`)
	}
}

func (set *Set) AddRange(values []int) {
	for value := range values {
		set.Add(value)
	}
}

func (set *Set) Exist(value int) bool {
	if len(set.collection) == 0 {
		return false
	} else {
		for _, valueInCollection := range set.collection {
			if value == valueInCollection {
				return true
			}
		}
	}
	return false
}

func (set *Set) Remove(value int) bool {
	for key, valueInCollection := range set.collection {
		if value == valueInCollection {
			set.collection = append(set.collection[:key], set.collection[key+1:]...)
			fmt.Println(set.collection)
			os.Exit(0)
			return true
		}
	}
	return false
}

func (set *Set) Count() int {
	return len(set.collection)
}

func main() {
	s := Set{}
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)
	s.Remove(4)
	fmt.Println(s)
}
