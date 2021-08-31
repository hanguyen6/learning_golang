package main

import "fmt" 

// https://leetcode.com/problems/add-two-numbers
type ListNode struct {
  Val int
  Next *ListNode 
}

func (l *ListNode) Display() {
	n := l  
	for n != nil {
		fmt.Printf("%v -> ", n.Val)
		n = n.Next
	}
	fmt.Println("nil")
}

func (l *ListNode) PushBack(n *ListNode) {
  curr := l 
  for curr.Next != nil {
    curr = curr.Next
  }
  curr.Next = n
}

// O(max(m,n)) for time complexity 
// O(max(m,n)) for space needed
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := new(ListNode)
    p := l1 
    q := l2 
    curr := dummyHead 
    carry := 0 
    var x, y int 
    
    for (p != nil || q != nil) {
        if (p != nil) {
            x = p.Val
        } else {
            x = 0
        }
        if (q != nil) {
            y = q.Val
        } else {
            y = 0 
        }
        
        sum := x + y + carry 
        carry = sum / 10 
        curr.Next = &ListNode{sum % 10, nil}
        curr = curr.Next 
        if (p != nil) {
            p = p.Next
        }
        
        if (q != nil) {
            q = q.Next       
        }
    }
    if (carry > 0) {
        curr.Next =  &ListNode{carry, nil}
    }
    return dummyHead.Next
}

func main() {
  fmt.Println("Sum two numbers stored in the linked list in reversed order ..  ")  

  l1 := &ListNode{5, nil}
  n := &ListNode{7, nil}
  l1.PushBack(n)
  l1.Display()
  
  l2 := &ListNode{4, nil}
  m := &ListNode{9, nil}
  l2.PushBack(m)
  l2.Display()
  
  fmt.Println("Result:")
  l3 := addTwoNumbers(l1, l2)
  l3.Display()
  
}
