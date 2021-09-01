/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
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
