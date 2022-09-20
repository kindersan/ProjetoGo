package main

import "fmt"
	// Function closures
/*As funções Go podem ser encerramentos. Um encerramento é um valor de função que referencia variáveis de fora de seu corpo.
A função pode acessar e atribuir as variáveis referenciadas; 
neste sentido a função está "ligada" às variáveis.

Por exemplo, a função somador retorna um encerramento. Cada encerramento está vinculado à sua própria variável de soma.*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
