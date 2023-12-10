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
	r := 0
	for i := 0; i < n; i++ {
		for r < n-1 && nums[r+1]-nums[i] <= k {
			r += 1
		}
		count += r - i
	}
	fmt.Println(count)
}
