package main

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input *node
		want  []int
	}{
		{"tree 1", tree1(), []int{8, 4, 9, 2, 10, 5, 1, 6, 3, 7}},
		{"tree 2", tree2(), []int{1, 2, 3, 4, 5, 6, 7}},
		{"tree 3", tree3(), []int{1, 2, 3, 4, 5, 6}},
		{"tree 4", tree4(), []int{1}},
		{"tree 5", tree5(), []int{0}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			inorderTraversal(tc.input, func(v int) { got = append(got, v) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    *node
		b    *node
		want bool
	}{
		{"Not Equal", tree1(), tree2(), false},
		{"Equal", tree1(), tree1(), true},
		{"Not Equal 2", tree2(), tree3(), false},
		{"Equal One", tree4(), tree4(), true},
		{"Not Equal 3", tree4(), tree5(), false},
		{"Equal Nil", tree5(), tree5(), true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if equal(tc.a, tc.b) != tc.want {
				t.Errorf("a = %v, b = %v, got = %v, want = %v", tc.a, tc.b, equal(tc.a, tc.b), tc.want)
			}
		})
	}
}

func TestInvertTree(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input *node
		want  *node
	}{
		{"tree 1", tree1(), tree1inv()},
		{"tree 2", tree2(), tree2inv()},
		{"tree 3", tree3(), tree3inv()},
		{"tree 4", tree4(), tree4()},
		{"tree 5", tree5(), tree5()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input
			invertTree(got)
			if !equal(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
