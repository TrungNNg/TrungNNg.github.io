package main
import (
    "fmt"
)
var p = fmt.Println

type Node struct {
    val int
    pri int   // priority
    lc  *Node // left child 
    rc  *Node // right child
}

type Treap struct {
    root *Node
}

func (t *Treap) Add(v int) {
    t.add(v, t.root)
}

func (t *Treap) add(v int, node *Node) {
    if node == nil {
        return &Node{v, }
    }
}

func main() {
    p("hello")
}

