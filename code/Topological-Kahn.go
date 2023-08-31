package main
import (
    "fmt"
)
var p = fmt.Println

func main() {
    // a directed graph implememented using adjacency list
    graph := map[int][]int {
        0 : {1,2},
        1 : {2,3},
        2 : {4},
        3 : {4},
        4 : {},
    }

    p(graph) 

    //p(TopologySort(graph))
    p(Kahn(graph))
}

// DFS
func TopologySort(graph map[int][]int) []int {
    visited := map[int]bool{}
    ans := []int{}
    for k, _ := range graph {
        if !visited[k] {
            T_helper(&ans, visited, k, graph)
        }
    }
    // to check if the sort successful compare ans size and number of vertex
    return ans
}

func T_helper(ans *[]int, visited map[int]bool, vertex int, graph map[int][]int) {
    if visited[vertex] {
        return
    }
    visited[vertex] = true
    for _,v := range graph[vertex] {
        T_helper(ans, visited, v, graph)
    }
    *ans = append([]int{vertex}, (*ans)...)
}

// Kahn's algorithm
func Kahn (graph map[int][]int) []int {
    in_degree := map[int]int{}
    for k := range graph {
        in_degree[k] = 0
    }
    for _, l := range graph {
        for _, v := range l {
            in_degree[v] += 1
        }
    }

    queue := []int{}
    for k, v := range in_degree {
        if v == 0 {
            queue = append(queue, k)
        }
    }

    ans := []int{}
    for len(queue) != 0 {
        vertex := queue[0]
        queue = queue[1:]
        for _, v := range graph[vertex] {
            in_degree[v] -= 1
            if in_degree[v] == 0 {
                queue = append(queue, v)
            }
        }
        ans = append(ans, vertex)
    }
    // to check if the sort successful compare ans size and number of vertex
    return ans
}























