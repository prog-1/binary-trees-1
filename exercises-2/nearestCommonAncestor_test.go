package main

import "testing"

type testCase struct {
	init func(tC *testCase)
	name string
	tree *node[int]
	a    *node[int]
	b    *node[int]
	want *node[int]
}

func TestNearestCommonAncestor(t *testing.T) {
	for _, tc := range []testCase{
		{func(tC *testCase) {
			tC.tree = Tree()
			tC.a = tC.tree.left.left  //4
			tC.b = tC.tree.left.right //5
			tC.want = tC.tree.left    //2
		}, "1", nil, nil, nil, nil},
		{func(tC *testCase) {
			tC.tree = Tree()
			tC.a = tC.tree.left.left   //4
			tC.b = tC.tree.right.right //7
			tC.want = tC.tree          //1
		}, "2", nil, nil, nil, nil},
		{func(tC *testCase) {
			tC.tree = Tree()
			tC.a = tC.tree.right.left    //6
			tC.b = tC.tree.right.left    //6
			tC.want = tC.tree.right.left //6
		}, "Same element", nil, nil, nil, nil},
		{func(tC *testCase) {
			tC.tree = nil
			tC.a = nil
			tC.b = nil
			tC.want = nil
		}, "Nil", nil, nil, nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.init(&tc)
			if NearestCommonAncestor(tc.tree, tc.a, tc.b) != tc.want {
				t.Errorf("got = %v, want = %v", NearestCommonAncestor(tc.tree, tc.a, tc.b), tc.want)
			}
		})
	}
}
