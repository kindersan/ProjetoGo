package main

import "fmt"
//Função anônima acessando variável em cada iteração do loop dentro do corpo da função.

func main() {
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}
}
