package main

import "fmt"

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j >= N-1-i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
			if j == N-1 {
				fmt.Print("\n")
			}
		}
	}

}
