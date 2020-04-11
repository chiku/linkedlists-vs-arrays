package linkedlist

import "strconv"

type node struct {
	value int
	next  *node
}

func newNode(value int, next *node) *node {
	return &node{value: value, next: next}
}

type Linkedlist struct {
	head *node
}

func NewLinkedlist() *Linkedlist {
	return &Linkedlist{head: nil}
}

func (list Linkedlist) IsEmpty() bool {
	return list.head == nil
}

func (list Linkedlist) Get(position int) int {
	var i int
	var p *node

	for i, p = 0, list.head; i < position && p != nil; i, p = i+1, p.next {
	}

	if p == nil {
		return 0
	}

	return p.value
}

func (list *Linkedlist) Insert(value int, position int) {
	if list.IsEmpty() || position == 0 {
		list.head = newNode(value, list.head)
		return
	}

	var i int
	var p *node

	for i, p = 0, list.head; i < position-1 && p.next != nil; i, p = i+1, p.next {
	}

	p.next = newNode(value, p.next)
}

func (list Linkedlist) String() string {
	result := "[ "

	for p := list.head; p != nil; p = p.next {
		result += "<"
		result += strconv.Itoa(p.value)
		result += "> -> "
	}

	result += "NULL ]"

	return result
}
