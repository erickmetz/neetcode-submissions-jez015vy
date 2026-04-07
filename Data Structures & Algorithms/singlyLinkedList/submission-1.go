type Node struct {
    val int
    next *Node
}

type LinkedList struct {
    head *Node
    len int
}

func NewLinkedList() *LinkedList {
    ll := &LinkedList{
        head: nil,
        len: 0,
    }

    return ll
}

func (ll *LinkedList) Get(index int) int {
    if index > ll.len-1 {
        return -1
    }
    n := ll.head
    idx := 0
    for n != nil {
        if idx == index {
            return n.val
        }
        idx++
        n = n.next
    }

    return -1
}

func (ll *LinkedList) InsertHead(val int) {
    oldHead := ll.head
    newHead := &Node{
        val: val,
        next: oldHead,
    }
    ll.head = newHead
    ll.len += 1
}

func (ll *LinkedList) InsertTail(val int) {
    newTail := &Node{
        val: val,
        next: nil,
    }
    if ll.head == nil {
        ll.head = newTail
        ll.len++
        return
    }
    n := ll.head
    for n != nil {
        if n.next == nil {
            n.next = newTail
            ll.len += 1
            return
        }
        n = n.next
    }
}

func (ll *LinkedList) Remove(index int) bool {
    if index < 0 || index >= ll.len {
        return false
    }

    var lastNode *Node

    n := ll.head
    idx := 0

    for n != nil {
        if idx == index {
            if idx == 0 {
                ll.head = ll.head.next
            } else {
                lastNode.next = n.next
            }
            ll.len -= 1
            return true
        }
        idx++
        lastNode = n
        n = n.next
    }

    return false
}

func (ll *LinkedList) GetValues() []int {
    res := make([]int, 0)

    n := ll.head
    for n != nil {
        res = append(res, n.val)
        n = n.next
    }

    return res
}
