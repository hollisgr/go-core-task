package main

import (
	"reflect"
	"testing"
)

func TestDiffNormal(t *testing.T) {
	str1 := []int{123, 234, 555, 112233, 90855}
	str2 := []int{9605846, 112233, 37465, 88775, 123}
	exp := []int{112233, 123}

	ok, test := FindDiff(str1, str2)

	if !ok {
		t.Errorf("expected: %v, got: %v", true, ok)
	}

	if !reflect.DeepEqual(exp, test) {
		t.Errorf("expected: %v, got: %v", exp, test)
	}
}

func TestDiffEmptySlice(t *testing.T) {
	str1 := []int{123, 234, 555, 112233, 90855}
	str2 := []int{9605846, 11233, 37465, 88775, 9948}
	exp := []int{}

	ok, test := FindDiff(str1, str2)

	if ok {
		t.Errorf("expected: %v, got: %v", false, ok)
	}

	if !reflect.DeepEqual(exp, test) {
		t.Errorf("expected: %v, got: %v", exp, test)
	}
}
