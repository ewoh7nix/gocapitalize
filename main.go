package main

import (
	"fmt"
	"gocapitalize/capitalize"
)

func main() {
	word := "i want to became a capitalist"
	cap_word := capitalize.Capitalize(word)
	fmt.Println(cap_word)
}
