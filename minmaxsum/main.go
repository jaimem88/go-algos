package main

import "fmt"
import "sort"

func main() {

	in := make([]int, 5)
	for i := range in {
		fmt.Scan(&in[i])

	}

	sort.Ints(in)
	fmt.Println(sumMin(in), sumMax(in))
}

func sumMax(arr []int) int {
	sum := 0
	for k, v := range arr {
		if k == 0 {
			continue
		}
		sum += v
	}
	return sum
}

func sumMin(arr []int) int {
	sum := 0
	for k, v := range arr {
		if k == len(arr)-1 {
			continue
		}
		sum += v
	}
	return sum
}
