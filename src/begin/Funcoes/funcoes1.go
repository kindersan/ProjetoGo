package main

import (
	"fmt"
	"math/rand"
	"time"
	)
	// definimos o tipo "binFunc", que é uma função binária; leva dois
	// inteiros e retorna um inteiro. Isso não é estritamente necessário para
	// este exemplo, mas digitar `binFunc` repetidamente é muito mais
	// conveniente do que digitar `func(int, int) int` em todos os lugares que você vê.
	//digite binFunc func(int, int) int
type binFunc func(int, int) int

func main() {
	// propaga seu gerador de números aleatórios.
	// nota: isso não funciona no Go Playground, porque o tempo é
	// congelado no Go Playground. time.Now() sempre retorna o mesmo
	// valor.
	rand.Seed(time.Now().Unix())

	// cria uma fatia de funções
	fns := []binFunc{
		func(x, y int) int { 
			return x + y 
		},
		
		func(x, y int) int { 
			return x - y 
		},
		
		func(x, y int) int { 
			return x * y 
		},
		
		func(x, y int) int { 
			return x / y 
		},
		
		func(x, y int) int { 
			return x % y 
		},
	}

	// pick one of those functions at random
	f := fns[rand.Intn(len(fns))]
	fmt.Println(len(fns))
	x, y := 12, 5
	fmt.Println(f(x, y))
}

