import Foundation

var arr:[Int] = [2,4,1,2,5,3]
print(arr)
//bubble_sort(&arr)
//selection_sort(&arr)
//insertion_sort(&arr)

arr = merge_sort(arr)

print(arr)

func merge_sort(_ arr: [Int]) -> [Int] {
    var arr = arr
    if arr.count <= 1 {
        return arr
    }
    let mid = arr.count / 2
    let left = merge_sort( [Int](arr[0..<mid]) )
    let right = merge_sort( [Int](arr[mid..<arr.count]) )
    merge(&arr, left, right)

    func merge(_ ans: inout [Int], _ left:[Int],_ right:[Int]) {
        var k = 0, a = 0, b = 0
        while a < left.count && b < right.count {
            if left[a] <= right[b] {
                ans[k] = left[a]
                a += 1
            } else if right[b] < left[a] {
                ans[k] = right[b]
                b += 1
            }
            k += 1
        }
        while a < left.count {
            ans[k] = left[a]
            a += 1
            k += 1
        }
        while b < right.count {
            ans[k] = right[b]
            b += 1
            k += 1
        }
    }

    return arr
}




func quick_sort(_ arr: inout [Int]) {
}
func optimize_quick_sort(_ arr: inout [Int]) {

}


func bubble_sort(_ arr: inout [Int]) {
    for _ in 0...arr.count-2 {
        for j in 0...arr.count-2 {
            if arr[j] > arr[j+1] {
                arr.swapAt(j,j+1)
            }
        }
    }
}

func insertion_sort(_ arr: inout [Int]) {
    for i in 1...arr.count-1 {
        let temp = arr[i]
        var hole = i
        while hole > 0 && arr[hole-1] > temp {
            arr[hole] = arr[hole-1]
            hole -= 1
        }
        arr[hole] = temp
    }
}

func selection_sort(_ arr: inout [Int]) {
    for i in 0...arr.count-2 {
        var mIndex = i
        for j in i+1...arr.count-1 {
            if arr[j] < arr[mIndex] {
                mIndex = j
            }
        }
        arr.swapAt(i,mIndex)
    }
}

