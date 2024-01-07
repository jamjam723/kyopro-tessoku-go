package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func distance(toX, fromX, toY, fromY int) float64 {
	return math.Sqrt(math.Pow(float64(toX)-float64(fromX), 2) + math.Pow(float64(toY)-float64(fromY), 2))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	xList := make([]int, n)
	yList := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &xList[i])
		fmt.Fscan(reader, &yList[i])
	}

	dp := make([][]float64, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = make([]float64, n)
	}
	inf := 100000000.0
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0.0
	for i := 0; i < 1<<n-1; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if (i/(1<<k))%2 == 1 {
					continue
				}

				dp[i+(1<<k)][k] = math.Min(dp[i+(1<<k)][k], dp[i][j]+distance(xList[k], xList[j], yList[k], yList[j]))
			}
		}
	}

	fmt.Println(dp[(1<<n)-1][0])
}
