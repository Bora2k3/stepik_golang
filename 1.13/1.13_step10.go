//Цифровой корень
//Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
//
//Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
//
//По данному числу определите его цифровой корень.
//
//Входные данные
//
//Вводится одно натуральное число n, не превышающее 10^710
//7
// .
//
//Выходные данные
//Вывести цифровой корень числа n.
//
//Sample Input:
//
//3456
//Sample Output:
//
//9

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for n >= 10 {
		newN := 0
		for n > 0 {
			newN += n % 10
			n /= 10
		}
		n = newN
	}
	fmt.Println(n)
}
