//Estudando variáveis de tipos primitivos em Go
//Declarar variaveis em go pode ser feito de modo implícito ou explicito e de maneira multipla ou singular
/*
O menor elemento de um programa que expressa uma ação a ser executada é a declaração de uma variável;
A combinação de um ou mais variáveis, que a linguagem interpreta e usa para produzir outro valor é uma expressão;
Em Go podemos criar nossos próprios tipos.
Tipos primitivos em Go:
->String: (nome := "Avatar" (Declaração implícita singular com operador curto de declaração)) trabalha com cadeia de caracteres, basicamente são textos;
->Integer: (idade, nome := 112, "Avatar" (Declaração implícita múltipla com operador curto de declaração)) trabalha com números inteiro, ou seja, sem vírgulas;
->Float: (var altura float64 = 1.45(Declaração explicita com valor definido)) trabalha com números decimais com vírgula.
  Existe o float32 e o float64, quando declaramos uma variável com valor decimal que possui virgula, por padrão ela será o float64;
  Então virou convenção usarmos apenas o float64;
->Bool: (var careca bool = true(Declaração explicita com valor definido)) trabalha unicamente com verdadeiros e falsos;
*/
package main

import (
	"fmt"
)
//Declaração de variáveis no corpo do importe:
var altura = 1.45

func main(){
//Declaração de variáveis no corpo da função:
nome, idade := "Avatar", 112
var careca bool = true
//Comando para mostrar os valores nas variáveis (%v) e o tipo da variável (%T) com formatação:
//Perceba que ao usar o parametro  %v %T precisamos usar a variável careca duas vezes;
//Uma para moatrar o conteudo da variável (%v) e outra para mostrar o tipo (%T) da variável. 
	fmt.Printf(" O %v Têm %v anos de idade;\n Têm %v de altura;\n E é careca?\t %v %T",nome, idade, altura, careca, careca)
}
