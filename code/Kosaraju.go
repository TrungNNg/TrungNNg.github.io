package main
import (
    "fmt"
)
var p = fmt.Println

func main() {
    // a directed graph implememented using adjacency list
    graph := map[int][]int {
        1 : {2},
        2 : {3},
        3 : {1,5},
        4 : {5},
        5 : {6},
        6 : {4},
        7 : {6},
    }

    SCCs := Kosaraju(graph)
    p(SCCs)
}

func Transpose(graph map[int][]int) map[int][]int {
    reverse_graph := map[int][]int{}
    for v, l := range graph {
        for _, vertex := range l {
            if _, ok := reverse_graph[vertex]; !ok {
                reverse_graph[vertex] = []int{}
            }
            reverse_graph[vertex] = append(reverse_graph[vertex], v)
        }
    }
    return reverse_graph
}

func Kosaraju(graph map[int][]int) [][]int {

    // pass 1: build the stack
    stack := []int{}
    visited := map[int]bool{}
    for v := range graph {
        if !visited[v] {
            helper(v, &stack, visited, graph)
        }
    }

    // pass 2: reverse the graph
    reverse_graph := Transpose(graph)
    visited = map[int]bool{}
    SCCs := [][]int{}
    for len(stack) != 0 {
        v := stack[0]
        stack = stack[1:]
        if !visited[v] {
            SCC := []int{}
            helper_2(v, stack, visited, reverse_graph, &SCC)
            SCCs = append(SCCs, SCC)
        }
    }

    return SCCs
}

// DFS on reverse graph and build SCC
func helper_2(vertex int, stack []int, visited map[int]bool, reverse_graph map[int][]int, SCC *[]int) {
    if visited[vertex] {
        return
    }
    visited[vertex] = true
    for _, v := range reverse_graph[vertex] {
        helper_2(v, stack, visited, reverse_graph, SCC)
    }
    *SCC = append(*SCC, vertex)
}

// DFS on original graph
func helper(vertex int, stack *[]int, visited map[int]bool, graph map[int][]int) {
    if visited[vertex] {
        return
    }
    visited[vertex] = true
    for _, v := range graph[vertex] {
        helper(v, stack, visited, graph)
    }
    *stack = append([]int{vertex}, (*stack)...)
}











