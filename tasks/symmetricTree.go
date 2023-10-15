package tasks

/*
В целом за основу взят метод из прошлой задачи, поменяла в нем ретерн только
 1. Если нил, то дерево симметрично
 2. Вызов IsSameTree2:
    Дерево симметрично, если его левый элемент равен правому
    Внутри IsSameTree2 рекурсивно вызываем метод, чтобы проверить что левый равен правому, а правый равен левому
    https://leetcode.com/problems/symmetric-tree/description/
*/
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if isSameTree2(root.Left, root.Right) {
		return true
	} else {
		return false
	}
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil || q == nil && p != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree2(p.Left, q.Right) && isSameTree2(p.Right, q.Left)
}
