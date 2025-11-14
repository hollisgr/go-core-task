package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestChannelTransformNormal(t *testing.T) {
	nums := []uint8{1, 2, 3, 4}
	in := make(chan uint8)
	out := make(chan float64)

	test := []float64{}
	exp := []float64{1, 8, 27, 64}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range out {
			test = append(test, n)
		}
	}()

	go TransformAndCube(in, out)

	for _, n := range nums {
		in <- n
	}

	close(in)

	wg.Wait()

	if !reflect.DeepEqual(exp, test) {
		t.Errorf("expected: %v, got: %v", exp, test)
	}
}
