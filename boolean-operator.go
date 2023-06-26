package main

import "fmt"

func main() {
	ujian := 80
	absensi := 80

	lulusUjian := ujian > 80
	lulusAbsensi := absensi > 80

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	lulus := lulusUjian && lulusAbsensi
	lulus2 := ujian > 80 && absensi > 80

	fmt.Println(lulus)
	fmt.Println(lulus2)
}
