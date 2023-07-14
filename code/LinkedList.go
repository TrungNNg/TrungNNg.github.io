package main

import (
    "fmt"
)
func main() {
    list := LinkedList{}
    list.Add(1)
    list.Add(2)
    list.Print()
    list.ReverseRecursive(list.Head)
    list.Print()
}

type LinkedList struct {
    Head *Node
}

type Node struct {
    Val int
    Next *Node
}

// reverse iterative
func (l *LinkedList) ReverseIterative() {
    var prev *Node
    curr := l.Head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    l.Head = prev
}

// reverse recursive
func (l *LinkedList) ReverseRecursive(n *Node) {
    if l.Head == nil { return }
    if n.Next == nil {
        l.Head = n
        return
    }
    l.ReverseRecursive(n.Next)
    n.Next.Next = n
    n.Next = nil
}

// if list has even number of node, return right node
func (l *LinkedList) FindMiddleNode() *Node {
    f, s := l.Head, l.Head
    for f != nil && f.Next != nil {
        s = s.Next
        f = f.Next.Next
    }
    return s
}

func (l *LinkedList) RemoveAll(val int) {
    dummy_node := Node{Val:0, Next: l.Head}
    curr,prev := l.Head, &dummy_node
    for curr != nil {
        if curr.Val == val {
            prev.Next = curr.Next
        } else {
            prev = curr
        }
        curr = curr.Next
    }
    l.Head = dummy_node.Next
}

// O(n * len(vals))
func (l *LinkedList) Add(val int) {
    new_node := Node{Val: val, Next: nil}
    if l.Head == nil {
        l.Head = &new_node
        return
    }
    curr := l.Head
    for curr.Next != nil {
        curr = curr.Next
    }
    curr.Next = &new_node
}

func (l *LinkedList) Print() {
    curr := l.Head
    s := ""
    for curr != nil {
       fmt.Print(s,curr.Val)
       s = " -> "
       curr = curr.Next
    }
    fmt.Println()
}


