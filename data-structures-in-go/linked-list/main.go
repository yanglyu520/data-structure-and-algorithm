package main

import "fmt"

//node struct consists of data and pointer to the next node

type node struct {
	data int
	next *node
}


type linkedlist struct {
	head   *node
	tail   *node
	length int
}

func (list *linkedlist)initLinkedList(){
	list.head = nil
	list.tail = nil
	list.length = 0
}


func (list linkedlist) printLinkedList() {
	fmt.Println("----start---")
	for list.head.next != nil {
		fmt.Printf("%d\n", list.head.data)
		list.head = list.head.next
	}
	fmt.Println("----end---")
}
//insert function using method receiver
//define the method outside the data type and use a receiver to point to the
func (list *linkedlist) prepend(node *node) {
	second := list.head
	list.head = node
	list.head.next = second
	list.length++
}



func (list *linkedlist) postpend(node *node){

}

func (list *linkedlist) insert(){

}

func main() {
	listlist := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 10}

	listlist.prepend(node1)
	listlist.prepend(node2)
	listlist.print()
}

func (list *linkedlist) delisteteNode (n *node){
   
}
