package chapter06

import "fmt"


type Queque struct {
	len int
	head *Node
	tail *Node
}

type Node struct {
	value string
	next *Node
}

func (q *Queque) enQueque(item string) {
	n := &Node{
		value: item,
		next: nil,
	}
	if q.len == 0 {
		q.head = n
	} else {
		q.tail.next = n			
	}
	q.tail = n
	q.len++
}  

func (q *Queque) deQueque() string {
	val := q.head.value
	nextPtr := q.head.next
	q.head = nil
	q.head = nextPtr
	q.len--
	return val
}

func (q *Queque) size() int {
	return q.len
}

func HasSellerFor(name string, graph map[string][]string) bool {
	search_queque := &Queque{}
	for _, v := range graph[name] {
		search_queque.enQueque(v)
	}	
	searched := make(map[string]struct{})
	for search_queque.size() > 0 {
		person := search_queque.deQueque()
		_, ok := searched[person]
		if !ok{
			if person_is_seller(person){
				fmt.Printf("Person %s is a mango seller ", person)
				println()
				return true
			} else {
				for _, v := range graph[person] {
					search_queque.enQueque(v)
				}
				searched[person] = struct{}{}	
			}
		}
	}
	return false
} 


func person_is_seller(name string) bool {
	return name[len(name)-1:] == "m" 
}



