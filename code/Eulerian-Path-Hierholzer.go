package main
import (
    "fmt"
)
var p = fmt.Println

func main() {
    // a directed graph implememented using adjacency list
    graph := map[int][]int {
        1 : {2,3},
        2 : {4,4},
        3 : {1,2,5},
        4 : {3,6},
        5 : {6},
        6 : {3},
    }
    Eulerian_path := Hierholzer(graph)
    p(Eulerian_path)
}

func Hierholzer(graph map[int][]int) []int {
    // in and out-degree array, start from index 1
    in  := make([]int, len(graph)+1)
    out := make([]int, len(graph)+1)

    for k, l := range graph {
        out[k] = len(l)
        for _, v := range l {
            in[v] += 1
        }
    }

    // check if there exist an Eulerian path in the graph
    // 3 conditions for an directed graph to has a Eulerian path:
    // - at most 1 vertex with 1 extra out edge (start vertex)
    // - at most 1 vertex with 1 extra in edge (end vertex)
    // - all other vertices have equal amount of in-out degree 
    start_vertex := 0
    end_vertex := 0
    for i := 1; i < len(in); i++ {
        if in[i] - out[i] > 1 || out[i] - in[i] > 1 {
            // no valid path
            return nil
        } else if in[i] - out[i] == 1 { // start vertex
            start_vertex += 1
        } else if out[i] - in[i] == 1 { // end vertex
            end_vertex += 1
        }
        // don't need to check in[i] == out[i] because there are only 3 cases 
        // ( |in - out| > 1), ( |in - out| == 1), (in == out)
        // we already check the first 2 cases, so the final case here is the (in == out)
    }

    if !((start_vertex == 0 && end_vertex == 0) ||  // Eulerian circuit case, also Eulerian path case
       (start_vertex == 1 && end_vertex == 1)) {   // Eulerian path
        // no valid path
        return nil
    }

    // use in and out to chose start vertex
    for i := 1; i < len(in); i++ {
        if out[i] - in[i] == 1 {
            start_vertex = i
            break
        }
        if out[i] > 0 {
            start_vertex = i
        }
    }

    // DFS
    E_path := []int{}
    DFS(start_vertex, graph, &E_path, out)
    return E_path
}

// This function modify original graph (remove visited edge) which is often un-acceptable.
// One way to fix this is to copy the graph, or use a map to mark visited edges
func DFS(vertex int, graph map[int][]int, E_path *[]int, out []int) {
    for len(graph[vertex]) > 0 {
        v := graph[vertex][0]
        graph[vertex] = graph[vertex][1:]
        DFS(v, graph, E_path, out)
    }
    *E_path = append([]int{vertex}, (*E_path)...)
}
















