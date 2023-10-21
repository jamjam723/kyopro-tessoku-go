package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)

	redCards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &redCards[i])
	}
	blueCards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &blueCards[i])
	}

	for i := 0; i < n; i++ {
		redCard := redCards[i]
		for j := 0; j < n; j++ {
			if redCard+blueCards[j] == k {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}
