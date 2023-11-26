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

	intervals := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &intervals[i])
	}

	l := 1
	r := 1000000000
	for l < r {
		m := (l + r) / 2

		count := 0
		for i := 0; i < n; i++ {
			count += m / intervals[i]
		}

		if count < k {
			l = m + 1
		} else {
			r = m
		}
	}

	fmt.Println(r)
}
