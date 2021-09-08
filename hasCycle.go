/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
* Keep track of visited node in hashmap  
* When node visited twice then there is a cycle 
* O(n) to space complexity 
* O(n) for time complexity 
*/
/*
func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]bool)
    
    curr := head
    for curr != nil {
        _, found := m[curr]
        if (found) {
            return true
        } else {
            m[curr] = true
        }
        curr = curr.Next
    }
    
    return false
}
*/

/*
* Use two pointers, one moving fast, one moving slow, starting from head 
* If fast pointer meet slow pointer before slow pointer comes to the end 
* there is cycle  
* O(n) for time complexity 
* O(1) for space needed
*/

func hasCycle(head *ListNode) bool {
    if (head == nil) { 
        return false
    }
    slow := head 
    fast := head.Next
    
    for (slow != nil && fast != nil && fast.Next != nil) {
        fast = fast.Next.Next 
        slow = slow.Next
        
        if (slow == fast) {
            return true
        }
    }
    
    return false
}
