package main
import (
    "fmt"
    "sort"
)
var p = fmt.Println

type Edge struct {
    fv int  // first vertex
    sv int  // second vertex
    w  int  // weight
}

func main() {

    // un-directed and weighted graph represented using edgelist
    graph := []Edge { 
        Edge{0,1,1},
        Edge{0,2,2},
        Edge{1,2,3},
        Edge{1,3,2},
        Edge{2,4,1},
        Edge{3,4,5},
    }
    p(graph)

    // sort the edgelist, or we can also use priority queue
    sort.Slice(graph, func (i, j int) bool {
        return graph[i].w < graph[j].w
    })
    p(graph)

}

