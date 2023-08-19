package main
import (
    "fmt"
)
var p = fmt.Println
func main() {
    stree := segment_tree([]int{-1,3,4,0,2,1})
    ans := min_range(stree, 2,4, 0, 5, 0)
    p(ans) // 0
    //p(min_range(stree,0,4,0,5,0)) // -1 ?
}

func segment_tree(arr []int) []int {
    segtree := make([]int, np2(len(arr)) * 2 - 1)
    build_tree(arr, segtree,0,len(arr)-1,0)
    return segtree
}

func build_tree(arr, segtree []int, l, r, pos int) {
    if l == r {
        segtree[pos] = arr[l]
        return
    }
    m := (l + r) / 2
    build_tree(arr, segtree, l, m, pos*2+1)
    build_tree(arr, segtree, m+1, r, pos*2+2)
    segtree[pos] = min(segtree[pos*2+1], segtree[pos*2+2])
}

func min_range(segtree []int, ql, qr , l, r, pos int) int {
    if ql <= l && qr >= r {
        return segtree[pos]
    }
    if ql > r || l > qr {
        return int(^uint(0) >> 1)
    }
    m := (l + r) / 2
    return min(min_range(segtree, ql, qr, l, m, pos*2+1), min_range(segtree, ql, qr, m+1, r, pos*2+2))
}


func min(a, b int) int {
    if a < b { return a }
    return b
}

func np2(a int) int {
    total := 1
    for i := 1; total < a; i++ { 
        total *= 2
    }
    return total
}


















