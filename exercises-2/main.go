package main

type node struct {
	val         int
	left, right *node
}

func nearestCommonAncestor(t *node, a, b *node) *node {
	p1 := FingAndGetParents(t, a)
	// fmt.Println(p1)
	p2 := FingAndGetParents(t, b)
	// fmt.Println(p2)
	if len(p1) > len(p2) {
		p1, p2 = p2, p1
	}
	for i := range p1 {
		if p1[i] != p2[i] {
			return p1[i-1]
		}
	}
	return p1[len(p1)-1]
}

func FingAndGetParents(t, a *node) (parents []*node) {
	if t == nil {
		return nil
	}
	if t == a {
		return []*node{t}
	}
	l, r := FingAndGetParents(t.left, a), FingAndGetParents(t.right, a)
	if len(r) != 0 {
		return append([]*node{t}, r...)
	}
	if len(l) != 0 {
		return append([]*node{t}, l...)
	}
	return nil
}

func main() {

}
