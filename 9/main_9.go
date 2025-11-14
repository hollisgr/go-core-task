package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := Input()

	uint8Chan := make(chan uint8)
	float64Chan := make(chan float64)

	go func() {
		for n := range float64Chan {
			fmt.Printf("value: %f\n", n)
		}
	}()

	go TransformAndCube(uint8Chan, float64Chan)

	for _, n := range nums {
		uint8Chan <- n
	}

	close(uint8Chan)

	fmt.Println("all done")
}

func TransformAndCube(in chan uint8, out chan float64) {
	for n := range in {
		cube := float64(n) * float64(n) * float64(n)
		out <- cube
	}
	close(out)
}

func Input() []uint8 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a numbers separated with space (0 - 255)")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	res := []uint8{}

	for t := range strings.SplitSeq(input, " ") {
		n, _ := strconv.Atoi(t)
		if n > 255 {
			fmt.Println("value is too large, exiting")
		}
		res = append(res, uint8(n))
	}
	return res
}
