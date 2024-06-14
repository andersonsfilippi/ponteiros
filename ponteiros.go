package main

import "fmt"

type Carro struct {
	Name string
}

func (c Carro) andou() {
	// sempre que o método andou for chamado, o nome será BMW
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func (c *Carro) andou2() {
	// sempre que o método andou for chamado, o nome será BMW
	// inclusive no ponteiro. fora deste escopo.
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

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

	// mudando o valor de "a" usando o ponteiro
	*ponteiro = 50
	fmt.Println("novo valor do ponteiro via de-reference:", *ponteiro)
	fmt.Println("valor de \"a\":", a)

	// exemplo de alteração de valor de uma variável usando uma função sem retorno
	variavel := 10
	fmt.Println("valor variavel:", variavel)
	abc(&variavel)
	fmt.Println("novo valor variavel:", variavel)

	// exemplo de método com alteração de variáveis em escopos diferentes
	carro := Carro{
		Name: "Ka",
	}
	// vai imprimir o BMW
	carro.andou()
	// aqui vai imprmir o nome do carro que foi atribuído ao invés de imprimir
	// o nome que está no método andou()
	fmt.Println("nome do carro:", carro.Name)

	// agora o exemplo usando um ponteiro
	carro2 := Carro{
		Name: "Ka",
	}
	// vai imprimir o BMW
	carro2.andou2()
	// também vai imprimir o BMW
	fmt.Println("nome do carro2:", carro2.Name)

}

func abc(a *int) {
	*a = 200
}
