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

	bitsInt := make([]int, 10)
	for i := 0; i < 10; i++ {
		divisor := int(math.Pow(2, float64(i)))
		quotient := n / divisor
		bit := quotient % 2
		bitsInt[10-i-1] = bit
	}

	for i := 0; i < 10; i++ {
		fmt.Print(bitsInt[i])
	}
	fmt.Println()
}
