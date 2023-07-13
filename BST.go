package main

import (
    "fmt"
)
var p = fmt.Println
func main() {
    
    bst := BST{}
    root := bst.Insert(5)
    a := bst.Insert(2)
    b := bst.Insert(6)
    c := bst.Insert(3)
    d := bst.Insert(1)

    p(root,a,b,c,d)

    bst.Root = bst.Remove(bst.Root, 3)

    p(root,a,b,c,d)

}

type BST struct {
    Root *Node    
}

type Node struct {
    Val int
    Left *Node
    Right *Node
}

// INSERT
func (tree *BST) Insert(val int) *Node {
    if tree.Root == nil {
        tree.Root = &Node{Val: val}
        return tree.Root
    }
    return insert(tree.Root, val)
}

func insert(root *Node, val int) *Node {
    if root.Val <= val {
        if root.Left == nil {
            root.Left = &Node{Val:val}
            return root.Left
        }
        return insert(root.Left, val)
    }
    if root.Right == nil {
        root.Right = &Node{Val:val}
        return root.Right
    }
    return insert(root.Right, val)
}

// SEARCH
func (tree *BST) Search(val int) *Node {
    return search(tree.Root, val)
}

func search(root *Node, val int) *Node {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    } else if root.Val < val {
        return search(root.Left, val)
    } else {
        return search(root.Right, val)
    }
}

// REMOVE
func (tree *BST) Remove(root *Node, val int) *Node {
    if root == nil {
        return nil
    }
    if val < root.Val {
        root.Left = tree.Remove(root.Left, val)
    } else if val > root.Val {
        root.Right = tree.Remove(root.Right, val)
    } else {
        if root.Left == nil && root.Right == nil {
            return nil
        } else if root.Left != nil {
            return root.Right
        } else if root.Right != nil {
            return root.Left
        } else {
            // target node has both left and right subtree
            min_node := find_min(root.Right)
            root.Val = min_node.Val
            root.Right = tree.Remove(root.Right, min_node.Val)
        }
    }
    return root
}

func find_min(n *Node) *Node {
    if n.Left != nil {
        return find_min(n.Left)
    }
    return n
}

//HEIGHT
func (tree *BST) Height(val int) int {
    node := tree.Search(val)
    if node == nil {
        return -1     // no node with given value
    }
    return height(node)
}

func height(n *Node) int {
    if n == nil {
        return -1
    }
    return max(height(n.Left), height(n.Right)) + 1
}

// DEPTH
func (tree *BST) Depth(val int) int {
    node := tree.Search(val)
    if node == nil {
        return -1     // no node with given value
    }
    return depth(tree.Root, val)
}

func depth(n *Node, val int) int {
    if n == nil {
        return -1
    } 
    if n.Val == val {
        return 0
    }
    if n.Val < val {
        return depth(n.Left, val) + 1
    }
    return depth(n.Right, val) + 1
}

// BFS traverse (level order traversal)
func (tree *BST) BFS() {
    if tree.Root == nil { return }
    queue := []*Node{tree.Root}
    for len(queue) != 0 {
        curr := queue[0]
        queue = queue[1:]
        p(curr.Val)
        if curr.Left != nil {
            queue = append(queue, curr.Left)
        }
        if curr.Right != nil {
            queue = append(queue, curr.Right)
        }
    }
}

// DFS In-order traverse
func (tree *BST) DFS() {
    dfs(tree.Root)
}
func dfs(n *Node) {
    if n == nil {
        return
    }
    dfs(n.Left)
    p(n.Val)
    dfs(n.Right)
}

// Utils functions
func max(a, b int) int {
    if a > b {return a}
    return b
}






