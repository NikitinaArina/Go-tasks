package tasks

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
https://leetcode.com/problems/same-tree/

	Задача 100 Same Tree

Проверяем на нил, если оба нил, то значит деревья одинаковые - возвращаем тру
Проверяем равен ли какой-нибудь из ноды нил, если да, то деревья уже не будут одинаковые - возвращаем фолз
Если значения нод не равны, то сразу фолз
Рекурсивно вызываем метод, чтобы проверить что левая нода 1-ого дерева == левой ноде 2-ого дерева и правая == правой
*/
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil || q == nil && p != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
