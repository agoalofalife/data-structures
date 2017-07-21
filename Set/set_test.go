package Set

import (
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := Set{}
	s.Add(2)

	if s.Count() != 1 {
		t.Error(`Add not work`)
	}
}

func TestSet_AddRange(t *testing.T) {
	s := Set{}
	arr := []int{1, 3, 4}
	s.AddRange(arr)

	if s.Count() != 3 {
		t.Error(`AddRange not work`)
	}
}

func TestSet_Exist(t *testing.T) {
	s := Set{}
	if s.Exist(2) != false {
		t.Error(`Exist not work`)
	}
}

func TestSet_Remove(t *testing.T) {
	s := Set{}
	s.Add(2)

	if s.Remove(2) != true {
		t.Error(`Remove not work`)
	}
}

func TestSet_Count(t *testing.T) {
	s := Set{}
	if s.Count() != 0 {
		t.Error(`Count not work`)
	}
}

func TestSet_Union(t *testing.T) {
	s := Set{}
	s.Add(1)
	s2 := Set{}
	s2.Add(2)

	if len(s.Union(s2).collection) != 2 {
		t.Error(`Method Union not work`)
	}
}

func TestSet_Intersection(t *testing.T) {
	s := Set{}
	s.Add(1)
	s.Add(2)
	s2 := Set{}
	s2.Add(2)
	s2.Add(1)

	if len(s.Intersection(s2).collection) != 2 {
		t.Error(`Method Intersection not work`)
	}
}

func TestSet_Difference(t *testing.T) {
	s := Set{}
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s2 := Set{}
	s2.Add(2)
	s2.Add(1)

	if s.Difference(s2).collection[0] != 3 {
		t.Error(`Method Difference not work`)
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	s := Set{}
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s2 := Set{}
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)
	s2.Add(6)

	if len(s.SymmetricDifference(s2).collection) != 4 {
		t.Error(`Method Difference not work`)
	}
}

func TestSet_IsSubset(t *testing.T) {
	s := Set{}
	s.Add(4)
	s2 := Set{}
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)
	s2.Add(6)

	if s.IsSubset(s2) != true {
		t.Error(`Method IsSubset not work`)
	}
}
