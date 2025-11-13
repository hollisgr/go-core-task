package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMultiplexerNormal(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}
	nums3 := []int{7, 8, 9}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	channels := []<-chan int{}

	channels = append(channels, NewChannel(nums1))
	channels = append(channels, NewChannel(nums2))
	channels = append(channels, NewChannel(nums3))

	res := Multiplexer(channels)

	test := []int{}

	for n := range res {
		test = append(test, n)
	}

	expLen := len(exp)
	testLen := len(test)

	if testLen != expLen {
		t.Errorf("expected: %d, got: %d", expLen, testLen)
	}

	sort.Ints(test)

	if !reflect.DeepEqual(exp, test) {
		t.Error("slices are not equal")
	}

}

func NewChannel(nums []int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for _, n := range nums {
			res <- n
		}
	}()
	return res
}
