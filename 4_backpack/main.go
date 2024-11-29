package main

import (
	"errors"
	"fmt"
)

// Дан рюкзак грузоподъемностью LC и n предметов, каждый из которых имеет свой вес Wi
// и ценность Vi. Все параметры задаются целыми числами. Необходимо найти максимальную
// ценность предметов, которую можно унести в рюкзаке.

// Пример решения для LC 4, весов 4, 1, 3 и стоимостей 5000, 3500, 2000
//
//          0    1    2    3    4
//
//          0    0    0    0    0
// 5000(4)  0    0    0    0 5000
// 3500(1)  0 3500 3500 3500 5000
// 2000(3)  0 3500 3500 3500 5500
//
// в другом порядке
//          0    0    0    0    0
// 5000(4)  0    0    0    0 5000
// 2000(3)  0    0    0 2000 5000
// 3500(1)  0 3500 3500 3500 5500

var (
	errNoVariants  = errors.New("no variants")
	errInvalidData = errors.New("invalid data")
)

func solve(lc int, w, v []int) (int, error) {
	if err := validate(lc, w, v); err != nil {
		return 0, err
	}

	t := make([][]int, len(w)+1)
	t[0] = make([]int, lc+1)

	for i := 1; i <= len(w); i++ {
		t[i] = make([]int, lc+1)

		for j := 1; j <= lc; j++ {
			t[i][j] = t[i-1][j]
			if w[i-1] <= j {
				t[i][j] = max(t[i][j], t[i-1][j-w[i-1]]+v[i-1])
			}
		}
	}

	res := t[len(w)][lc]
	if res > 0 {
		return res, nil
	}

	return 0, errNoVariants
}

func validate(lc int, w, v []int) error {
	if lc < 1 || len(w) < 1 || len(w) != len(v) {
		return errInvalidData
	}

	for i := 0; i < len(w); i++ {
		if w[i] < 1 || v[i] < 1 {
			return errInvalidData
		}
	}

	return nil
}

func solveAndPrint(lc int, w, v []int) {
	fmt.Println("Load capacity:", lc)
	fmt.Println("Weights:", w)
	fmt.Println("Values:", v)

	if res, err := solve(lc, w, v); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println("-----")
}

func main() {
	solveAndPrint(0, []int{4, 1, 3}, []int{4, 3, 2})
	solveAndPrint(5, nil, []int{4, 3, 2})
	solveAndPrint(5, []int{1, 3}, []int{4, 3, 2})
	solveAndPrint(5, []int{4, 1, 3}, []int{0, 3, 2})

	solveAndPrint(1, []int{4, 2, 3}, []int{5000, 3500, 2000})

	solveAndPrint(5, []int{4, 1, 3}, []int{5000, 3500, 2000})
	solveAndPrint(4, []int{4, 1, 3}, []int{5000, 3500, 2000})
}
