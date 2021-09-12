/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/* 
* Starting from root, compare current value vs target val 
* If current Value < target Val -> Insert into right subTree
* Else Insert into left subTree  
* Return root 
* O(log(N)) for time complexity 
* O(log(N)) for space complexity 
*/

func insertIntoBST(root *TreeNode, val int) *TreeNode {    
     
    if (root == nil) {
        return &TreeNode{val, nil, nil}
    } 
    
    if (root.Val < val) {
        root.Right = insertIntoBST(root.Right, val)
    } else {
        root.Left = insertIntoBST(root.Left, val)
    }

    return root
    
}

* -- Iterative solution -- 
* O(log(N)) for time complexity 
* O(1) space need
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {     
    curr := root
    for (curr != nil) {
        if (curr.Val > val) {
            if (curr.Left == nil) {
                curr.Left = &TreeNode{val, nil, nil}
                return root
            } else {
                curr = curr.Left
            }
        } else {
            if (curr.Right == nil) {
                curr.Right = &TreeNode{val, nil, nil}
                return root
            } else {
                curr = curr.Right
            }
        }
    }
    return &TreeNode{val, nil, nil}
    
    
}
