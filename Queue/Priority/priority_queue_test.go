package PriorityQueue

import (
	"os"
	"testing"
)

var q PriorityQueue // init variable fot test

func TestMain(t *testing.M) {
	q.Enqueue(1, 6)
	q.Enqueue(2, 3)
	q.Enqueue(3, 1)
	os.Exit(t.Run())
}

func TestEnqueue(t *testing.T) {
	if q.Size() != 3 {
		t.Error(`Error enqueue method`)
	}
}

func TestDequeue(t *testing.T) {
	if q.Dequeue() != 3 {
		t.Error(`Error dequeue method`)
	}
}

func TestSize(t *testing.T) {
	if q.Size() != 3 {
		t.Error(`Method not size from collection`)
	}
}

func TestIsEmpty(t *testing.T) {
	if q.IsEmpty() != false {
		t.Error(`Method not IsEmpty`)
	}
}
