//Количество минимумов
//Найдите количество минимальных элементов в последовательности.
//
//Входные данные
//
//Вводится натуральное число N, а затем N целых чисел последовательности.
//
//Выходные данные
//
//Выведите количество минимальных элементов последовательности.
//
//Sample Input:
//
//3
//21
//11
//4
//Sample Output:
//
//1

package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	fmt.Scan(&x)
	min := x
	number := 1

	for i := 1; i < n; i++ {
		fmt.Scan(&x)

		if x < min {
			min = x
			number = 0
		}
		if x == min {
			number++
		}
	}
	fmt.Println(number)
}
