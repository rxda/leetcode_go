package src

//  11
// 7  2     sum=18
//h(11,18) -> h(7,7)->true
//         -> h(2,7)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
		return true
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
