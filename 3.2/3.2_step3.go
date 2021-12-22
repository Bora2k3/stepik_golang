//Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.
//
//Считывать и выводить ничего не нужно!
//
//Считайте что функция main и пакет main уже объявлены!
//
//Sample Input:
//
//Sample Output:
package main

import "fmt"

func convert(num int64) uint16 {
	return uint16(num)
}

func main() {
	var num = int64(1000)
	var num2 = convert(num)
	fmt.Printf("%T %v\n", num2, num2)
}
