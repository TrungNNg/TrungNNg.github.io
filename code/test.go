package main
import (
    "fmt"
)
var p = fmt.Println

type Node struct {
    prev *Node
    next *Node
    val int
}

type DLL struct {
    head *Node
    tail *Node
}

// TODO fix this 
func (l *DLL) addFront(v int) {
    l.head.prev = &Node{nil, l.head, v}
    l.head = l.head.prev
}

func (l *DLL) removeTail() {
    l.tail = l.tail.prev
    l.tail.next = nil
}

func main() {
    // cache := LRU{DLL{}, map[int]*Node{}}
    l := DLL{}
    l.addFront(1)
    l.addFront(2)
    l.addFront(3)
    l.addFront(4)

    p(l.head.val)
}

type LRU struct {
    list DLL
    m map[int]*Node
}

func (c *LRU) add(v int) {
    c.list.addFront(v)
}

func (c *LRU) contains(v int) {

}


