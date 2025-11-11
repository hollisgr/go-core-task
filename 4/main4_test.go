package main

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	str1 := []string{"test", "golang", "diff", "strings", "true"}
	str2 := []string{"test", "golang", "water", "juice", "false"}
	exp := []string{"diff", "strings", "true"}

	test := FindDiff(str1, str2)
	if !reflect.DeepEqual(exp, test) {
		t.Errorf("expected: %v, got: %v", exp, test)
	}
}
