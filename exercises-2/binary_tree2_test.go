package exercises2

import (
	"testing"
)

func fullTree(depth int) *node {
	if depth <= 0 {
		return nil
	}
	return &node{left: fullTree(depth - 1), right: fullTree(depth - 1)}
}

func tree(depth int) func() *node {
	return func() *node { return fullTree(depth) }
}

func root(t *node) *node { return t }

func child(dirs ...func(*node) *node) func(*node) *node {
	return func(t *node) *node {
		for _, d := range dirs {
			t = d(t)
		}
		return t
	}
}

func left(t *node) *node  { return t.left }
func right(t *node) *node { return t.right }

func TestNearestAncestor(t *testing.T) {
	for _, tc := range []struct {
		name string
		t    func() *node
		a, b func(t *node) *node
		want func(t *node) *node
	}{
		{name: "empty", t: tree(0),
			a:    root,
			b:    root,
			want: root},
		{name: "single", t: tree(1),
			a:    root,
			b:    root,
			want: root},
		{name: "direct parent 1", t: tree(2),
			a:    child(left),
			b:    child(right),
			want: root},
		{name: "direct parent 2", t: tree(3),
			a:    child(left, left),
			b:    child(left, right),
			want: child(left)},
		{name: "direct parent 3", t: tree(3),
			a:    child(right, right),
			b:    child(right, left),
			want: child(right)},

		// TODO: ... MORE MORE MORE TESTS !!!

	} {
		t.Run(tc.name, func(t *testing.T) {
			root := tc.t()
			if got := NearestAncestor(root, tc.a(root), tc.b(root)); got != tc.want(root) {
				printTree(root, t) //meeh, t.log was unfortunate idea, but fmt.Print does not work here...
				t.Error("returned unexpected node")
			}
		})
	}
}

func printTree(root *node, t *testing.T) {
	printer(root, nil, t)
}

func printNonemptyHistory(prev []int, printLast bool, t *testing.T) {
	if len(prev) == 0 {
		panic("must not happen")
	}
	last := len(prev)
	if !printLast {
		last--
	}
	for _, p := range prev[:last] {
		if p > 0 {
			t.Log("| ")
		} else {
			t.Log("  ")
		}
	}
}

func directChildren(n *node, t *testing.T) (num int) {
	if n == nil {
		return 0
	}
	if n.left != nil {
		num++
	}
	if n.right != nil {
		num++
	}
	return num
}

func printer(n *node, prev []int, t *testing.T) {
	if n == nil {
		return
	}
	if len(prev) > 0 {
		printNonemptyHistory(prev, true, t)
		t.Log()
		printNonemptyHistory(prev, false, t)
		prev[len(prev)-1]--
	}
	t.Log("+->", n.val)
	prev = append(prev, directChildren(n, t))
	if n.left != nil {
		printer(n.left, prev, t)
	}
	if n.right != nil {
		printer(n.right, prev, t)
	}
}
