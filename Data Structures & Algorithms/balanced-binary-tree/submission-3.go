/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    res := dfs(root)
	if res == -1 {
		return false
	}
	return true
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left)
	right := dfs(node.Right)

	if left == -1 || right == -1 || math.Abs(float64(left - right)) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

