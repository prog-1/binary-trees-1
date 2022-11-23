package exercises2

type node struct {
	val         int
	left, right *node
}

func EasierNearestAncestor(root, a, b *node) *node {
	if root == nil || root == a || root == b {
		return root
	}
	left := EasierNearestAncestor(root.left, a, b)   // this is where i stuck on Divide and Conquer hw, it was always infinite recursion
	right := EasierNearestAncestor(root.right, a, b) // when i assigned recursion to variable
	if left != nil && right != nil {                 // i just placed first if wrong
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
	//point from a,b to nearest ancestor
}

func Walk(root, EndPoint *node) []*node {
	//remember all nodes from root to a and b
	// in some way similar with easier version
	if root == nil || root == EndPoint { // this if really saves from infinite recursion
		return []*node{root}
	}
	left := Walk(root.left, EndPoint)
	right := Walk(root.right, EndPoint)
	if len(left) != 0 && len(right) != 0 { // maybe i need one more if?
		return nil
	}
	if len(left) != 0 {
		return append([]*node{root}, left...)
	}
	return append([]*node{root}, right...)
}

func NearestAncestor(root, a, b *node) *node {
	WalkToA := Walk(root, a)
	WalkToB := Walk(root, b)
	if len(WalkToA) > len(WalkToB) { // ensure B is longer
		WalkToA, WalkToB = WalkToB, WalkToA
	}
	i := 0
	for ; i != len(WalkToA); i++ { // searching... (in shortest array)
		if WalkToA[i] != WalkToB[i] { // error is definetly down  here
			return WalkToA[i-1]
		}
	}
	if i < len(WalkToA) {
		return WalkToB[len(WalkToB)-1]
	} else /*if i < len(WalkToB) */ {
		return WalkToA[len(WalkToA)-1]
	}
}

// it worked at first but then i did rewrote something wrong... aaaa
