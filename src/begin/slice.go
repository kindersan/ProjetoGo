/*package main
import ("fmt")

func main() {
  arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))
}*/

package main

import "fmt"
/*Padrões de fatia
Ao fatiar, você pode omitir os limites alto ou baixo para usar seus padrões.
O padrão é zero para o limite inferior e o comprimento da fatia para o limite superior.
Para a matriz
var a [10]int
essas expressões de fatia são equivalentes:

a[0:10]
um[:10]
a[0:]
uma[:]*/

/*package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a)
	fmt.Println(b)

	b[0] = "CARSON"
	fmt.Println(a, b)
	fmt.Println(names)
}*/


func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
} 

	//Slices can contain any type, including other slices.


/*package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "Mdi"
	board[2][2] = "O"
	board[1][2] = "Didi"
	board[1][0] = "O"
	board[0][2] = "carson"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
*/
