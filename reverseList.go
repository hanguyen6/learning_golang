/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
* https://leetcode.com/problems/reverse-linked-list
* - Use a pointer to store the previous node 
* - Use a temp pointer p to store the next node 
* - While iterate through the list 
* - Redirect the next of current node to previous node 
* - Move previous node become current node 
* - Move current next become the next temp node 
* - O(n) for time complexity
*/
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head 
    for (curr != nil) {
        nextTemp := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextTemp
    }
    return prev
}
