package main
import (
    "fmt"
)
var p = fmt.Println

func main() {
    // a un-directed graph implememented using adjacency list
    graph := map[int][]int {
        0 : {1,2},
        1 : {0,2},
        2 : {0,1,3},
        3 : {2,4},
        4 : {3,5,6},
        5 : {4,6,7},
        6 : {4,5},
        7 : {5},
    }
    a_points := Tarjan(graph)
    fmt.Print("articulation points: ")
    for k := range a_points {
        fmt.Print(k, " ")
    }
    p()
}

func Tarjan(graph map[int][]int) map[int]bool {
    visited   := map[int]bool{}
    disc      := map[int]int{}
    low       := map[int]int{}
    ans       := map[int]bool{}
    disc_time := 0
    // in this example, -1 will represent no parent
    DFS(3,-1, &disc_time, graph, disc, low, visited, ans)

    // to find bridges (edges which removal increase number of connected components)
    // an edge is a bridge if both 2 conditions are true
    // 1. it connect to an articulation point
    // 2. the articulation point has lower low than its children
    //    so the child subgraph does not have back-edge to the articulation point or its parent
    bridges := [][2]int{}
    for k := range ans {
        for _, v := range graph[k] {
            if low[k] < low[v] {
                bridges = append(bridges, [2]int{k,v})
            }
        }
    }
    p("bridges:", bridges)
    return ans
}

func DFS(vertex, parent int, disc_time *int, graph map[int][]int, 
         disc, low map[int]int, visited map[int]bool, 
         ans map[int]bool) {

    disc[vertex] = *disc_time
    low[vertex]  = *disc_time
    *disc_time = *disc_time + 1
    visited[vertex] = true
    child_count := 0
    for _, v := range graph[vertex] {
        if v == parent { continue }

        if !visited[v] {
            DFS(v, vertex, disc_time, graph, disc, low, visited, ans)
            child_count += 1
        }

        if low[v] < low[vertex] {
            low[vertex] = low[v]
        }

        // check if its parent is an articulation point
        // handle special case for root first
        if parent == -1 {
            if child_count >= 2 {
                ans[vertex] = true
            }
        } else if low[v] >= disc[vertex] {
                ans[vertex] = true
        }
    }
}







