package main

import "fmt"

// Найти n-е число Фибоначчи

// наивное решение через рекурсию
func fibRec(n int) int {
	if n <= 2 {
		return 1
	}

	return fibRec(n-2) + fibRec(n-1)
}

// рекурсия с мемоизацией
func fibRecMemo(n int) int {
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

// самое эффективное решение
func fib(n int) int {
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
	fmt.Println(fib(n))
	fmt.Println(fibRecMemo(n))
	fmt.Println(fibRec(n)) // при больших n (>40) будет заметная пауза
}
