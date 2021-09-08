/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/** 
* Use dummyHead / sentinel node to avoid headLess situation 
* when deleting element(s) in the head of linked list 
* Use two pointer prev and curr to move through the list 
* When curr pointer point to node to be deleted 
* - Set next pointer of prev to the node next to the one to delete 
* O(n) time complexity 
* O(1) space need 
*/
func removeElements(head *ListNode, val int) *ListNode {
    
    dummyHead := new(ListNode)
    dummyHead.Next = head 
    
    prev := dummyHead 
    curr := head 
    
    for curr != nil {
        if (curr.Val == val) {
            prev.Next = curr.Next
        } else {
            prev = curr
        }
        curr = curr.Next
    }
    
    return dummyHead.Next
    
}
