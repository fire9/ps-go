package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("hank", "hello world"))
}

func convert(name, course string) (a, b string) {
	name = strings.ToUpper(name)
	course = strings.Title(course)

	return name, course
}
