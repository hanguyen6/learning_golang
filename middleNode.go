/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* 
* https://leetcode.com/problems/middle-of-the-linked-list/
* Use fast and slow pointer 
* fast pointer moves double speed when compared with slow pointer 
* when fast pointer reaches to the end, the slow pointer should be in the middle 
* O(n) time complexity 
* O(1) space 
*/
func middleNode(head *ListNode) *ListNode {
    slow := head
    fast := head 
    for (fast != nil && fast.Next != nil) {
        fast = fast.Next
        fast = fast.Next
        slow = slow.Next
    }
    return slow
}
