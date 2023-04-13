package tasks

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	s := make([]*TreeNode, 0, 15)
	s = append(s, root)

	var res int

	for len(s) > 0 {
		var i *TreeNode
		i, s = s[len(s)-1], s[:len(s)-1]

		if i.Left == nil && i.Right == nil {
			count++
			if count == k {
				res = i.Val
				break
			}
			continue
		}

		if i.Right != nil {
			s = append(s, i.Right)
		}

		s = append(s, &TreeNode{Val: i.Val})

		if i.Left != nil {
			s = append(s, i.Left)
		}
	}

	return res
}
