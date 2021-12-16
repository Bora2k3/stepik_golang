// По данному трехзначному числу определите, все ли его цифры различны.
//
//Формат входных данных
//На вход подается одно натуральное трехзначное число.
//
//Формат выходных данных
//Выведите "YES", если все цифры числа различны, в противном случае - "NO".
//
//Sample Input 1:
//
//237
//Sample Output 1:
//
//YES
//Sample Input 2:
//
//117
//Sample Output 2:
//
//NO

package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	if n[0] != n[1] && n[0] != n[2] && n[1] != n[2] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
