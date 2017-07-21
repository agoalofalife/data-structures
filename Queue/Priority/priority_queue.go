//---------------------------------------------------------------------------
//-[ Priority Queue ]-----------------------------------------------------------
//---------------------------------------------------------------------------
//---------------------------------------------------------------------------

package PriorityQueue

import (
	"fmt"
)

type PriorityQueue struct {
	collection map[int]int
}

// add new value depend of priority
func (queue *PriorityQueue) Enqueue(value int, priority int) {
	if queue.IsEmpty() {
		queue.collection = make(map[int]int)
		queue.collection[priority] = value
	} else {
		queue.collection[priority] = value
	}

}

func (queue *PriorityQueue) Dequeue() (middle int) {
	var middlePriority int
	for priority, value := range queue.collection {
		if middlePriority == 0 && priority > middlePriority {
			middle = value
			middlePriority = priority
		} else if middlePriority > priority {
			middle = value
			middlePriority = priority
		}
	}
	return middle
}

func (queue *PriorityQueue) Size() int {
	return len(queue.collection)
}

func (queue *PriorityQueue) IsEmpty() bool {
	return len(queue.collection) == 0
}

func main() {
	// Example FIFO
	q := PriorityQueue{}
	q.Enqueue(1, 6)
	q.Enqueue(2, 3)
	q.Enqueue(3, 1)

	fmt.Println(q.Dequeue())
}
