package main

import "fmt"
import "math"

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	primSum := 0
	secSum := 0
	for i := 0; i < N; i++ {
		// fmt.Println("i", i)
		for j := 0; j < N; j++ {
			// fmt.Println("j", j)
			tmp := 0
			fmt.Scanf("%d", &tmp)

			if j == i {
				// fmt.Println("prim", tmp)
				primSum += tmp
			}
			if j == N-i-1 {
				// fmt.Println("sec", tmp)
				secSum += tmp
			}
		}

	}
	// fmt.Println(primSum, secSum)
	fmt.Println(math.Abs(float64(primSum) - float64(secSum)))

}
