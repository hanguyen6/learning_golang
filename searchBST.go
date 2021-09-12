* -- Recursive --- 
* Search in BST: if current value < target, visit right substree 
* Otherwise visit left subtree 
* Until currentNode = target return currentNode 
* Or nill -> return nul
* O(log(n)) for time and complexity 
*/
func searchBST(root *TreeNode, val int) *TreeNode {
    if (root == nil) {
        return nil
    } else if (root.Val == val) {
        return root
    } else if (root.Val > val) {
        return searchBST(root.Left, val)
    } else {
        return searchBST(root.Right, val)
    } 
}

* -- Iterative solution -- 
* O(log(n)) time complexity 
* O(1) space need 
func searchBST(root *TreeNode, val int) *TreeNode {
    for (root != nil && root.Val != val) {
        if (root.Val > val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return root
}
