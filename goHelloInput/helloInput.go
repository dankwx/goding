package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Digite seu nome:")
	fmt.Scanln(&name)
	fmt.Printf("Olá %s, qual sua idade?", name)
	fmt.Scanln(&age)
	fmt.Printf("Ok, estas são suas informações:\n"+
		"Nome: %s\n"+
		"Idade:%d", name, age)
}
