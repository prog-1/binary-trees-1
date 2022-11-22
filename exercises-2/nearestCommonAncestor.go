package main

import "fmt"

func NearestCommonAncestor(root, a, b *node[int]) *node[int] {
	var aS, bS []*node[int]

	//find all sequence from element to root for a
	//find all sequence from element to root for b
	if !sequence(root, a, &aS) || !sequence(root, b, &bS) {
		return nil
	}

	//find element, where sequences are going sideways
	var i int
	for ; i < len(aS) && aS[i] == bS[i]; i++ {
	}
	return aS[i-1]

}

func sequence(root *node[int], target *node[int], seq *[]*node[int]) bool {
	if root == nil {
		return false
	}

	*seq = append(*seq, root)

	if root == target || root.left != nil && sequence(root.left, target, seq) || root.right != nil && sequence(root.right, target, seq) {
		return true
	}

	(*seq) = (*seq)[:len(*seq)-1]
	return false

}

type node[T any] struct {
	val         T
	left, right *node[T]
}

type nodeInt = node[int]

func n[T any](v T) *node[T]              { return &node[T]{val: v} }
func (n *node[T]) l(l *node[T]) *node[T] { n.left = l; return n }
func (n *node[T]) r(r *node[T]) *node[T] { n.right = r; return n }

func Tree() *nodeInt {
	return n(1).l(n(2).l(n(4)).r(n(5))).r(n(3).l(n(6)).r(n(7)))
}

func main() {
	tree := Tree()
	a := tree.right.left
	b := tree.right.left
	fmt.Println(NearestCommonAncestor(tree, a, b))
}
