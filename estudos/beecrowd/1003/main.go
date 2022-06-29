package main

import "fmt"

func main()  {
	var a int
	var b int

	fmt.Sscan("%d", &a)
	fmt.Sscan("%d", &b)

	print("SOMA = ", a+b)
}