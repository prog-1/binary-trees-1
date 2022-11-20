package main

import (
	"reflect"
	"testing"
)

func TestNodeCount(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want int
	}{
		{"case-1", nil, 0},
		{"case-2", &node{val: 1}, 1},
		{"case-3", &node{val: 1, left: &node{val: 2}}, 2},
		{"case-4", &node{val: 1, left: &node{val: 2}, right: &node{val: 3}}, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := nodeCount(tc.root); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestHeight(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want int
	}{
		{"case-1", nil, 0},
		{"case-2", &node{val: 1}, 0},
		{"case-3", &node{val: 1, left: &node{val: 2}}, 1},
		{"case-4", &node{val: 1, left: &node{val: 2}, right: &node{val: 3}}, 1},
		{"case-5", &node{val: 1, left: &node{val: 2, right: &node{val: 4}}, right: &node{val: 3}}, 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := height(tc.root); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", &node{val: 1}, []int{1}},
		{"case-3", &node{val: 1, left: &node{val: 2}}, []int{2, 1}},
		{"case-4", &node{val: 1, left: &node{val: 2}, right: &node{val: 3}}, []int{2, 1, 3}},
		{"case-5", &node{val: 1, left: &node{val: 2, left: &node{val: 4}}, right: &node{val: 3}}, []int{4, 2, 1, 3}},
		{"case-6", &node{val: 1, left: &node{val: 2, right: &node{val: 4}}, right: &node{val: 3}}, []int{2, 4, 1, 3}},
		{"case-7", &node{val: 1, left: &node{val: 2, left: &node{val: 4}, right: &node{val: 5}}, right: &node{val: 3}}, []int{4, 2, 5, 1, 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			inorderTraversal(tc.root, func(v int) { got = append(got, v) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	for _, tc := range []struct {
		name string
		a, b *node
		want bool
	}{
		{"case-1", nil, nil, true},
		{"case-2", nil, &node{val: 1}, false},
		{"case-3", &node{val: 1}, &node{val: 2}, false},
		{"case-4", &node{val: 1}, &node{val: 1, left: &node{val: 1}}, false},
		{"case-5", &node{val: 1}, &node{val: 1}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if equal(tc.a, tc.b) != tc.want {
				t.Errorf("got = %v, want = %v", equal(tc.a, tc.b), tc.want)
			}
		})
	}
}

func TestInvertTree(t *testing.T) {
	for _, tc := range []struct {
		name       string
		root, want *node
	}{
		{"case-1", nil, nil},
		{"case-2", &node{val: 1}, &node{val: 1}},
		{"case-3", &node{val: 1, left: &node{val: 2}}, &node{val: 1, right: &node{val: 2}}},
		{"case-4", &node{val: 1, left: &node{val: 2}, right: &node{val: 3}}, &node{val: 1, right: &node{val: 2}, left: &node{val: 3}}},
		{"case-5", &node{val: 1, left: &node{val: 2, right: &node{val: 4}}, right: &node{val: 3}}, &node{val: 1, right: &node{val: 2, left: &node{val: 4}}, left: &node{val: 3}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			invertTree(tc.root)
			if !equal(tc.root, tc.want) {
				t.Errorf("equal: %v", equal(tc.root, tc.want))
			}
		})
	}
}
