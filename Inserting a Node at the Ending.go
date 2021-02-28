package main

type ListNode struct {
	data interface{}
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) InsertAtEnd(data interface{}) {
	node := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current != nil {
			current = current.next
		}
		current.next = node
	}
	ll.size++
	return
}
