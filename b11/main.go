package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	slices.Sort(nums)

	var q, x int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &x)
		pos, _ := slices.BinarySearch(nums, x)
		fmt.Println(pos)
	}
}
