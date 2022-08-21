package chapter02



import "fmt"



type LinkedList struct{
	len 	int
	head   *Node
	tail   *Node
}

type Node struct{
	value int
	next  *Node
}

func (l *LinkedList) add(item int){
	n := &Node{
		value: item,
		next: nil,
	}
	if l.len == 0 {
		l.head = n
	} else {
		l.tail.next = n			
	}
	l.tail = n
	l.len++
}	


func (l *LinkedList) insertAt(index, item int) {
	if index > l.len {
		return
	}
	n := &Node{
		value: item,
		next: nil,
	}
	l.len++
	if index == 0 {
		n.next = l.head
		l.head = n
		return
	}
	ib := l.head   // node before the target insertion position 
 	i := 1
	for i < index {
		ib = ib.next
		i++
	}  
	n.next, ib.next = ib.next, n // make swap of next field
	if n.next == nil {
		l.tail = n
	}
	return 
}

func (l *LinkedList) deleteAt(index int) {
	if index > l.len {
		return
	}
	l.len--
	if index == 0 {
		target := l.head
		l.head = target.next
		target = nil
		return
	}
	ib := l.head   // node before the target deletiom position 
 	i := 1
	for i < index {
		ib = ib.next
		i++
	}
	target := ib.next
	ib.next = target.next
	if target.next == nil {
		l.tail = ib
	}
	target = nil
	return 
}


func (l *LinkedList) indexOf(item int) int {
	c := l.head
	i := 0
	for c != nil {
		if c.value == item {
			return i
		}
		c = c.next
		i++ 
	}
	return -1  
}


func (l *LinkedList) getAt(index int) int {
	c := l.head
	i := 0
	for c != nil {
		if i == index {
			return c.value
		}
		c = c.next
		i++ 
	}
	return -1  
}

func (l *LinkedList) size() int {
	return l.len
}


func (l LinkedList) display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.value)
		l.head = l.head.next
	}
	fmt.Println()
}
