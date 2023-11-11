package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscan(reader, &n, &q)

	var num int
	numCumSum := 0
	numCumSumList := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &num)
		numCumSum += num
		numCumSumList[i] = numCumSum
	}

	var start, end int
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &start, &end)
		fmt.Println(numCumSumList[end] - numCumSumList[start-1])
	}
}
