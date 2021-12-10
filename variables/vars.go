package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	name   = os.Getenv("$USER")
	course = "Intro Golang"
	module = "4"
	clip   = 2
)

func main() {
	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}
	fmt.Println("Memory address of *course* variable is", &course)
	var ptr *string = &course
	fmt.Println("Pointing course variable at address,", ptr, "which holds this value.", *ptr)
	// updateC := updateCourse("hello world")
	// fmt.Println(updateC)
	updateCourse(&course)
}

func updateCourse(course *string) string {
	*course = "Getting Start Docker"
	fmt.Println("Update course to", *course)
	return *course
}
