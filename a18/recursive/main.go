package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dp   [][]bool
	nums []int
)

func partialSum(rest int, cur int) bool {
	if rest == 0 {
		return true
	}

	if cur < 0 || rest < 0 {
		return false
	}

	if dp[cur][rest] {
		return true
	}

	ok := partialSum(rest, cur-1) || partialSum(rest-nums[cur], cur-1)
	dp[cur][rest] = ok
	return ok
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, s int
	fmt.Fscan(reader, &n, &s)

	nums = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, s+1)
	}
	ok := partialSum(s, n-1)
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
