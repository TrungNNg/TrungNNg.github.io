package main

import (
    "fmt"
)
var p = fmt.Println
func main() {
    p("hello")

    //matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}

    dfs(matrix,0,0)

}


func dfs(matrix [][]int, row, col int) {
    visited := map[[2]int]bool{[2]int{row,col}:true}
    stack := [][2]int{[2]int{row,col}}
    for len(stack) != 0 {
        curr := stack[len(stack)-1]
        row = curr[0]
        col = curr[1]
        stack = stack[:len(stack)-1]
        p(matrix[row][col])

        // row - 1, col
        if row - 1 >= 0 && !visited[[2]int{row-1,col}] {
            stack = append(stack, [2]int{row-1, col})
            visited[[2]int{row-1, col}] = true
        }
        // row + 1, col
        if row + 1 < len(matrix) && !visited[[2]int{row+1,col}] {
            stack = append(stack, [2]int{row+1, col})
            visited[[2]int{row+1, col}] = true
        }
        // row, col - 1
        if col - 1 >= 0 && !visited[[2]int{row,col-1}] {
            stack = append(stack, [2]int{row, col-1})
            visited[[2]int{row, col-1}] = true
        }
        // row, col + 1
        if col + 1 < len(matrix[curr[0]]) && !visited[[2]int{row,col+1}] {
            stack = append(stack, [2]int{row, col+1})
            visited[[2]int{row, col+1}] = true
        }
    }
}






