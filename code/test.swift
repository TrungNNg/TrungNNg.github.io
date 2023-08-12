import Foundation

let arr : [Int] = [1,3,2,4]
let prefix = create_prefix(arr)
print(prefix)
func create_prefix(_ arr : [Int]) -> [Int] {
    var prefix = Array(repeating: 0, count: arr.count+1)
    for i in 0..<arr.count {
        prefix[i+1] = prefix[i] + arr[i]
    }
    return prefix
}

