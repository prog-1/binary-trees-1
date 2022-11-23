package main

import "fmt"

type node struct {
	v     int
	left  *node
	right *node
}

func findAncestor(root, p1, p2 *node) *node {
	var path1, path2 []*node
	findPath(root, p1, &path1)
	findPath(root, p2, &path2)
	var i int
	for i < len(path1) && i < len(path2) && path1[i] == path2[i] {
		i++
	}
	return path1[i-1]
}

func findPath(root, p *node, path *[]*node) bool {
	*path = append(*path, root)
	if root == p || (root.left != nil && findPath(root.left, p, path)) || (root.right != nil && findPath(root.right, p, path)) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func main() {
	a := &node{v: 5, left: &node{v: 9}, right: &node{v: 10}}
	b := &node{v: 8}
	ancesr := findAncestor(&node{v: 1, left: &node{v: 2, left: &node{v: 4, right: b}, right: a}, right: &node{v: 3}}, a, b)
	fmt.Println(format(ancesr))
}

func format(l *node) string {
	var b string
	for e := l; e != nil; {
		b = b + fmt.Sprint(e.v)
		if e.left != nil && e.right != nil {
			b = b + " -> "
		}
		if e.left != nil && e.right == nil {
			e = e.left
		} else {
			e = e.right
		}
	}
	return b
}
