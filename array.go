package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Fajar"
	names[1] = "Dwi"
	names[2] = "Yunanto"

	values := [3]int{
		90,
		98,
		66,
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(values)
	fmt.Println(len(values))
}
