package main

import "fmt"

// this is a comment

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	variables()
}

func variables() {
	var a1 string = "1. init variables with form - var a1 string, after init "
	fmt.Println(a1)
	var a2 string = "2. init variables with form - var a1 string = \"\" "
	fmt.Println(a2)
	a3 := "3. init variables with form - a3 := \"\", compiler define type"
	fmt.Println(a3)
	const a4 string = "4. init const variables - const a4 string = \"\""
	fmt.Println(a4)

}
