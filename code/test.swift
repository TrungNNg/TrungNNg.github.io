import Foundation


var arr = [2,5,1,1,4,3]

bubble_sort(&arr)

print(arr)

func bubble_sort(_ arr: inout [Int]) {
    for i in 0...arr.count-1 {
        for j in 0...arr.count-2 {
            if arr[i] < arr[j] {
                arr.swapAt(i,j)
            }
        }
    }
}



