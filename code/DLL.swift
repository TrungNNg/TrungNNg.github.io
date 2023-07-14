import Foundation

class DoublyLinkedList {
    class Node {
        var value: Int
        var next: Node?
        var previous: Node?
        init(_ value: Int) {
            self.value = value
        }
        init(_ value: Int,_ next: Node?,_ previous: Node?) {
            self.value = val
            self.next = next
            self.previous = previous
        }
    }

    var head: Node?
    var tail: Node?

    // add to end of list
    func append(_ val: Int) {
        if head == nil {
            head = Node(val)
            tail = head
        } else {
            tail.next = Node(val,nil,tail)
            tail = tail.next
        }
    }

    // remove all occuence
    func remove_all(_ val: Int) {

    }
    
    // reverse list
    func reverse() {

    }

}



