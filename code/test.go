package main

import (
    "fmt"
)
var p = fmt.Println

func main() {
    h := Constructor(7, []int{-10,1,3,1,4,10,3,9,4,5,1})
    
    p(h.Add(3))
    p(h.Add(2))
    p(h.Add(3))
    p(h.Add(1))
    p(h.Add(2))
    p(h.Add(4))
    p(h.Add(5))
    p(h.Add(5))
    p(h.Add(6))
    p(h.Add(7))
    p(h.Add(7))
    p(h.Add(8))
    p(h.Add(2))
    p(h.Add(3))
    p(h.Add(1))
    p(h.Add(1))
    p(h.Add(1))
    p(h.Add(10))
    p(h.Add(11))
    p(h.Add(5))
    p(h.Add(6))
    p(h.Add(2))
    p(h.Add(4))
    p(h.Add(7))
    p(h.Add(8))
    p(h.Add(5))
    p(h.Add(6))
   
}

type KthLargest struct {
    arr []int
    k int
}
func Constructor(k int, nums []int) KthLargest {
    heap := KthLargest{arr: []int{}, k:k}
    for _, v := range nums {
        heap.arr = append(heap.arr, v)
        up(heap.arr, len(heap.arr) - 1)
    }
    fix(heap.arr, 0)
    return heap
}

func fix(arr []int, index int) {
    if index > len(arr) { 
        return 
    } else {
        lc := index * 2 + 1
        rc := index * 2 + 2
        if rc < len(arr) && arr[rc] < arr[index] {
            arr[rc], arr[index] = arr[index], arr[rc]
        } else {
            if lc < len(arr) && arr[lc] < arr[index] {
                arr[lc], arr[index] = arr[index], arr[lc]
            }
        }
        fix(arr, lc)
        fix(arr, rc)
    }
}

func (this *KthLargest) Add(val int) int {
    this.arr = append(this.arr, val)
    up(this.arr, len(this.arr) - 1)
    for len(this.arr) > this.k {
        this.Pop()
    }
    return this.arr[0]
}

func up(arr []int, index int) {
    parentIndex := (index - 1) / 2
    for parentIndex >= 0 {
        if arr[index] < arr[parentIndex] {
            arr[index], arr[parentIndex] = arr[parentIndex], arr[index]
        }
        if parentIndex == 0 {
            break
        }
        index = parentIndex
        parentIndex = (index - 1) / 2
    }
}

func (this *KthLargest) Pop() { 
    last_index := len(this.arr) - 1
    this.arr[0], this.arr[last_index] = this.arr[last_index], this.arr[0]
    this.arr = this.arr[:last_index]
    down(this.arr, 0)
}

func down(arr []int, index int) {
    leftChild := index * 2 + 1
    for leftChild <= len(arr) - 1 {
        rightChild := leftChild + 1
        if rightChild <= len(arr) - 1 && arr[rightChild] < arr[leftChild] && arr[rightChild] < arr[index] {
            arr[rightChild], arr[index] = arr[index], arr[rightChild]
            index = rightChild
        } else {
            if arr[leftChild] < arr[index] {
                arr[leftChild], arr[index] = arr[index], arr[leftChild]
            }
            index = leftChild
        }
        leftChild = index * 2 + 1
    }
}


















