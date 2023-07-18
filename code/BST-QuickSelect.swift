import Foundation


var arr: [Int] = [1, 3, 2]

print(QuickSelect(arr, 2))

func BinarySearch(_ arr: [Int],_ target: Int) -> Int {
    var left = 0, right = arr.count-1
    var mid = -1
    while left <= right {
        mid = (right - left) / 2 + left
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

// return k largest from unsorted array
func QuickSelect(_ arr: [Int],_ k: Int) -> Int {
    var arr = arr
    var left = 0, right = arr.count - 1
    let k = arr.count - k
    var p = partition(&arr,left,right)
    while p != k {
        if p < k {
            left = p + 1
            p = partition(&arr,left,right)
        }
        if p > k {
            right = p - 1
            p = partition(&arr,left, right)
        }
    }
    return arr[p]
}

func partition(_ arr: inout [Int],_ start: Int,_ end: Int) -> Int {
    let pivot = arr[end]
    var pIndex = start
    for i in start..<end {
        if arr[i] < pivot {   // if the pivot has duplicate. all duplicates will go right
            arr.swapAt(i, pIndex)
            pIndex += 1
        }
    }
    arr.swapAt(pIndex, end)
    return pIndex
}








