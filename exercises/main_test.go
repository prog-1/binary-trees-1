package main

import (
	"reflect"
	"testing"
)

func small() *node {
	return &node{val: 1, left: &node{val: 2}, right: &node{val: 3}}
}

func large() *node {
	return &node{val: 1, left: &node{val: 2, right: &node{val: 3}}, right: &node{val: 4, left: &node{val: 5}}}
}

func incomplete() *node {
	return &node{val: 1, left: &node{val: 2}}
}

func TestNodeCount(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input *node
		want  int
	}{
		{"nil", nil, 0},
		{"small", small(), 3},
		{"large", large(), 5},
		{"incomplete", incomplete(), 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := nodeCount(tc.input); got != tc.want {
				t.Errorf("Test: %v, got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestHeight(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input *node
		want  int
	}{
		{"nil", nil, 0},
		{"small", small(), 1},
		{"large", large(), 2},
		{"incomplete", incomplete(), 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := height(tc.input); got != tc.want {
				t.Errorf("Test: %v, got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input *node
		want  []int
	}{
		{"nil", nil, nil},
		{"small", small(), []int{2, 1, 3}},
		{"large", large(), []int{2, 3, 1, 5, 4}},
		{"incomplete", incomplete(), []int{2, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			inorderTraversal(tc.input, func(v int) { got = append(got, v) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Test: %v, got = %v, want = %v", tc.name, got, tc.want)
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
		{"nil", nil, nil, true},
		{"one nil", nil, small(), false},
		{"not the same", large(), small(), false},
		{"not the same2", large(), incomplete(), false},
		{"not the same3", small(), incomplete(), false},
		{"same size dif. values", large(), &node{val: 1, left: &node{val: 2, right: &node{val: 3}}, right: &node{val: 4, left: &node{val: 6}}}, false},
		{"same size dif. values2", large(), &node{val: 8, left: &node{val: 12, right: &node{val: 23}}, right: &node{val: 4, left: &node{val: 5}}}, false},
		{"same", large(), large(), true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if equal(tc.a, tc.b) != tc.want {
				t.Errorf("Test: %v, got = %v, want = %v", tc.name, equal(tc.a, tc.b), tc.want)
			}
		})
	}
}

func TestInvertTree(t *testing.T) {
	for _, tc := range []struct {
		name        string
		input, want *node
	}{
		{"nil", nil, nil},
		{"small", small(), &node{val: 1, left: &node{val: 3}, right: &node{val: 2}}},
		{"large", large(), &node{val: 1, left: &node{val: 4, left: &node{val: 5}}, right: &node{val: 2, right: &node{val: 3}}}},
		{"incomplete", incomplete(), &node{val: 1, right: &node{val: 2}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			invertTree(tc.input)
			if !equal(tc.input, tc.want) {
				t.Errorf("Test name: %v, equal: %v", tc.name, equal(tc.input, tc.want))
			}
		})
	}
}
