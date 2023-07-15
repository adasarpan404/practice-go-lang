package core

type Queue struct {
	Front *Node
	Rear  *Node
}

func NewQueue() *Queue {
	return &Queue{Front: nil, Rear: nil}
}

func (q *Queue) Enqueue(data string) {
	newNode := &Node{Data: data, Next: nil}

	if q.IsEmpty() {
		q.Front = newNode
		q.Rear = newNode
	} else {
		q.Rear.Next = newNode
		q.Rear = newNode
	}
}

func (q *Queue) Dequeue() string {
	if q.IsEmpty() {
		return ""
	}

	data := q.Front.Data
	q.Front = q.Front.Next
	if q.IsEmpty() {
		q.Rear = nil
	}
	return data
}

func (q *Queue) IsEmpty() bool {
	return q.Front == nil
}
