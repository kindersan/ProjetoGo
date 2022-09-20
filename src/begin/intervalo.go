package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for i, j := range fruits {
     fmt.Printf("%v\t%v\n", i, j)
  }
}

/*	Aqui, queremos omitir os valores ( idx armazena o índice, valarmazena o valor):
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}

  for idx, _ := range fruits {
     fmt.Printf("%v\n", idx)
  }
}
*/
