// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//test in main
	var s Stack
	//push data in stack
	s.push(1)
	s.push(7)
	s.push(9)
	s.push(11)
	//print the stack
	s.print()
	//pop an item from the stack
	s.pop()
	//print the stack
	s.print()

}

// stack data structure
// 1. create a type
type Stack struct {
	data []int
}

// 2. implement push function
func (stack *Stack) push(item int) {
	stack.data = append(stack.data, item)
}

// 3. give it the pop  unction
func (stack *Stack) pop() {
	l := len(stack.data)
	fmt.Println("pop up item: ", stack.data[l-1])
	if l == 0 {
		fmt.Println("stack is empty, cannot pop anymore")
	} else {
		stack.data = stack.data[:l-1]
	}
}

// 4. print the stack
func (stack *Stack) print() {
	fmt.Println("-start-")
	for i := len(stack.data) - 1; i >= 0; i-- {
		fmt.Println(stack.data[i])
	}
	fmt.Println("-finish-")
}
