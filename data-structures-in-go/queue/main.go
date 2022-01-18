package main

import "fmt"

func main() {
	//test in main
	q := Queue{}
	//enqueue some items
	q.enqueue(2)
	q.enqueue(6)
	q.enqueue(9)
  //display
	q.display()
	//dequeue
	q.dequeue()
	//display
	q.display()
}

// stack data structure
// 1. create a type
type Queue struct {
	data []int
}

//2. enqueue func , add from the beginning of the slice
func (q *Queue) enqueue(item int) {
	q.data = append(q.data, item)
}

//3. dequeue func
func (q *Queue) dequeue() int {
	l := len(q.data)
	if l != 0 {
		popout := q.data[0]
		q.data = q.data [1:]
		fmt.Printf("removed item: %d\n", popout)
    return popout
	} else {
		fmt.Println("q is empty, cannot dequeue")
		return 0
	}
}

//4. Display the queue
func (q Queue) display() {
	fmt.Println("tail <------- head")
	for _,v := range q.data {
		fmt.Printf("%d<-", v)
	}
}
