package main

import (
    "fmt"
)
var p = fmt.Println
func main() {
    n := 2
    arr := [n]int{}
    p(arr)
}

// generate all subset
func subsets(nums []int) [][]int {
    ans := [][]int{}
    //m := map[int]bool{}
    subsets_helper(&ans, nums, []int{})
    return ans
}

func subsets_helper(ans *[][]int, remain, current []int) {
    if len(remain) == 0 {
        *ans = append(*ans, append([]int{}, current...))
        return
    }
    current = append(current, remain[0])
    subsets_helper(ans, remain[1:], current)
    current = current[:len(current)-1]
    subsets_helper(ans, remain[1:], current)
}

// permute
func permute(nums []int) [][]int {
    ans := [][]int{}
    m := map[int]bool{}
    permute_helper(&ans, nums, []int{}, m)
    return ans
}

func permute_helper(ans *[][]int, remain, current []int, chosen map[int]bool) {
    if len(current) == len(remain) {
        *ans = append(*ans, append([]int{}, current...))
        return
    }
    for i := range remain {
        if !chosen[remain[i]] {
            current = append(current, remain[i])
            chosen[remain[i]] = true
            permute_helper(ans, remain, current, chosen)
            current = current[:len(current)-1]
            chosen[remain[i]] = false
        }
    }
}

// combination
func combine(n int, k int) [][]int {
    ans := [][]int{}
    remain := []int{}
    for i := 1; i <= n; i++ {
        remain = append(remain, i)
    }
    combine_helper(&ans, remain, []int{}, k, 0)
    return ans
}

func combine_helper(ans *[][]int, remain, current []int, k int, index int) {
    if len(current) == k {
        *ans = append(*ans, append([]int{}, current...))
    } else {
        for i := index; i < len(remain); i++ {
            current = append(current, remain[i])
            combine_helper(ans, remain, current, k, i+1)
            current = current[:len(current)-1]
        }
    }
}




