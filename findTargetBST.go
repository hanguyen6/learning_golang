/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/* 
* Use a hashTable to keep track of remainingSum = target Value - current Value 
* As traversing through the tree, 
    - If remainingSum in the hashTable, return true
    - add remainingSum to the hashTable if remainingSum not found 
* O(N) for time and space complexity 
*/
func findTarget(root *TreeNode, k int) bool {
    m := make(map[int]bool, 0)
    return find(root, k, m)
}

func find(root *TreeNode, k int, m map[int]bool) bool {
    if (root == nil) {
        return false
    }
    
    _, found := m[k - root.Val]
    if (found) {
        return true
    } else {
        m[root.Val] = true
    }
    
    return find(root.Left, k, m) || find(root.Right, k, m)
}
