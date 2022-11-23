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
			tC.a = tC.tree.left.left   //420
			tC.b = tC.tree.right.right //999
			tC.want = tC.tree          //100
		}, "a", nil, nil, nil, nil},
		{func(tC *testCase) {
			tC.tree = Tree()
			tC.a = tC.tree.left.right //420
			tC.b = tC.tree.right.left //69
			tC.want = tC.tree         //365
		}, "b", nil, nil, nil, nil},
		{func(tC *testCase) {
			tC.tree = Tree()
			tC.a = tC.tree.left.left  //78
			tC.b = tC.tree.left.right //69
			tC.want = tC.tree.left    //100
		}, "c", nil, nil, nil, nil},
		{func(tC *testCase) {
			tC.tree = nil
			tC.a = nil
			tC.b = nil
			tC.want = nil
		}, "Nil", nil, nil, nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.init(&tc)
			if nca(tc.tree, tc.a, tc.b) != tc.want {
				t.Errorf("got = %v, want = %v", nca(tc.tree, tc.a, tc.b), tc.want)
			}
		})
	}
}
