package main

import (
	"fmt"
)

func main() {
	var num int = 2
	var opcao int = 1
	for opcao == 1 {
		fmt.Printf("Choose an option\n")
		fmt.Printf("1. Continue\n")
		fmt.Printf("2. Exit\n")
		fmt.Scanf("%d", &opcao)
		if opcao == 1 {
			fmt.Printf("\nType a number integer\n")
			fmt.Scanf("%d", &num)
			if num % 2 == 0 {
				fmt.Printf("Even\n")
			} else {
				fmt.Printf("Odd\n")
			}
		} else {
			fmt.Printf("Bye\n")
		}
	}
}