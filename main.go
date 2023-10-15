package main

import (
	"fmt"
	"leetcode/tasks"
)

func main() {
	//First task
	tree1 := tasks.TreeNode{
		Val:  1,
		Left: nil,
		Right: &tasks.TreeNode{
			Val:  2,
			Left: nil,
			Right: &tasks.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	tree2 := tasks.TreeNode{
		Val:  1,
		Left: nil,
		Right: &tasks.TreeNode{
			Val:  2,
			Left: nil,
			Right: &tasks.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(tasks.IsSameTree(&tree1, &tree2))

	//Second task
	fmt.Println(tasks.IsSymmetric(&tree1))

	//Third task
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	tasks.Merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	//Fourth task
	fmt.Println(tasks.Reverse(4152))

	//Fifth task
	fmt.Println(tasks.MyAtoi("4193 with words"))
}
