package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var d, n int
	fmt.Fscan(reader, &d, &n)

	var first, last int
	diffs := make([]int, d+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &first, &last)
		diffs[first] += 1
		if last != d {
			diffs[last+1] -= 1
		}
	}

	totals := make([]int, d+1)
	for i := 1; i <= d; i++ {
		totals[i] = totals[i-1] + diffs[i]
		fmt.Println(totals[i])
	}
}
