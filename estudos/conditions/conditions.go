package main

import (
	"fmt"
)

func main()  {
	
	// variables can be defined by their types
	var a int 
	var c int

	// or it can be declared without the type
	// recebendo apenas :=
	// the colons means we're declaring the variable 
	// so we can't use := on the same variable twice
	b := "hello"

	// once is declared, to change its value we only 
	// need to use the '='
	b = "world"
	a = 2

	fmt.Println(b)
	fmt.Println("a = ", a)
	fmt.Scanln("%d", &c)
	if a == 1{
		fmt.Println("Bem vindo")
	} else if a == 2{
		fmt.Println("Welcome")
	} else {
		fmt.Println()
	}
}