/*
- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print
*/

package main

import (
	"fmt"
)

func main() {
	x, y, z := 42, "James Bond", true

	//(1)Demonstrar os valores nestas variáveis com uma única declaração print
	fmt.Println(x, y, z)
	//(2)Demonstrar os valores nestas variáveis com múltiplas declarações print
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T", z, z)
}
