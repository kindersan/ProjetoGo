package main

import "fmt"

var i, j int = 1, 2

func main() {
var k int = 42
var f float64 = float64(k)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Printf("tipo:%T \n->valor: %v \n",f,f)
	fmt.Println(f)
}

