//На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу: число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после запятой. Но если число равно или меньше нуля - выводить:
//"число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине. А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).
//
//Sample input:
//-000012.2123
//Sample output:
//число -12.21 не подходит
//
//Sample input:
//1000001
//Sample output:
//1.000001e+06
//
//Sample Input:
//
//12.12345678
//Sample Output:
//
//146.9782

package main

import "fmt"

func main() {
	var n float32
	fmt.Scan(&n)

	pow_n := n * n

	if n <= 0 {
		fmt.Printf("число %2.2f не подходит", n)
	} else if pow_n > 10_000 {
		fmt.Printf("%e", n)
	} else {
		fmt.Printf("%.4f", pow_n)
	}
}
