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

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}

	fmt.Println(dp[n])
}
