package main
import (
    "fmt"
)
var p = fmt.Println

func main() {
    // a un-directed graph implememented using adjacency matrix
    graph := [][]int{
        []int{0,1,1,0,1},
        []int{1,0,1,1,1},
        []int{1,1,0,1,0},
        []int{0,1,1,0,1},
        []int{1,1,0,1,0},
    }

    Hamiltonian(graph)
}

// return all hamiltonian cycle in the graph
func Hamiltonian(graph [][]int) {
    visited := map[int]bool{}

    // to store the current valid path
    path := make([]int, len(graph))

    // if the (un-directed) graph have an Hamiltonian cycle
    // then we can start at any arbitrary vertex
    // in this case, I chose vertex 0
    DFS(0, graph, path, visited, 1)

}

func DFS(vertex int, graph [][]int, path []int, visited map[int]bool, index int) {    
    if visited[vertex] {
        return
    }

    // special case, last index
    // if the last vertex have connection to start vertex which is 0 in this case, 
    // then we found a valid cycle 
    if index == len(graph) {
        if graph[vertex][0] == 1 {
            p("valid path:", path)
        }
        return
    }

    visited[vertex] = true
    for v, connect := range graph[vertex] {
        if connect == 1 && !visited[v] {
            path[index] = v
            DFS(v, graph, path, visited, index+1)
        }
    }
    delete(visited, vertex)
}











