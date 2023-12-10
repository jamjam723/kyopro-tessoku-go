package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tmp int
	fmt.Fscan(reader, &tmp)
	n := float64(tmp)

	left := 0.0
	right := 100000.0
	for right-left > 0.001 {
		mid := (left + right) / 2
		if math.Pow(mid, 3)+mid >= n {
			right = mid
		} else {
			left = mid
		}
	}

	fmt.Println(right)
}
