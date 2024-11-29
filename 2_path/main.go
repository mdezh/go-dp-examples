package main

import "fmt"

// Дано прямоугольное поле размером n*m клеток. Можно совершать шаги
// длиной в одну клетку вправо или вниз. Посчитать, сколькими способами
// можно попасть из левой верхней клетки в правую нижнюю.

// 1 1 1
// 1 2 3
// 1 3 6

func solve(n, m int) int {
	if n < 1 || m < 1 {
		return 0
	}

	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
	}
	a[0][0] = 1

	for i := range n {
		for j := range m {
			if i > 0 {
				a[i][j] += a[i-1][j]
			}
			if j > 0 {
				a[i][j] += a[i][j-1]
			}
		}
	}

	return a[n-1][m-1]
}

// Решение через комбинаторику, более эффективное
func solveComb(n, m int) int {
	if n < 1 || m < 1 {
		return 0
	}
	return binominal(n+m-2, n-1)
}

// Расчет биноминального коэффициента (число сочетаний из n по k)
// Функция взята из https://pkg.go.dev/gonum.org/v1/gonum/stat/combin#Binomial
// (убраны проверки на валидность аргументов)
func binominal(n, k int) int {
	if k > n/2 {
		k = n - k
	}

	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}

	return b
}

func solveAndPrint(n, m int) {
	fmt.Println("n:", n, "m:", m, "dp:", solve(n, m), "binominal:", solveComb(n, m))
}

func main() {
	solveAndPrint(0, 10)
	solveAndPrint(10, 0)
	solveAndPrint(1, 1)
	solveAndPrint(1, 2)
	solveAndPrint(1, 3)
	solveAndPrint(2, 1)
	solveAndPrint(3, 1)
	solveAndPrint(2, 2)
	solveAndPrint(2, 3)
	solveAndPrint(3, 2)
	solveAndPrint(3, 3)
	solveAndPrint(4, 4)
	solveAndPrint(5, 5)
	solveAndPrint(6, 6)
	solveAndPrint(10, 10)
}
