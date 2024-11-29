package main

import "fmt"

// Посчитать число последовательностей нулей и единиц длины n,
// в которых не встречаются две идущие подряд единицы

// ... | n-2 | n-1 | 0
// ... | n-2 |  0  | 1

func solve(n int) int {
	if n < 1 {
		return 0
	}

	prev, cur := 2, 3

	switch n {
	case 1:
		return prev
	case 2:
		return cur
	}

	for i := 3; i <= n; i++ {
		cur, prev = cur+prev, cur
	}

	return cur
}

func main() {
	n := 64
	fmt.Println(n, solve(n))
}
