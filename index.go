package main

import (
	"fmt"
)

type List struct {
	head, tail *Point
	size       int
}

type Point struct {
	data      int
	next, pre *Point
}

func create() (list List) {
	var point = new(Point)
	point.data = 0
	point.next = nil
	point.pre = nil
	list.head = point
	return list
}

func (list *List) exist(data int) bool {
	var exist = false
	for now := list.head; now != nil; now = now.next {
		if now.data == data {
			exist = true
			break
		}
	}
	return exist
}

func (list *List) append(data int) {
	var point = new(Point)
	point.data = data
	list.tail.next = point
	list.size++
}
func (list *List) insert(data int) {
	var point = new(Point)
	point.data = data
	var flag = false
	for now := list.head; now != nil; now = now.next {
		if now.data <= data && now.next.data >= data {
			point.next = now.next
			now.next = point
			list.size++
			flag = true
			break
		}
	}
	if !flag {
		list.append(data)
	}
}

func (list *List) delete(data int) {
	for now := list.head; now != nil; now = now.next {
		if now.data == data {
			if now.next != nil {
				now.pre = now.next
			} else {
				now.pre = nil
			}
			list.size--
		}
	}
}

func (list *List) print() {
	for now := list.head; now != nil; now = now.next {
		fmt.Println("the data is: ", now.data)
	}
	fmt.Println("list size is: ", list.size)
}

func main() {
	var list = create()
	list.append(5)
	list.append(10)
	list.insert(2)
	list.insert(8)
	list.print()
}
