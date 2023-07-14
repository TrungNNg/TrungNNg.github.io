package main

import (
    "fmt"
)
var p = fmt.Println
func main() {
    
    bst := BST{}
    bst.Insert(5)
    bst.Insert(2)
    bst.Insert(6)
    bst.Insert(3)
    bst.Insert(1)

//    bst.BFS()

    p(bst.Depth(3))

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
func (tree *BST) Insert(val int) {
    tree.Root = insert(tree.Root, val)
}

func insert(root *Node, val int) *Node {
    if root == nil {
        return &Node{Val:val}
    } else if val <= root.Val {
        root.Left = insert(root.Left, val)
    } else {
        root.Right = insert(root.Right, val)
    }
    return root
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
        return search(root.Right, val)
    } 
    return search(root.Left, val)    
}

// REMOVE
func (tree *BST) Remove(val int) {
    tree.Root = remove(tree.Root, val)
}

func remove(root *Node, val int) *Node {
    if root == nil {
        return nil
    }
    if val < root.Val {
        root.Left = remove(root.Left, val)
    } else if val > root.Val {
        root.Right = remove(root.Right, val)
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
            root.Right = remove(root.Right, min_node.Val)
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

func height(root *Node) int {
    if root == nil {
        return -1
    }
    return max(height(root.Left), height(root.Right)) + 1
}

// DEPTH
func (tree *BST) Depth(val int) int {
    depth := 0
    curr := tree.Root
    for curr != nil {
        if val < curr.Val {
            curr = curr.Left
        } else if val > curr.Val {
            curr = curr.Right
        } else if val == curr.Val {
            return depth
        }
        depth += 1
    }
    return -1
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






