package hashtable

import (
	"testing"
	"reflect"
)

func TestNewHashTable(t *testing.T) {
	if reflect.TypeOf(*NewHashTable()).Name() != `HashTable` {
		t.Error(`Constructor HashTable not work`)
	}
}

func TestHashFunction(t *testing.T)  {
	if reflect.TypeOf(hash(`key`, 10)).Kind()  != reflect.Int {
		t.Error(`Hash function not make int`)
	}
}

func TestHashTable_Add(t *testing.T) {
	hash := NewHashTable()
	hash.Add(`one`, `number`)
	if hash.Search(`one`) != `number` {
		t.Error(`Hash table not have Add value`)
	}
}

func TestHashTable_Search(t *testing.T) {
	hash := NewHashTable()
	hash.Add(`one`, `number`)
	hash.Add(`two`, `number2`)
	if hash.Search(`two`) != `number2` {
		t.Error(`Hash table not work search method`)
	}
}

func TestHashTable_Remove(t *testing.T) {
	hash := NewHashTable()
	hash.Add(`one`, `number`)
	hash.Add(`two`, `number2`)

	if 	hash.Remove(`one`) != true{
		t.Error(`Method Remove not work`)
	}
}