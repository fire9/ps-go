package main

import (
	"fmt"
	"os"
	"reflect"
)

var (
	name   string
	course string
	module int
	clip   int
)

func main() {
	fmt.Println(os.Getenv("USER"))
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	name, course := "Hank", "Intro Golang"
	module, clip := 4, 2
	SayHello(name, course, module, clip)
}

func SayHello(name, course string, module, clip int) {
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
}
