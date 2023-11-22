package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	caps := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &caps[i])
	}

	cumMaxFromL := make([]int, n+2)
	cumMaxFromR := make([]int, n+2)
	for i := 1; i <= n; i++ {
		cumMaxFromL[i] = int(math.Max(float64(cumMaxFromL[i-1]), float64(caps[i])))
		rIndex := n - i + 1
		cumMaxFromR[rIndex] = int(math.Max(float64(cumMaxFromR[rIndex+1]), float64(caps[rIndex])))
	}

	var d, l, r int
	fmt.Fscan(reader, &d)
	for i := 0; i < d; i++ {
		fmt.Fscan(reader, &l, &r)
		maxCap := int(math.Max(float64(cumMaxFromL[l-1]), float64(cumMaxFromR[r+1])))
		fmt.Println(maxCap)
	}
}
