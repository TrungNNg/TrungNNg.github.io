package main
import (
    "fmt"
)
var p = fmt.Println

func main() {
    mh := MinHeap{[]int{}}
    mh.Push(5)
    mh.Push(2)
    mh.Push(1)
    mh.Push(3)
    mh.Push(4)

    p(mh.arr)

    //mh.Update(1, 0)

    //p(mh.arr)

}

type MinHeap struct {
    arr []int
}

// Push
func (h *MinHeap) Push (val int) {
    h.arr = append(h.arr, val)
    up(h.arr, len(h.arr)-1)
}

func up(arr []int, index int) {
    for index > 0 {
        parentIndex := (index - 1) / 2
        if arr[index] < arr[parentIndex] {
            arr[index], arr[parentIndex] = arr[parentIndex], arr[index]
        }
        index = parentIndex
    }
}

// Pop
// panic if heap is empty
func (h *MinHeap) Pop () int {
    ans := h.arr[0]
    lastIndex := len(h.arr) - 1
    h.arr[0], h.arr[lastIndex] = h.arr[lastIndex], h.arr[0]
    h.arr = h.arr[:lastIndex]
    down(h.arr, 0)
    return ans
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

// Update
func (h *MinHeap) Update (index, val int) {
    h.arr[index] = val
    // h.arr[index] can be smaller than its parent or bigger than it childrens
    parentIndex := (index - 1) / 2
    if h.arr[index] < h.arr[parentIndex] {
        up(h.arr, index)
    } else {
        down(h.arr, index)
    }
}


