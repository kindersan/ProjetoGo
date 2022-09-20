/*package main

import "fmt"
// Em Go, apenas constantes são imutáveis. No entanto, como os argumentos são passados por valor,
// uma função que recebe um argumento de valor e o muta, não mudará o valor original.


type Artist struct {
	Name, Genero string
	Musicas       int
}

func novolancamento(a Artist) int {
	a.Musicas++
	return a.Musicas
}

func main() {
	me := Artist{Name: "Matt", Genero: "Electro", Musicas: 42}
	fmt.Printf("%s released their %dth musicas\n", me.Name, novolancamento(me))
	fmt.Printf("%s has a total of %d Musicas\n", me.Name, me.Musicas)
}*/

//Como você pode ver, a quantidade total de músicas no valor da variável não foi alterada. Para mutar o valor passado, 
//precisamos passá-lo por referência, usando um ponteiro.me

package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}


