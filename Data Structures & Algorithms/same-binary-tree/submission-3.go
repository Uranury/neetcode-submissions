/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if dfs(p, q) == -1 {
		return false
	}
	return true
}

func dfs(p *TreeNode, q *TreeNode) int {
	if p == nil && q == nil {
		return 0
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return -1
	}
	if p.Val != q.Val {
		return -1
	}

	left, right := dfs(p.Left, q.Left), dfs(p.Right, q.Right)

	if left == -1 || right == -1 {
		return -1
	}
	return min(left, right)
}