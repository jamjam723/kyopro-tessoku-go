package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func subSum(nums []int) []int {
	subSum := make([]int, 1<<len(nums))
	for i := 0; i < (1 << (len(nums))); i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			wari := 1 << j
			if (i/wari)%2 == 1 {
				sum += nums[j]
			}
		}
		subSum[i] = sum
	}
	return subSum
}

func printAns(yes bool) {
	if yes {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	subSum1 := subSum(nums[:n/2])
	subSum2 := subSum(nums[n/2:])
	slices.Sort(subSum2)

	ok := false
	for i := 0; i < (1 << (n / 2)); i++ {
		diff := k - subSum1[i]
		_, ok = slices.BinarySearch(subSum2, diff)
		if ok {
			break
		}
	}
	printAns(ok)
}
