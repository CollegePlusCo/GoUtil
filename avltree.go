package util

type Comparable interface {
	LessThan(comparable2 Comparable) bool
	MoreThan(comparable2 Comparable) bool
	EqualTo(comparable2 Comparable) bool
	GetInteger() int
}

type TreeNode[T Comparable] struct {
	data   T
	left   *TreeNode[T]
	right  *TreeNode[T]
	height int
}

type AVLTree[T Comparable] struct {
	Root *TreeNode[T]
	Size int
}

// height of the tree
func (tree *AVLTree[T]) height(currentNode *TreeNode[T]) int {
	if currentNode == nil {
		return 0
	}
	return currentNode.height
}

// A utility function to get maximum
// of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (node *TreeNode[T]) GreaterThan(data2 any) bool {
	data := node.data
	return data.MoreThan(data2.(Comparable))
}

func (node *TreeNode[T]) LesserThan(data2 any) bool {
	data := node.data
	return data.LessThan(data2.(Comparable))
}

func (tree *AVLTree[T]) Insert(curr **TreeNode[T], x T) {
	if (*curr) == nil {
		tree.Size++
		*curr = &TreeNode[T]{data: x}
	} else if (*curr).GreaterThan(x) {
		tree.Insert(&(*curr).left, x)
		//curr.left = &TreeNode[T]{data: x}
		if tree.height((*curr).left)-tree.height((*curr).right) > 1 {
			if (*curr).left.GreaterThan(x) {
				tree.case1(curr)
			} else {
				tree.case2(curr)
			}
		}

	} else if (*curr).LesserThan(x) {
		tree.Insert(&(*curr).right, x)
		//if curr.right == nil {
		//	curr.right = &TreeNode[T]{data: x}
		if tree.height((*curr).right)-tree.height((*curr).left) > 1 {
			if (*curr).right.LesserThan(x) {
				tree.case4(curr)
			} else {
				tree.case3(curr)
			}
		}
	} else {
		tree.Insert(&(*curr).left, x)
		//curr.left = &TreeNode[T]{data: x}
		if tree.height((*curr).left)-tree.height((*curr).right) > 1 {
			if (*curr).left.GreaterThan(x) {
				tree.case1(curr)
			} else if (*curr).left.LesserThan(x) {
				tree.case2(curr)
			} else {
				tree.case1(curr)
			}
		}
	}
	(*curr).height = max(tree.height((*curr).left), tree.height((*curr).right)) + 1
}

// prints the values in the AVLTree in order
func (tree *AVLTree[T]) Inorder(node *TreeNode[T]) {
	if node == nil {
		return
	}
	tree.Inorder(node.left)
	tree.Inorder(node.right)
}

// prints the values in the AVLTree in order
func (tree *AVLTree[T]) GetInorder(node *TreeNode[T], ordered *[]any) {
	if node == nil {
		return
	}
	tree.GetInorder(node.left, ordered)
	*ordered = append(*ordered, node.data)
	tree.GetInorder(node.right, ordered)
}

// prints the values in the AVLTree in order
func (tree *AVLTree[T]) GetReversed(node *TreeNode[T], ordered *[]T) {
	if node == nil {
		return
	}
	tree.GetReversed(node.right, ordered)
	*ordered = append(*ordered, node.data)
	tree.GetReversed(node.left, ordered)
}

func (tree *AVLTree[T]) Exists(curr *TreeNode[T], x int) bool {
	if curr == nil {
		return false
	}
	if curr.GreaterThan(x) {
		return tree.Exists(curr.left, x)
	} else if curr.LesserThan(x) {
		return tree.Exists(curr.right, x)
	}
	return true
}

// left left rotation
func (tree *AVLTree[T]) case1(k2 **TreeNode[T]) {
	k1 := (*k2).left
	if k1 != nil {
		(*k2).left = k1.right
		k1.right = *k2
	}
	(*k2).height = max(tree.height((*k2).left), tree.height((*k2).right)) + 1
	if k1 != nil {
		k1.height = max(tree.height(k1.left), (*k2).height) + 1
		*k2 = k1
	}
}

// right right rotation
func (tree *AVLTree[T]) case4(k1 **TreeNode[T]) {
	k2 := (*k1).right
	if k2 != nil {
		(*k1).right = k2.left
		k2.left = *k1
	}
	(*k1).height = max(tree.height((*k1).left), tree.height((*k1).right)) + 1
	if k2 != nil {
		k2.height = max(tree.height(k2.left), (*k1).height) + 1
		*k1 = k2
	}
}

// left right rotation
func (tree *AVLTree[T]) case2(k3 **TreeNode[T]) {
	tree.case4(&(*k3).left)
	tree.case1(k3)
}

// right left rotation
func (tree *AVLTree[T]) case3(k1 **TreeNode[T]) {
	tree.case1(&(*k1).right)
	tree.case4(k1)
}
