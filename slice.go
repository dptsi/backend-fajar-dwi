package main

import "fmt"

func main() {
	month := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice := month[4:7]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := month[11:]
	fmt.Println(slice2)

	newSlice := make([]string, 2, 5)

	copySlice := make([]string, len(newSlice), cap(newSlice))

	fmt.Println(newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
