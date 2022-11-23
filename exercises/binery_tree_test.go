package main

import (
	"reflect"
	"testing"
)

func small() *node {
	return &node{val: 1}
}

func small2() *node {
	return &node{val: 1, right: &node{val: 2}}
}

func small3() *node {
	return &node{val: 1, left: &node{val: 2, right: &node{val: 3}}}
}
func large() *node {
	return &node{val: 2,
		left: &node{val: 1},
		right: &node{val: 4,
			left:  &node{val: 3},
			right: &node{val: 5},
		}}
}

func large2() *node {
	return &node{val: 2,
		left: &node{val: 1},
		right: &node{val: 4,
			left: &node{val: 3},
			right: &node{val: 5,
				left: &node{val: 9},
			},
		}}
}

func TestNodeCount(t *testing.T) {
	for _, tc := range []struct {
		name string
		tree *node
		want int
	}{
		{"large bin tree", large(), 5},
		{"large 2 bin tree", large2(), 6},
		{"small bin tree", small(), 1},
		{"small 2 bin tree", small2(), 2},
		{"small 3 bin tree", small3(), 3},
		{"NIL tree", nil, 0},
	} {
		if got := nodeCount(tc.tree); got != tc.want {
			t.Errorf("%v; got = %v, want = %v", tc.name, got, tc.want)
		}
	}
}

func TestHeight(t *testing.T) {
	for _, tc := range []struct {
		name string
		tree *node
		want int
	}{
		{"large bin tree", large(), 2},
		{"large 2 bin tree", large2(), 3},
		{"small bin tree", small(), 0},
		{"small 2 bin tree", small2(), 1},
		{"small 3 bin tree", small3(), 2},
		{"NIL tree", nil, 0},
	} {
		if got := height(tc.tree); got != tc.want {
			t.Errorf("%v; got = %v, want = %v", tc.name, got, tc.want)
		}
	}
}

func TestEqual(t *testing.T) {
	for _, tc := range []struct {
		name  string
		tree1 *node
		tree2 *node
		want  bool
	}{
		{"almost equal", small2(), small3(), false},
		{"equal", small2(), small2(), true},
		{"almost equal2", large(), large2(), false},
		{"equal2", large(), large(), true},
		{"equal3", large2(), large2(), true},
	} {
		if got := equal(tc.tree1, tc.tree2); got != tc.want {
			t.Errorf("%v; got = %v, want = %v", tc.name, got, tc.want)
		}
	}
}

func TestInorderTraversal(t *testing.T) {
	for _, tc := range []struct {
		tree *node
		want []int
	}{
		{large(), []int{1, 2, 3, 4, 5}},
		{large2(), []int{1, 2, 3, 4, 9, 5}},
		{small(), []int{1}},
		{small2(), []int{1, 2}},
	} {
		var got []int
		inorderTraversal(tc.tree, func(v int) { got = append(got, v) })
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}

func TestInvertTree(t *testing.T) {
	for _, tc := range []struct {
		tree *node
		want []int
	}{
		{large(), []int{5, 4, 3, 2, 1}},
		{large2(), []int{5, 9, 4, 3, 2, 1}},
		{small(), []int{1}},
		{small2(), []int{2, 1}},
	} {
		var got []int
		invertTree(tc.tree)
		inorderTraversal(tc.tree, func(v int) { got = append(got, v) })
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
