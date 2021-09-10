/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/* --- Recursion -- 
* Traversing from root node to left and right node. As doing that, recording max depth  
* H(root) = 1 + max(H(root-leftsubtree), H(root-rightSubtree))
* O(n) for time complexity to traverse through all nodes in the tree 
* O(n) for space need for call stack when tree is completely unbalanced 
*/


func max(x int, y int) int {
    if (x > y) {
        return x
    }  else {
        return y
    }
}

/*
func maxDepth(root *TreeNode) int {
    if (root == nil) {
        return 0
    } else {
        return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
    }
}
*/

/** -- Use Stack -- 
* Keep track of visting nodes in a stack 
* As stack is not empty, keep visiting nodes and put it into the stack with its depth 
* Keep track of max depth 
* O(n) for time and space complexity 
*/

type TreeNodeWithDepth struct {
    node *TreeNode
    depth int
}

func maxDepth(root *TreeNode) int {
        
    stack := make([]TreeNodeWithDepth, 0)
    
    if (root != nil) {
        stack = append(stack, TreeNodeWithDepth{root, 1})
    }

    res := 0 
    for len(stack) > 0 {
        elem := stack[0]
        stack = stack[1:]
        
        if (elem.node != nil) {
            res = max(elem.depth, res)
            stack = append(stack, TreeNodeWithDepth{elem.node.Left, elem.depth + 1})
            stack = append(stack, TreeNodeWithDepth{elem.node.Right, elem.depth + 1})
        }
    }
    
    return res
    
}
