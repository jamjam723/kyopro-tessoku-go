package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscan(reader, &a, &b)

	bases := make([]int, 61)
	bases[0] = a
	for i := 1; i < 61; i++ {
		bases[i] = (bases[i-1] * bases[i-1]) % 1000000007
	}

	ans := 1
	for i := 0; i < 61; i++ {
		if (b/int(math.Pow(2, float64(i))))%2 == 1 {
			ans *= bases[i]
			ans %= 1000000007
		}
	}

	fmt.Println(ans)
}
