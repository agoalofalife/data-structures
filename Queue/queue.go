//---------------------------------------------------------------------------
//-[ Queue ]-----------------------------------------------------------------
//---------------------------------------------------------------------------
//---------------------------------------------------------------------------

package Queue

import "fmt"

type Queue struct {
	collection []int
}

func (queue *Queue) Enqueue(value int) {
	queue.collection = append(queue.collection, value)
}

func (queue *Queue) Dequeue() (value int) {
	value = queue.collection[:1][0]
	queue.collection = queue.collection[1:]
	return
}

func (queue *Queue) Size() int {
	return len(queue.collection)
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.collection) == 0
}

func main() {
	// Example FIFO
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue())
}
