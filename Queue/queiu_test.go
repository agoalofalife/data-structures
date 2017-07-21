package Queue

import (
	"testing"
	"os"
)

var q Queue // init variable fot test

func TestMain(m *testing.M) {
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	os.Exit(m.Run())
}

func TestEnqueue(t *testing.T){
	q.Enqueue(4)
	if q.Size() != 4 {
		t.Error(`Method Enqueue not work`)
	}
}

func TestDequeue(t *testing.T) {
	q.Dequeue()
	if q.Size() != 3 {
		t.Error(`Method not dequeue from collection`)
	}
}


func TestSize(t *testing.T) {
	if q.Size() != 3 {
		t.Error(`Method not size from collection`)
	}
}

func TestIsEmpty(t *testing.T)  {
	if q.IsEmpty() != false {
		t.Error(`Method not IsEmpty`)
	}
}