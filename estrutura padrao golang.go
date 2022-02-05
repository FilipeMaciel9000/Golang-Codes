/*
Entendendo a estrutura de um programa em go:
1) Declaração de pacote principal do nosso programa, é por esse pacote que o programa irá começar;
2) Declaração do comando import, para importar códigos úteis de outros pacotes;
3) "fmt" é abreviação de format, que implementa formatação de input e output para o nosso código;
4) Declaração da função principal do nosso programa, é a porta de entrada e de saida do seu programa em go
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Olá mundo!")
}

//Aqui seu código termina!
