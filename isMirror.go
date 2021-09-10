/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/** -- Recursive Solution -- 
* Compare two trees with same root. They are mirrored if: 
* - values are the same 
* - leftSubTree of r1 is the mirror of rightSubTree of r2 
* - rightSubTree of r1 is the mirror of leftSubTree of r2 
O(N) time complexity since we have to travel the entire tree 
O(n) for space complexity: the recursive call is bound to the height of the tree which is O(N) in the worse case 
*/


func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root)
}


func isMirror(r1 *TreeNode, r2 *TreeNode) bool {
    if (r1 == nil && r2 == nil) {
        return true
    }
    
    if (r1 == nil || r2 == nil) {
        return false
    }
    
    return r1.Val == r2.Val && isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
}


/** -- Iterative solution --  
*  Use an queue to store nodes in the tree 
* Initially, the queue contain root and root 
* While queue is not empty, pop two elements from the queue 
- If both are null, continue 
- Either is null, return, false
- Compare their values if none is null. return false if values are not equal 
- Otherwise, insert left and right subtree of two nodes in opposite order 
O(N) for time and space complexity 
*/

func isSymmetric(r1 *TreeNode) bool {
    
    q := make([]*TreeNode, 0)
    q = append(q, r1)
    q = append(q, r1)
    
    for len(q) > 0 {
        t1 := q[0]
        t2 := q[1]
        q = q[2:]
        
        if (t1 == nil && t2 == nil) {
            continue
        }
        
        if (t1 == nil || t2 == nil) {
            return false
        }
        
        if (t1.Val != t2.Val) {
            return false 
        }
        
        q = append(q, t1.Left)
        q = append(q, t2.Right)
        q = append(q, t1.Right)
        q = append(q, t2.Left)
    }
    
    return true
}
