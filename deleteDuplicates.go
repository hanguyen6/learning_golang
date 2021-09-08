/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// O(n) time complexity and O(1) space needed
func deleteDuplicates(head *ListNode) *ListNode {
    dummyHead := new(ListNode)
    dummyHead.Next = head 
    
    curr := head 
    
    for (curr != nil && curr.Next != nil)  {
        if (curr.Val == curr.Next.Val) {
            // Skip next node if found duplicate
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    
    return dummyHead.Next
    
}
