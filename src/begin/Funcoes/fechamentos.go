package main

import "fmt"
//Fechamentos são um caso especial de funções anônimas.
//Fechamentos são funções anônimas que acessam as variáveis definidas fora do corpo da função.


func main() {
	l := 20
	b := 30

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
}
