package algorithm

import "fmt"

type LinkedList struct {
	size int
	head *DoublyListNode
	tail *DoublyListNode
}

func NewLinkedList() *LinkedList {
	head := &DoublyListNode{}
	tail := &DoublyListNode{}
	head.next = tail
	tail.prev = head
	return &LinkedList{
		head:head,
		tail:tail,
		size:0,
	}
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) Add(data interface{}) {
	node := &DoublyListNode{
		prev: ll.tail.prev,
		next: ll.tail,
		Data: data,
	}
	ll.tail.prev.next = node
	ll.tail.prev = node
	ll.size++
}

func (ll *LinkedList) Remove(data interface{}) {
	for temp := ll.head.next; temp != ll.tail; temp = temp.next {
		fmt.Println(temp.Data, data)
		if temp.Data == data {
			temp.LeaveChain()
			break
		}
	}
}

func (ll *LinkedList) Show() {
	for temp := ll.head.next; temp != ll.tail; temp = temp.next {
		fmt.Print(temp.Data)
		if temp.next != ll.tail  {
			fmt.Print(",")
		}
	}
	fmt.Println()
}