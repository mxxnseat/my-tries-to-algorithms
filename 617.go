package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNodeOrNull(node *TreeNode, side string) *TreeNode {
	if node == nil {
		return nil
	}
	if side == "left" {
		return node.Left
	}
	return node.Right
}

func bfs(ans *TreeNode, root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	newValue := 0
	if root1 != nil {
		newValue += root1.Val
		if root1.Left != nil {
			ans.Left = &TreeNode{}
		}
		if root1.Right != nil {
			ans.Right = &TreeNode{}
		}
	}
	if root2 != nil {
		newValue += root2.Val
		if root2.Left != nil && ans.Left == nil {
			ans.Left = &TreeNode{}
		}
		if root2.Right != nil && ans.Right == nil {
			ans.Right = &TreeNode{}
		}
	}
	ans.Val = newValue
	bfs(ans.Left, getNodeOrNull(root1, "left"), getNodeOrNull(root2, "left"))
	bfs(ans.Right, getNodeOrNull(root1, "right"), getNodeOrNull(root2, "right"))
	return ans
}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return bfs(&TreeNode{}, root1, root2)
}
