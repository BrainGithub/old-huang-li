package main

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(h *ListNode) *ListNode {
    if h == nil || h.Next == nil {
        return h
    }

    var pre * ListNode = nil
    var cur = h
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}

func dfs_reverseList(h *ListNode) *ListNode {
    if h == nil || h.Next == nil {
        return h
    }

    p := dfs_reverseList(h)
    h.Next.Next = p
    h.Next = nil
    return p
}
