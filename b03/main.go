package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	costs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &costs[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if costs[i]+costs[j]+costs[k] == 1000 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}
