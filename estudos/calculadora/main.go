package main

import "fmt"

func main()  {
	var opcao, n1, n2 int
	fmt.Println("------------------ ")
	fmt.Println("--- CALCULADORA ---")
	fmt.Println("------------------ ")
	fmt.Println("1. SOMA")
	fmt.Println("2. SUBTRAÇÃO")
	fmt.Println("3. MULTIPLICAÇÃO")
	fmt.Println("4. DIVISÃO")
	fmt.Println(" ------------------ ")
	fmt.Println("Escolha uma opção:")
	fmt.Scanf("%d", &opcao)
	fmt.Println("------------------ ")
	if opcao == 1 {
		fmt.Println("1. SOMA")
		fmt.Println("------------------ ")
		fmt.Scanf("%d %d", n1, n2)
	} else if opcao == 2 {
		fmt.Println("2. SUBTRAÇÃO")
		fmt.Println("------------------ ")
		fmt.Scanf("Numero 1: %d", &n1)
		fmt.Scanf("Numero 2: %d", &n2)
	} else if opcao == 3 {
		fmt.Println("3. MULTIPLICAÇÃO")
		fmt.Println("------------------ ")
		fmt.Scanf("Numero 1: %d", &n1)
		fmt.Scanf("Numero 2: %d", &n2)
	} else if opcao == 4 {
		fmt.Println("4. SUBTRAÇÃO")
		fmt.Println("------------------ ")
		fmt.Scanf("Numero 1: %d", &n1)
		fmt.Scanf("Numero 2: %d", &n2)
	}
	fmt.Println("Numero 1: ", n1)
	fmt.Println("Numero 2: ", n2)
}