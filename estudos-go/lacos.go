package main

import "fmt"

func main(){

	// condicao simples
	var m int
	m = 1
	for m <= 5{
		fmt.Printf("m = %d\n", m)
		m += 1
	}

	// condicao classica
	for e:=6; e<=10; e++{
		fmt.Printf("m = %d\n", e)
	}
}