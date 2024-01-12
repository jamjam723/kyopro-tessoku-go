package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"golang.org/x/exp/slices"
)

type box struct {
	v int
	h int
}

func maxInt(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)

	boxes := make([]box, n)
	var v int
	var h int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &v)
		fmt.Fscan(reader, &h)
		boxes[i] = box{
			v,
			h,
		}
	}
	slices.SortFunc(boxes, func(a, b box) bool {
		if a.v == b.v {
			return a.h < b.h
		}
		return b.v < a.v
	})

	dp := make([]int, n)
	maxBoxes := make([]box, n+1)
	maxBoxes[0] = box{
		v: 500001,
		h: 500001,
	}
	for i := 0; i < n; i++ {
		pos, _ := slices.BinarySearchFunc(maxBoxes, boxes[i], func(a, b box) int {
			return b.h - a.h
		})
		dp[i] = maxInt(1, pos)
		maxBoxes[pos] = boxes[i]
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = maxInt(ans, dp[i])
	}
	fmt.Println(ans)
}
