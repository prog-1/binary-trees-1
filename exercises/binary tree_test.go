package main

import (
	"reflect"
	"testing"
)

func TestNodeCount(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    *node
		want int
	}{
		{"test-1", nil, 0},
		{"test-2", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, &node{6, &node{5, nil, nil}, &node{7, nil, nil}}}, 7},
		{"test-3", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, nil}, 4},
		{"test-4", &node{3, &node{5, &node{6, nil, nil}, &node{2, &node{7, nil, nil}, &node{4, nil, nil}}}, &node{1, &node{0, nil, nil}, &node{8, nil, nil}}}, 9},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := nodeCount(tc.n)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestHeight(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    *node
		want int
	}{
		{"test-1", nil, 0},
		{"test-2", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, &node{6, &node{5, nil, nil}, &node{7, nil, nil}}}, 2},
		{"test-3", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, nil}, 2},
		{"test-4", &node{3, &node{5, &node{6, nil, nil}, &node{2, &node{7, nil, nil}, &node{4, nil, nil}}}, &node{1, &node{0, nil, nil}, &node{8, nil, nil}}}, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := height(tc.n)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    *node
		want []int
	}{
		{"test-1", nil, nil},
		{"test-2", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, &node{6, &node{5, nil, nil}, &node{7, nil, nil}}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"test-3", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, nil}, []int{1, 2, 3, 4}},
		{"test-4", &node{3, &node{5, &node{6, nil, nil}, &node{2, &node{7, nil, nil}, &node{4, nil, nil}}}, &node{1, &node{0, nil, nil}, &node{8, nil, nil}}}, []int{6, 5, 7, 2, 4, 3, 0, 1, 8}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			inorderTraversal(tc.n, func(v int) { got = append(got, v) })
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
		{"test-1", nil, nil, true},
		{"test-2", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, &node{6, &node{5, nil, nil}, &node{7, nil, nil}}},
			&node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, &node{6, &node{5, nil, nil}, &node{7, nil, nil}}}, true},
		{"test-3", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, nil},
			&node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, &node{6, &node{5, nil, nil}, &node{7, nil, nil}}}, false},
		{"test-4", &node{3, &node{5, &node{6, nil, nil}, &node{2, &node{7, nil, nil}, &node{4, nil, nil}}}, &node{1, &node{0, nil, nil}, &node{8, nil, nil}}}, nil, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := equal(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestInvertTree(t *testing.T) {
	for _, tc := range []struct {
		name    string
		n, want *node
	}{
		{"test-1", nil, nil},
		{"test-2", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, &node{6, &node{5, nil, nil}, &node{7, nil, nil}}}, &node{4, &node{6, &node{7, nil, nil}, &node{5, nil, nil}}, &node{2, &node{3, nil, nil}, &node{1, nil, nil}}}},
		{"test-3", &node{4, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, nil}, &node{4, nil, &node{2, &node{3, nil, nil}, &node{1, nil, nil}}}},
		{"test-4", &node{3, &node{5, &node{6, nil, nil}, &node{2, &node{7, nil, nil},
			&node{4, nil, nil}}}, &node{1, &node{0, nil, nil}, &node{8, nil, nil}}}, &node{3, &node{1, &node{8, nil, nil}, &node{0, nil, nil}}, &node{5, &node{2, &node{4, nil, nil}, &node{7, nil, nil}}, &node{6, nil, nil}}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if invertTree(tc.n); !equal(tc.n, tc.want) {
				t.Errorf("got = %v, want = %v", tc.n, tc.want)
			}
		})
	}
}
