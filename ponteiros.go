package main

import "fmt"

func main() {

	// memoria-endereco(40) <--- a <--- 40
	//variável com valor
	a := 10
	// valor de "a"
	fmt.Println("valor de \"a\":", a)

	// endereço de memória da variável "a"
	fmt.Println("endereco de memória da variável \"a\":", &a)

	// define um ponteiro inteiro e atribui o endereço da variável "a"
	var ponteiro *int = &a

	// imprime o endereço do próprio ponteiro
	fmt.Println("endereço do ponteiro:", &ponteiro)

	// imprime o valor do ponteiro, que é o endereço de "a"
	fmt.Println("valor do ponteiro, o endereço de \"a\":", ponteiro)

	// faz o de-reference, buscando o valor que está contido no endereço de "a"
	fmt.Println("valor de \"a\" usando de-reference do ponteiro", *ponteiro)
}
