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

	var result int
	cumSum := 0
	resultCumSumList := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &result)
		cumSum += result
		resultCumSumList[i] = cumSum
	}

	var start, end int
	var q int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &start, &end)
		tryNum := end - start + 1
		atariNum := resultCumSumList[end] - resultCumSumList[start-1]
		hazureNum := tryNum - atariNum
		if atariNum == hazureNum {
			fmt.Println("draw")
		} else if atariNum > hazureNum {
			fmt.Println("win")
		} else {
			fmt.Println("lose")
		}
	}
}
