package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter first line separated by spaces")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	fmt.Println("Enter second line separated by spaces")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	slice1 := strings.Split(input1, " ")
	slice2 := strings.Split(input2, " ")

	fmt.Println("First slice: ", slice1)
	fmt.Println("Second slice: ", slice2)

	res := FindDiff(slice1, slice2)
	fmt.Println("difference: ", res)
}

func FindDiff(s1 []string, s2 []string) []string {
	res := []string{}

	data := map[string]int{}

	for _, str := range s2 {
		// записываю длину строки в значение, т.к. решил не использовать map[string]bool
		data[str] = len(str)
	}

	for _, str := range s1 {
		_, ok := data[str]
		if !ok {
			res = append(res, str)
		}
	}
	return res
}
