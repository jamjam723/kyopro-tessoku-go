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

	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	cumSums := make([]int, n+1)
	cumSums[0] = 0
	for i := 1; i <= n; i++ {
		cumSums[i] = cumSums[i-1] + nums[i]
	}

	ans := 0
	right := 0
	for i := 0; i < n; i++ {
		left := i
		for right < n && cumSums[right+1]-cumSums[left] <= k {
			right += 1
		}
		ans += (right - left)
	}

	fmt.Println(ans)
}
