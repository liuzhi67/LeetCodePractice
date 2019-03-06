/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, target := findLowestCommonAncestor(root, p, q)
	return target
}

func findLowestCommonAncestor(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	hitl, tl := findLowestCommonAncestor(root.Left, p, q)
	hitr, tr := findLowestCommonAncestor(root.Right, p, q)
	if tl != nil {
		return 2, tl
	}
	if tr != nil {
		return 2, tr
	}
	hitc := 0
	if root.Val == q.Val || root.Val == p.Val {
		hitc = 1
	}
	hit := hitl + hitr + hitc
	if hit == 2 {
		return 2, root
	}
	return hit, nil
}
