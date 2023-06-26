package main

import "fmt"

func main() {
	name1 := "Fajar"
	name2 := "Dwi"

	number1 := 70
	number2 := 60

	result := name1 == name2
	result2 := number1 <= number2
	result3 := number1 > number2
	result4 := number1 == number2

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
