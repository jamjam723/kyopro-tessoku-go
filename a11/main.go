package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, x int
	fmt.Fscan(reader, &n, &x)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	k := -1
	l := 0
	r := n
	for {
		if r <= l {
			break
		}

		k = (l + r) / 2
		if nums[k] == x {
			break
		}

		if nums[k] < x {
			l = k + 1
		} else {
			r = k
		}
	}

	fmt.Println(k + 1)
}
