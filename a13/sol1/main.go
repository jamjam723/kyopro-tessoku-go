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

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		left := i
		right := n
		for right-left > 1 {
			m := (left + right) / 2
			if nums[m]-nums[i] <= k {
				left = m
			} else {
				right = m
			}
		}

		count += left - i
	}
	fmt.Println(count)
}
