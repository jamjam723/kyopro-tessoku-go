package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, s int
	fmt.Fscan(reader, &n, &s)

	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, s+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			ifUse := false
			if j-nums[i] >= 0 {
				ifUse = dp[i-1][j-nums[i]]
			}
			dp[i][j] = dp[i-1][j] || ifUse
		}
	}
	if !dp[n][s] {
		fmt.Println(-1)
		return
	}

	sum := s
	selected := make([]int, 0, n)
	for i := n; i >= 1; i-- {
		if dp[i-1][sum] {
			continue
		}

		selected = append(selected, i)
		sum -= nums[i]
	}

	fmt.Println(len(selected))
	for i := len(selected) - 1; i >= 0; i-- {
		fmt.Print(selected[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
