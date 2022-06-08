package main

import "fmt"

// há várias formas para declarar variáveis:

// em grupo
var (
	nome string = "José"
	idade int 
	casado bool
	estado_civil string
)

func main(){

	// varáveis previamente declaradas apenas recebem "="	
	idade = 32

	// novas variaveis devem ser definidas com ":=", com isso ele recebe o tipo da variável
	carteira := 15.35
	estado_civil := "solteiro"
	
	if (casado == true) {
		estado_civil = "casado"
	}

	fmt.Println(nome, "tem", idade, "anos")
	fmt.Println(nome, "é", estado_civil)
	fmt.Println("Ele tem R$", carteira)

	// ou individual
	nome := "Hello world"
	fmt.Println(nome)

	var idade int
	idade = 18
	fmt.Println(idade)

	// para o print formatado:
	var i int
	i = 5
	var message string
	message = "Host not found"
	fmt.Printf("%d %s\n", i, message)

}