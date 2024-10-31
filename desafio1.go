// CURSO DE GO DA DIO
// "Aprendendo Estruturas em Go"

// criar um código que exibe todos os números entre 1 e 100 que são divisíveis por 3
// sempre que aparecer um multiplo de 3 printar pin
// sempre que aparecer multiplo de 5 printar pan
// sem printar os números, apenas os cõdigos pin e pan.
// de 1 a 100

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 00 && i%5 == 0 {
			fmt.Println("PinPan!")
		}
		if i%3 == 0 {
			fmt.Println("Pin!")
		} else if i%5 == 0 {
			fmt.Println("Pan!")
		} else {
			fmt.Println(i)
		}
	}
}
