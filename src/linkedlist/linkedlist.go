package linkedlist

import "fmt"

type LinkedList struct{
	Head * Node
}

func (list *LinkedList) Add(Value int) {
	newNode := &Node{
		Value: Value,
	}
	if list.Head == nil {
		list.Head = newNode
	} else{
		currentNode := list.Head
		for currentNode.Next != nil{
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}

func (list *LinkedList) Print() {
	current := list.Head
    for current != nil {
        fmt.Print(current.Value, " ")
        current = current.Next
    }
    fmt.Println()
}