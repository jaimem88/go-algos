package main

import "fmt"
import "sort"

func main() {
	initial := 0
	fmt.Scanf("%d", &initial)
	in := make([]int, initial)
	for i := range in {
		fmt.Scan(&in[i])

	}

	sort.Ints(in)
	fmt.Println(sumCandles(in))
}
func sumCandles(arr []int) int {
	max := arr[len(arr)-1]
	sum := 0
	for _, v := range arr {
		if v == max {
			sum++
		}

	}
	return sum
}
