/**
* Start traversing the tree from the root node 
* If both node p and q are in the left subtree, continue the search with left subTree 
* If both node p, q are in the right subtree, continue the search with right subTree
* Otherwise, the current node is common to node p and q subtree
* O(N) for time and space complexity 
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if (root.Val > p.Val && root.Val > q.Val) {
        return lowestCommonAncestor(root.Left, p, q)
    } else if (root.Val < p.Val && root.Val < q.Val) {
        return lowestCommonAncestor(root.Right, p, q)
    } else {
        return root
    }
}
