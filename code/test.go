package main

import (
    "fmt"
    "container/heap"
)
var p = fmt.Println

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
    old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
    //h := &MaxHeap{1,2,5,3,6}
    arr := []int{1,2,5,3,6}
    var h MaxHeap
    h = arr
    heap.Init(&h)
    heap.Push(&h, 3)
	fmt.Printf("maximum: %d\n", h[0])
	for h.Len() > 0 {
        z := heap.Pop(&h)
        fmt.Printf("%T", z)
	}
}



