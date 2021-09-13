/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/** -- Recursive -- 
* Compare current value with a range and update the range while traversing through the tree 
* At root, range = -inf -> +inf 
* When validate left subtree of the root, range need to be from -inf -> root.Val
* When validate right subtree of the root, range need to be from root.Val -> +inf 
* 
*/
const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
    MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
    MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
    MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

func isValidBST(root *TreeNode) bool {
    return validate(root, MinInt, MaxInt)
}

func validate(root *TreeNode, low int, high int) bool {
    if (root == nil) {
        return true
    }
    if ((root != nil && root.Val <= low) || 
        (root != nil && root.Val >= high)) {
            return false
    }
    return validate(root.Left, low, root.Val) && validate(root.Right, root.Val, high)
}
