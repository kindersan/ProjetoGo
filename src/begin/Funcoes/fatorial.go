package main

import "fmt"

func main(){
	y := factorial(2)
	fmt.Println(y)

}
func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}
