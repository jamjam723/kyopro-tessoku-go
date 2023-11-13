package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(reader, &t, &n)

	nums := make([]int, t+2)
	var start, end int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &start, &end)
		nums[start] += 1
		nums[end] -= 1
	}

	cumSums := make([]int, t+1)
	cumSums[0] = nums[0]
	for i := 1; i <= t; i++ {
		cumSums[i] = cumSums[i-1] + nums[i]
	}

	for i := 0; i < t; i++ {
		fmt.Println(cumSums[i])
	}
}
