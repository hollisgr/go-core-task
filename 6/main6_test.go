package main

import (
	"sync"
	"testing"
)

func TestRngNormal(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	params := Params{
		Min:  1,
		Max:  10,
		Size: 15,
	}

	wg.Add(1)
	go RandomGenerator(ch, &wg, params)

	test := []int{}

	for len(test) < params.Size {
		n := <-ch
		test = append(test, n)
	}

	close(ch)
	wg.Wait()

	if len(test) != params.Size {
		t.Errorf("expected: %v, got: %v", params.Size, len(test))
	}

	for _, v := range test {
		if v < params.Min || v > params.Max {
			t.Errorf("incorrect value")
		}
	}
}
