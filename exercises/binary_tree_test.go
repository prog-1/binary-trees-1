package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	name          string
	input         func() *node // Func ensures each test wil get its own tree.
	wantNodeCount int
	wantHeight    int
	wantInorder   []int // Again, to ensure that we get fresh values.
	wantCompRes   []bool
	wantInverted  func() *node
}

var testCases = []testCase{
	{"nil", func() *node { return nil }, 0, 0, nil, []bool{true, false, false}, func() *node { return nil }},
	{"complete tree", func() *node {
		return &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, &node{6, &node{5, nil, nil}, &node{7, nil, nil}}}
	}, 7, 2, []int{1, 2, 3, 4, 5, 6, 7}, []bool{false, true, false}, func() *node {
		return &node{4, &node{6, &node{7, nil, nil}, &node{5, nil, nil}}, &node{2, &node{3, nil, nil}, &node{1, nil, nil}}}
	}},
	{"incomplete tree", func() *node {
		return &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, nil}
	}, 4, 2, []int{1, 2, 3, 4}, []bool{false, false, true}, func() *node {
		return &node{4, nil, &node{2, &node{3, nil, nil}, &node{1, nil, nil}}}
	}},
}

func TestNodeCount(t *testing.T) {
	for _, tc := range testCases {
		if got := nodeCount(tc.input()); got != tc.wantNodeCount {
			t.Errorf("got = %v, want = %v", got, tc.wantNodeCount)
		}
	}
}

func TestHeight(t *testing.T) {
	for _, tc := range testCases {
		if got := height(tc.input()); got != tc.wantHeight {
			t.Errorf("got = %v, want = %v", got, tc.wantHeight)
		}
	}
}

func TestInorderTraversal(t *testing.T) {
	for _, tc := range testCases {
		var got []int
		inorderTraversal(tc.input(), func(v int) { got = append(got, v) })
		if !reflect.DeepEqual(got, tc.wantInorder) {
			t.Errorf("got = %v, want = %v", got, tc.wantInorder)
		}
	}
}

func TestEqual(t *testing.T) {
	for _, tcA := range testCases {
		for i, tcB := range testCases { // Comparing each test to itself and all others
			if equal(tcA.input(), tcB.input()) != tcA.wantCompRes[i] {
				t.Errorf("got = %v, want = %v", !tcA.wantCompRes[i], tcA.wantCompRes[i])
			}
		}
	}
}

func TestInvertTree(t *testing.T) {
	treeToSlice := func(root *node) []int {
		var s []int
		inorderTraversal(root, func(v int) { s = append(s, v) })
		return s
	}
	for _, tc := range testCases {
		got := tc.input()
		invertTree(got)
		if !equal(got, tc.wantInverted()) {
			t.Errorf("got = %v, want = %v", treeToSlice(got), treeToSlice(tc.wantInverted()))
		}
	}
}
