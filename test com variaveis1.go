//Usando variáveis pela primeira vez
/*
Entendendo a estrutura de um programa em go:
1) Declaração de pacote principal do nesso programa, é por esse pacote que o programa ira começar;
2) Declaração do comando import, para importar códigos úteis de outros pacotes;
3) "fmt" é abreviação de format, que implementa formatação de input e output para o nosso código;
4) Declaração da função principal do nesso programa, é a porta de entrada e de saida do seu programa em go
*/
package main //(1)

import ( //(2)
	"fmt" //(3)
)

func main() { //(4)
	//Aqui seu código começa;
	//Usamos o operador curto de declaração(A marmota :=) para declarar três variáveis diferentes em apenas uma linha;
	//Preste atenção na formatação: três strings diferentes foram criadas e atribuidas para três variaveis diferentes separadas por vírgula;
	//Mas detalhes sobre tipos de variáveis serão vistos em test com variaveis2;
	x, y, z := "Olá,", "Mundo!", "Pronto me livrei da maldição"
	//O comando usado para mostrar a variável na tela é fmt(package).Printf(identificador);
	//Esse comando precisa ser escrito exatamente desse modo(fmt.Printf) para poder funcionar corretamente;
	//Preste atenção na formatação: Com fmt.Printf("%v", x) mostramos o valor que está contido na variável x;
	//O \n em "%v %v\n" serve para pular uma linha;
	//Assim fmt.Printf("%v %v\n", x, y) mostra os valores atribuidos a x e y e pulará uma linha;
	fmt.Printf("%v %v\n", x, y)
	fmt.Printf(z)
}

//Aqui seu código termina!
