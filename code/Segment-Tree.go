package main
import (
    "fmt"
)
var p = fmt.Println
func main() {
    stree := segment_tree([]int{-1,3,4,0,2,1})
    ans := min_range(stree, 2,4, 0, 5, 0)
    p(ans)
   // p(min_range(stree,0,4,0,5,0)) // 0 ?
}

func segment_tree(arr []int) []int {
    stree := make([]int, np2(len(arr)) * 2 - 1)
    build_tree(arr, stree, 0, len(arr)-1, 0)
    return stree
}

func build_tree(arr, stree []int, l, h, pos int) {
    if l == h {
        stree[pos] = arr[l]
        return
    }
    m := (l + h) / 2
    build_tree(arr, stree, l, m, pos*2+1)
    build_tree(arr, stree, m+1, h, pos*2+2)
    stree[pos] = min(stree[pos*2+1], stree[pos*2+2])
}

func min_range(stree []int, ql, qh, l, h, pos int) int {
    // if query range >= current range return value of current range
    if ql <= l && qh >= h {
        return stree[pos]
    } else if ql > h || qh < l {
        // if query range is outside current range return max
        return int(^uint(0) >> 1)
    }
    // if query range partial overlap, compare with left and right subtree
    m := (h + l) / 2
    return min( min_range(stree, ql, qh, l, m, pos*2+1),
                min_range(stree, ql, qh, m+1, h, pos*2+2) )
}

func min(a, b int) int {
    if a < b { return a}
    return b
}

func np2(a int) int {
    ans := 1
    for ans < a {
        ans *= 2
    }
    return ans
}

