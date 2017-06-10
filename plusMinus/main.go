package main

import "fmt"

// import "math"

func main() {
	N := float64(0)
	fmt.Scanf("%f", &N)
	pos := float64(0)
	neg := float64(0)
	zeroes := float64(0)
	for i := float64(0); i < N; i++ {
		tmp := float64(0)
		fmt.Scanf("%f", &tmp)
		if tmp > 0 {
			pos++
		} else if tmp < 0 {
			neg++
		} else {
			zeroes++
		}
	}
	// fmt.Println(primSum, secSum)
	fmt.Println(pos / N)
	fmt.Println(neg / N)
	fmt.Println(zeroes / N)

}
