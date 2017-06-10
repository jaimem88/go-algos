package main

import "log"
import "strings"

func stringBubbleSort(arr []string) {

	if len(arr) <= 1 {
		return
	}
	swap := false
	for i := 0; i < len(arr)-1; i++ {
		if strings.Compare(arr[i], arr[i+1]) > 0 {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
			swap = true
		}
	}
	if !swap {
		return
	}
	stringBubbleSort(arr[0 : len(arr)-1])
}

func bubbleSort(arr []int, desc bool) {

	if len(arr) <= 1 {
		return
	}
	swap := false

	for i := 0; i < len(arr)-1; i++ {
		f := false
		if desc { // if descending order
			f = arr[i] < arr[i+1]
		} else { // ascending order
			f = arr[i] > arr[i+1]
		}

		if f {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
			swap = true
		}
	}
	if !swap {
		return
	}
	bubbleSort(arr[0:len(arr)-1], desc)
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	log.Println("Before", a)
	bubbleSort(a, false)
	log.Println("After", a)

	b := []int{1, 2, 3, 4}
	log.Println("Before", b)
	bubbleSort(b, false)
	log.Println(b)

	c := []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	bubbleSort(c, true)
	log.Println(c)

	d := []string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine", "dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"}
	log.Println(d)
	stringBubbleSort(d)
	log.Println(d)
}
