package main

import (
	"log"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	search := 3
	log.Println("a", a, search, " at index:", binarySearch(search, a, 0))
	log.Println("b", b, search, " at index:", binarySearch(search, b, 0))

}

func binarySearch(val int, arr []int, lo, hi int) (index int) {
	arrLen := len(arr)
	// log.Println(arr)
	if arrLen < 1 {
		return -1
	}
	mid := lo +(hi - lo)/2
	// log.Println(" current", current, " mid ", mid)
	if arr[mid] == val {
		log.Println("never here?", arr[mid])
		return mid
	} else if val > arr[mid] {
		
		//check right
		return binarySearch(val, arr[mid+1:], mid)
	} else {
		//check left
		current = mid
		return binarySearch(val, arr[:mid-1], mid)
	}

}
