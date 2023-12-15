package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	sorted := make([]int, n)
	copy(sorted, nums)
	slices.Sort(sorted)
	deduped := slices.Compact(sorted)

	ans := make([]string, n)
	for i, v := range nums {
		pos, _ := slices.BinarySearch(deduped, v)
		ans[i] = strconv.Itoa(pos + 1)
	}

	fmt.Println(strings.Join(ans, " "))
}
