/*
Quando declaramos uma variável com a palavra-chave var e não atribuímos nenhum valor a esta variável;
Automaticamente o compilador designa para esta um valor padrão, chamado de valor zero.
Os zero em GO:
->String: ;
->Integer: 0;
->Float: 0;
->Bool: false;
Obs: Em Go existe um caractere que permite "jogar fora" um valor, esse caractere é _
_ permite que você diga ao programa que NÃO vai utilizar o valor retornado por uma função

*/
package main

import (
	"fmt"
)

var a int
var b float64
var c string
var d bool

func main() {
	//
	fmt.Printf("O valor na variavel a é %v; do tipo: %T\n", a, a)
	fmt.Printf("O valor da variável b é %v; do tipo: %T\n", b, b)
	fmt.Printf("O valor da variável C é %v; do tipo: %T\n", c, c)
	fmt.Printf("O valor da variável d é %v; do tipo: %T\n", d, d)
}
