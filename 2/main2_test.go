package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	src := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	exp := []int{0, 2, 4, 6, 8}
	test := sliceExample(src)
	if !reflect.DeepEqual(exp, test) {
		t.Errorf("expected: %v, got: %v", exp, test)
	}
}

func TestAddElements(t *testing.T) {
	src := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	test := addElements(src, 10)
	if !reflect.DeepEqual(exp, test) {
		t.Errorf("expected: %v, got: %v", exp, test)
	}
}

func TestCopySlice(t *testing.T) {
	src := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test := copySlice(src)
	if !reflect.DeepEqual(src, test) {
		t.Errorf("expected: %v, got: %v", src, test)
	}

	test = addElements(src, 10)
	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exp_src := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(test, exp) && !reflect.DeepEqual(src, exp_src) {
		t.Errorf("expected test: %v, got: %v", exp, test)
		t.Errorf("expected src: %v, got: %v", exp_src, src)
	}
}

func TestRemoveSlice(t *testing.T) {
	src := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test := removeSlice(src, 4)
	exp := []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(exp, test) {
		t.Errorf("expected: %v, got: %v", exp, test)
	}
}
