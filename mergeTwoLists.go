/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
* - use two pointers starting from each list 
* - iterate through two list concurrently 
* - compare values and put smaller value into the returning merged list 
* - move pointer of smaller-value list to the next position 
* - until end of either list 
* - the larger list still have remaining values will be merged into the returning list
* O(m + n) for time complexity 
* O(m + n) for space needed 
*/ 
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := new(ListNode)
    p1 := l1
    p2 := l2
    curr := dummyHead 
    
    for (p1 != nil &&  p2 != nil) {
        if (p1.Val < p2.Val) {
            curr.Next = &ListNode{p1.Val, nil}
            p1 = p1.Next
        } else {
            curr.Next = &ListNode{p2.Val, nil}
            p2 = p2.Next
        }
        curr = curr.Next
    }
    
    for (p1 != nil) {
        curr.Next = &ListNode{p1.Val, nil}
        curr = curr.Next
        p1 = p1.Next
    }
    
    for (p2 != nil) {
        curr.Next = &ListNode{p2.Val, nil}
        curr = curr.Next
        p2 = p2.Next
    }
    
    return dummyHead.Next
}
