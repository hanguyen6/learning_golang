/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 /**
 * -- Recursive Solution
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
   
    if (root == nil) {
        return false
    }
    
    targetSum = targetSum - root.Val
    
    if (root.Left == nil && root.Right == nil) {
        return targetSum == 0
    } else {
        return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
    }
}

/* -- Iterative Solution --
* Visit each node with DFS pre-order strategy and updating remaining sum for each visit 
* Use a stack containing the nodes with their remaining sum 
* Pop the current node out of stack and return true if remaining sum = 0 and we on leaf node
* Otherwise push the child nodes and corresponding remaining sum to stack 
* O(N) for time & space complexity 
*/

type TreeNodeWithSum struct {
    Node *TreeNode
    Rsum int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if (root == nil) {
        return false
    }
    
    stack := make([]*TreeNodeWithSum, 0)
    stack = append(stack, &TreeNodeWithSum{root, targetSum - root.Val})
    
    for len(stack) > 0 {
        curr := stack[(len(stack) - 1)]
        stack = stack[0:len(stack)-1]
        
        // Check remaining sum if curr is leaf node 
        if (curr.Node.Left == nil && curr.Node.Right == nil && curr.Rsum == 0) {
            return true
        }
        
        // Else push child nodes into stack 
        
        if (curr.Node.Right != nil) {
            rNode := &TreeNodeWithSum{curr.Node.Right, curr.Rsum - curr.Node.Right.Val}
            stack = append(stack, rNode)
        }
        
        if (curr.Node.Left != nil) {
            lNode := &TreeNodeWithSum{curr.Node.Left, curr.Rsum - curr.Node.Left.Val}
            stack = append(stack, lNode)
        }
        
    }
    
    return false
}
