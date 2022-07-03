package main

import "fmt"

func main()  {

	// for and while loop

	// simple for
	soma := 0
	for i:= 1; i < 5; i++{
		soma += i
		fmt.Println(soma)
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(soma)

	/*
	simple while
	for [condition]{
		code
	}
	*/
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}