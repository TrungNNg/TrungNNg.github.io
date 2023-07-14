import Foundation

let list = DoublyLinkedList()
list.append(1)
list.append(2)
list.append(3)
list.append(4)
list.append(5)

list.remove_all(5)

list.toString()

class DoublyLinkedList {
    class Node {
        var value: Int
        var next: Node?
        var previous: Node?
        init(_ value: Int) {
            self.value = value
        }
        init(_ value: Int,_ next: Node?,_ previous: Node?) {
            self.value = value
            self.next = next
            self.previous = previous
        }
    }

    // init var
    var head: Node?
    var tail: Node?

    // add to end of list
    func append(_ val: Int) {
        if head == nil {
            head = Node(val)
            tail = head
        } else {
            tail?.next = Node(val,nil,tail)
            tail = tail?.next
        }
    }

    // remove all
    func remove_all(_ val: Int) {
        var curr = head
        while curr != nil {
            if curr!.value == val {
                curr?.previous?.next = curr!.next
                curr?.next?.previous = curr!.previous
            }
            curr = curr?.next
        }
    }
    
    // reverse list
    func reverse() {
        var arr :[Int] = []
        var curr = head
        while curr != nil {
            arr.append(curr!.value)
            curr = curr?.next
        }
        curr = tail
        while curr != nil {
            curr?.value = arr.removeLast()
            curr = curr?.previous
        }
    }

    // print list
    func toString() {
        var curr = head
        while curr != nil {
            print(curr!.value)
            curr = curr?.next
        }
    }
}



