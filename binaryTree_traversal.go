/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
/**
* -- Recursive Solutions ---  
**/
func preorderTraversal(root *TreeNode) []int {
    if (root == nil) {
        return []int{}
    } else {
        l1 := append([]int{root.Val}, preorderTraversal(root.Left)...)
        return append(l1, preorderTraversal(root.Right)...)
    }
}

func inorderTraversal(root *TreeNode) []int {
    if (root == nil) {
        return []int{}
    } else {
        l1 := append(inorderTraversal(root.Left), []int{root.Val}...)
        return append(l1, inorderTraversal(root.Right)...)
    }
}

func postorderTraversal(root *TreeNode) []int {
    if (root == nil) {
        return []int{}
    } else {
        l1 := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
        return append(l1, []int{root.Val}...)
    }
}

/* -- Iterative Solutions -- 
* Use a stack to store nodes 
* Pop current node out of stack and push its child nodes in 
* Order: Top -> Bottom and Left -> Right 
*/
func preorderTraversal(root *TreeNode) []int {
    
    if (root == nil) {
        return []int{}
    } 
    
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    
    res := make([]int, 0)
    
    for len(stack) > 0 {
        curr := stack[(len(stack) - 1)]
        stack = stack[0:len(stack)-1]
        
        if (curr != nil) {
            res = append(res, curr.Val)
            
            if (curr.Right != nil) {
                stack = append(stack, curr.Right)
            }
            
            if (curr.Left != nil) {
                stack = append(stack, curr.Left)
            }
        }
    }
    
    return res
    
}

/**
* Use a stack to contain visiting nodes 
* Starting from root, 
*   - Update stack with its left subTrees 
*   - As poping element from the stack
*        - Update its value to result list 
*        - Visiting its right subtree, update it to stack 
*/
func inorderTraversal(root *TreeNode) []int {
    if (root == nil) {
        return []int{}
    } 
    
    stack := make([]*TreeNode, 0)
    fmt.Println("Initial stack ", stack)
    curr := root 
    res := make([]int, 0)
    
    for (curr != nil || len(stack) > 0) {        
        // Keep adding left subtrees to stack 
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        
        // Poping element 
        curr = stack[len(stack) - 1]
        stack = stack[0:len(stack)-1]

        // Update element to result list
        res = append(res, curr.Val)
        
        // Visiting right subtree
        curr = curr.Right
        
    }
    return res
}
