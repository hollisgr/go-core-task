package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter numbers separated by spaces")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	fmt.Println("Enter numbers separated by spaces")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	slice1 := StringToIntSlice(input1)
	slice2 := StringToIntSlice(input2)

	fmt.Println("First slice:  ", slice1)
	fmt.Println("Second slice: ", slice2)

	ok, res := FindDiff(slice1, slice2)
	fmt.Printf("Success: %v, result: %v\n", ok, res)
}

func FindDiff(s1 []int, s2 []int) (bool, []int) {
	res := []int{}
	data := map[int]int{}
	counter := 0
	for _, v := range s1 {
		data[v] = rand.Intn(10) // опять же не хотелось создавать map[int]bool, поэтому записываю рандомное значение от 0 до 9
	}

	for _, v := range s2 {
		_, ok := data[v]
		if ok {
			res = append(res, v)
			counter++
		}
	}

	if counter == 0 {
		return false, res
	}

	return true, res
}

func StringToIntSlice(s string) []int {
	res := []int{}
	for t := range strings.SplitSeq(s, " ") {
		t = strings.TrimSpace(t)
		n, _ := strconv.Atoi(t)
		res = append(res, n)
	}
	return res
}
