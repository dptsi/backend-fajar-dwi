package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "fajar",
		"address": "tulungagung",
	}

	person["Title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["title"])
	fmt.Println(person["Title"])

	delete(person, "Title")

	fmt.Println(person)
}
