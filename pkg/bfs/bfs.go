package bfs

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Bfs(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		l := queue.Len()
		arr := make([]int, 0)
		for i := 0; i < l; i++ {
			t := queue.Front().Value.(*TreeNode)
			if t.Left != nil {
				queue.PushBack(t.Left)
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
			}
			arr = append(arr, t.Val)
			queue.Remove(queue.Front())
		}
		res = append(res, arr)
	}
	return res
}
