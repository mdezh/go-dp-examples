package main

import "fmt"

// Найти n-е число Фибоначчи

func FibRec(n int) int {
	if n <= 2 {
		return 1
	}

	return FibRec(n-2) + FibRec(n-1)
}

func FibRecMemo(n int) int {
	memo := make([]int, n+1)

	var fib func(int) int
	fib = func(n int) int {
		if n <= 2 {
			return 1
		}

		if memo[n] == 0 {
			memo[n] = fib(n-1) + fib(n-2)
		}

		return memo[n]
	}

	return fib(n)
}

func Fib(n int) int {
	if n <= 2 {
		return 1
	}

	prev, cur := 1, 1

	for i := 3; i <= n; i++ {
		prev, cur = cur, prev+cur
	}

	return cur
}

func main() {
	n := 45
	fmt.Println(n)
	fmt.Println(FibRec(n))
	fmt.Println(FibRecMemo(n))
	fmt.Println(Fib(n))

}
