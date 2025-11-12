package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Params struct {
	Min  int
	Max  int
	Size int
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	params := Input()

	output := make([]int, 0, params.Size)

	wg.Add(1)
	go RandomGenerator(ch, &wg, params)

	for len(output) < params.Size {
		n := <-ch
		output = append(output, n)
	}

	close(ch)

	wg.Wait()

	fmt.Println("result:", output)
}

func RandomGenerator(ch chan int, wg *sync.WaitGroup, p Params) {
	defer wg.Done()
	counter := 0
	for counter < p.Size {
		ch <- GetRandInt(p.Min, p.Max)
		counter++
	}
}

func GetRandInt(min int, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	res := min + rng.Intn(max-min+1)
	return res
}

func Input() Params {
	res := Params{}
	fmt.Println("Enter a min number")
	fmt.Scanf("%d", &res.Min)
	fmt.Println("Enter a max number")
	fmt.Scanf("%d", &res.Max)
	fmt.Println("Enter a size")
	fmt.Scanf("%d", &res.Size)
	return res
}
