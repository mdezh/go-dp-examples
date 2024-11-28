package main

import (
	"fmt"
)

// Дана последовательность целых чисел. Необходимо найти длину ее самой
// длинной строго возрастающей подпоследовательность (подпоследовательность
// получается из последовательности удалением 0 или больше ее элементов).

func solve(a []int) int {
	if len(a) < 2 {
		return len(a)
	}

	// длина максимальной найденной подпоследовательности
	maxLen := 1

	// длины максимальных подпоследовательностей для участков последовательности от 0 до i
	maxLens := make([]int, len(a))
	maxLens[0] = 1

	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] <= a[j] {
				continue
			}

			if maxLens[j] > maxLens[i] {
				maxLens[i] = maxLens[j]
			}
		}

		maxLens[i]++
		if maxLens[i] > maxLen {
			maxLen = maxLens[i]
		}
	}

	return maxLen
}

func solveAndPrint(a []int) {
	fmt.Println(a, solve(a))
}

func main() {
	solveAndPrint(nil)
	solveAndPrint([]int{})
	solveAndPrint([]int{1})
	solveAndPrint([]int{1, 1})
	solveAndPrint([]int{1, 2})
	solveAndPrint([]int{1, 2, 3})
	solveAndPrint([]int{1, 2, 0})
	solveAndPrint([]int{5, 2, 3})
	solveAndPrint([]int{1, 3, 5})
	solveAndPrint([]int{1, 2, 3, 4, 5})
	solveAndPrint([]int{1, 1, 3, 4, 5})
	solveAndPrint([]int{1, 2, 1, 4, 5})
	solveAndPrint([]int{1, 2, 1, 4, 4})
	solveAndPrint([]int{2, 1, 2, 4, 4})
	solveAndPrint([]int{2, 1, 5, 4, 4})
}
