package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	originalSlice := NewRandSlice()
	fmt.Println("original slice: ", originalSlice)

	example := sliceExample(originalSlice)
	fmt.Println("example slice:  ", example)

	add := addElements(example, 2025)
	fmt.Println("add slice:      ", add)

	copy := copySlice(add)
	fmt.Println("copy slice:     ", copy)

	remove := removeSlice(originalSlice, 5)
	fmt.Println("remove slice:   ", remove)

}

func removeSlice(s []int, index int) []int {
	res := []int{}
	for i, n := range s {
		if i != index {
			res = append(res, n)
		}
	}
	return res
}

func copySlice(s []int) []int {
	copy := []int{}
	for _, n := range s {
		copy = append(copy, n)
	}
	return copy
}

func addElements(s []int, num int) []int {
	return append(s, num)
}

func sliceExample(s []int) []int {
	res := []int{}
	for _, n := range s {
		if n%2 == 0 {
			res = append(res, n)
		}
	}
	return res
}

func NewRandSlice() []int {
	arr := make([]int, 10)
	for i := range arr {
		num := GetRandInt()
		arr[i] = num
	}
	return arr
}

func GetRandInt() int {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	min := 0
	max := 9
	res := min + rng.Intn(max+1)
	return res
}
