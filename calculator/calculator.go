package main

import (
	"fmt"
)

func sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mult(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func main() {
	var num1, num2 int
	var escolha int

	for {
		fmt.Println("Que operação deseja realizar")
		fmt.Scanln(&escolha)

		if escolha == 0 {
			break
		}

		fmt.Println("Escolha o primeiro número:")
		fmt.Scanln(&num1)
		fmt.Println("Escolha o segundo número:")
		fmt.Scanln(&num2)
		operation := escolha

		switch operation {
		case 1:
			fmt.Println(sum(num1, num2))
		case 2:
			fmt.Println(sub(num1, num2))
		case 3:
			fmt.Println(mult(num1, num2))
		case 4:
			fmt.Println(div(num1, num2))

		}

	}

}
