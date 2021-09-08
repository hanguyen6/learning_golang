/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
/* Two passes 
* First iteration: Move through the list to calculate length of the list: l
* Second iteration: Delete the node at position (l -n)
* O(2 * n) time complexity 
* O(1) space 
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := 0 
    
    curr := head 
    for (curr != nil) {
        l +=1 
        curr = curr.Next
    }
    
    pos := l - n
    
    dummyHead := new(ListNode)
    dummyHead.Next = head 
    
    prev := dummyHead 
    curr = head 
    for (curr != nil) {
        if (pos == 0) {
            prev.Next = curr.Next
        } else {
            prev = curr 
        }
        curr = curr.Next
        pos--
    }
    
    return dummyHead.Next
}


/** One pass 
* Use two pointers, one starts at the beginning of the list 
* One starts at (n) position 
* When second pointer reaches the end of the list 
* first pointer goes to the node where its next node need to be deleted 
* 
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyHead := new(ListNode)
    dummyHead.Next = head 
    
    fast := head 
    slow := dummyHead 
    
    // advance fast pointer to the (n+1) position 
    for (n !=0 && fast != nil) {
        fast = fast.Next
        n--
    }
    
    // start slow pointer while keep fast pointer moving til the end 
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    
    // Delete next node after slow pointer 
    slow.Next = slow.Next.Next
    
    return dummyHead.Next
}
