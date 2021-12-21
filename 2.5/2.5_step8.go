//На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
//
//Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").
//
//Sample Input:
//
//топот
//Sample Output:
//
//Палиндром
package main

import "fmt"

func ToReversed(text string) string {
	rs := []rune(text)

	reversedBs := make([]rune, len(rs))
	for i := 0; i < len(reversedBs); i++ {
		reversedBs[i] = rs[len(reversedBs)-i-1]
	}
	return string(reversedBs)
}

func main() {
	var text string
	_, _ = fmt.Scan(&text)

	reversed := ToReversed(text)

	if text == reversed {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
