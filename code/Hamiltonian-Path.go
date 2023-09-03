package main
import (
    "fmt"
)
var p = fmt.Println

/*
    0 ----- 1 ----\ 
     \      \      4
      3 ----- 2 --/

    H cycles:
    0 - 1 - 4 - 2 - 3 - 0
*/

func main() {
    // a un-directed graph implememented using adjacency list
    graph := [][]int{
        []int{1,3},   // 0
        []int{0,2,4}, // 1
        []int{1,3,4},   // 2
        []int{0,2},   // 3
        []int{1,2},   // 4
    }

    Hamiltonian(graph)
}

func Hamiltonian(graph [][]int) {
    visited := map[int]bool{}

    // to store the current valid path
    path := make([]int, len(graph))

    // if the (un-directed) graph have an Hamiltonian cycle
    // then we can start at any arbitrary vertex
    // in this case, I chose vertex 0
    DFS(0, graph, path, visited, 1)

    // Drawback!
    // this backtracking algorithm will produce duplicate paths.
    //
    // 2 Hamiltonion paths considered a same path if they consist of same edges,
    // Ex: these 2 paths are the same
    // [0 1 4 2 3]
    // [0 3 2 4 1]
    // 
    // need to eliminate duplicate paths to find unique path
}

func DFS(vertex int, graph [][]int, path []int, visited map[int]bool, index int) {    
    if visited[vertex] {
        return
    }

    // special case, last index
    // if the last vertex have connection to start vertex which is 0 in this case, 
    // then we found a valid cycle 
    if index == len(graph) {
        for _, v := range graph[vertex] {
            if v == 0 {
                p(path)
                return
            }
        }
        return
    }

    visited[vertex] = true
    for _, v := range graph[vertex] {
        if !visited[v] {
            path[index] = v
            DFS(v, graph, path, visited, index+1)
        }
    }
    delete(visited, vertex)
}











