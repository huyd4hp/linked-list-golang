package main

import (
	"linkedlist/src/linkedlist"
)

func main(){
	list := &linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Print()
}