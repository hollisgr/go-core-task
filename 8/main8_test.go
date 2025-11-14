package main

import (
	"reflect"
	"sort"
	"sync"
	"testing"
)

func TestCustomWaitGroupNormal(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}
	nums3 := []int{7, 8, 9}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	test := []int{}

	wg := MyWaitGroup{}
	mu := sync.Mutex{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		for _, v := range nums1 {
			mu.Lock()
			test = append(test, v)
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for _, v := range nums2 {
			mu.Lock()
			test = append(test, v)
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for _, v := range nums3 {
			mu.Lock()
			test = append(test, v)
			mu.Unlock()
		}
	}()

	wg.Wait()

	sort.Ints(test)

	expLen := len(exp)
	testLen := len(test)

	if expLen != testLen {
		t.Errorf("expected: %v, got: %v", expLen, testLen)
	}

	if !reflect.DeepEqual(exp, test) {
		t.Errorf("slices are not equal! exp: %v, got: %v", exp, test)
	}
}
