package main
import (
    "fmt"
    "math/rand"
)
var p = fmt.Println

// higher pri mean higher priority
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
    t.root = t.add(v, t.root)
}

func (t *Treap) add(v int, node *Node) *Node {
    if node == nil {
        // a psudo-random number from [0 - 1000)
        ran := rand.Intn(1000)
        return &Node{v, ran, nil, nil}
    } else if v < node.val {
        node.lc = t.add(v, node.lc)
        if node.lc.pri < node.pri {
            node = t.leftRotate(node.lc, node)
        }
    } else { // v >= node.val
        node.rc = t.add(v, node.rc)
        if node.rc.pri > node.pri {
            node = t.rightRotate(node, node.rc)
        }
    }
    return node
}

func (t *Treap) leftRotate(l, r *Node) *Node {
    r.lc = l.rc
    l.rc = r
    return l
}

func (t *Treap) rightRotate(l, r *Node) *Node {
    l.rc = r.lc
    r.lc = l
    return r
}

func (t *Treap) BFS() {
    if t.root == nil { return }
    queue := []*Node{t.root}
    for len(queue) != 0 {
        curr := queue[0]
        queue = queue[1:]
        p(curr.val)
        if curr.lc != nil {
            queue = append(queue, curr.lc)
        }
        if curr.rc != nil {
            queue = append(queue, curr.rc)
        }
    }
}

func main() {
    treap := Treap{}
    treap.Add(1)
    treap.Add(2)
    treap.Add(3)
    treap.Add(4)
    treap.Add(5)
    treap.Add(6)
    treap.BFS()
    p("here")
}









