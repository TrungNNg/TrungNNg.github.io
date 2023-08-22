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

func (l *DLL) addFront(v int) *Node {
    if l.head == nil {
        l.head = &Node{val:v}
        l.tail = l.head
        return l.head
    }
    l.head.prev = &Node{nil, l.head, v}
    l.head = l.head.prev
    return l.head
}

func (l *DLL) removeTail() {
    if l.tail != nil {
        l.tail = l.tail.prev
        l.tail.next = nil
    }
}

func (l *DLL) remove(npt *Node) {
    if l.head == npt && l.tail == npt {
        l.head = nil
        l.tail = nil
    } else if l.tail == npt {
        l.removeTail()
    } else if l.head == npt {
        l.head = npt.next
        if l.head != nil {
            l.head.prev = nil
        }
    } else {
        npt.prev.next = nil
        npt.next.prev = nil
    }
}

// print head to tail
func (l *DLL) pht() {
    curr := l.head
    for curr != nil {
        fmt.Print(curr.val, " ")
        curr = curr.next
    }
    p()
}

// print tail to head
func (l *DLL) pth() {
    curr := l.tail
    for curr != nil {
        fmt.Print(curr.val, " ")
        curr = curr.prev
    }
    p()
}

func main() {

    /* test DLL
    l := DLL{}
    l.addFront(1)
    l.addFront(2)
    l.addFront(3)
    l.addFront(4)

    // head to tail
    l.pht()

    l.removeTail()

    l.pht()
    */

    // an LRU cache with capacity 5
    cache := LRU{DLL{}, map[int]*Node{}, 0, 5}
    cache.add(1)
    cache.add(2)
    cache.add(3)
    cache.add(4)
    cache.add(5)
    cache.pLRU()
    cache.remove(1)
    cache.pLRU()
    cache.add(6)
    cache.pLRU()
}

type LRU struct {
    list    DLL
    m       map[int]*Node
    size    int
    max_cap int
}

func (c *LRU) add(v int) {
    if c.size >= c.max_cap {
        // evict last recently use
        c.list.removeTail()
        c.size -= 1
    }
    npt := c.list.addFront(v)
    c.size += 1
    c.m[v] = npt
}

func (c *LRU) contains(v int) bool {
     _, ok := c.m[v]
     return ok
}

func (c *LRU) remove(v int) {
    if npt, ok := c.m[v]; ok {
        c.list.remove(npt)
        c.size -= 1
    }
}

// print LRU cache
func (c *LRU) pLRU() {
    c.list.pht()
}





