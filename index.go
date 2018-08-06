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

func (list *List) init() {
	var head, tail = new(Point), new(Point)
	head.next = tail
	tail.pre = head
	list.head = head
	list.tail = tail
	list.size = 0
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

	point.next = list.tail
	point.pre = list.tail.pre

	point.pre.next = point
	list.tail.pre = point

	list.size++
}
func (list *List) insert(data int) {
	var point = new(Point)
	point.data = data
	var flag = false
	for now := list.head; now != nil; now = now.next {
		if now.data <= data && now.next.data >= data {
			point.next = now.next
			point.pre = now
			now.next.pre = point
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
				now.next.pre = now.pre
				now.pre.next = now.next
			} else {
				now.pre = nil
			}
			list.size--
		}
	}
}

func (list *List) print() {
	fmt.Println("list size is: ", list.size)
	fmt.Println("left to right is: ")
	for now := list.head; now != nil; now = now.next {
		fmt.Print(now.data, " ")
	}
	fmt.Println()
	fmt.Println("right to left is: ")
	for now := list.tail; now != nil; now = now.pre {
		fmt.Print(now.data, " ")
	}
}

type Say interface {
	say()
}

type Animal struct {
	name string
}

func (animal *Animal) say() {
	fmt.Println(animal.name)
}

func (animal *Animal) set(name string) {
	animal.name = name
}

func hello(animal Say) {
	animal.say()
}

func main() {
	var list = new(List)
	list.init()
	list.append(5)
	list.append(10)
	list.insert(2)
	list.insert(8)
	list.delete(5)
	list.insert(4)
	list.print()
	var dog = new(Animal)
	var cat = new(Animal)
	dog.set("dogg")
	cat.set("Crystal")
	hello(dog)
	hello(cat)
	dog.say()
	cat.say()

}
