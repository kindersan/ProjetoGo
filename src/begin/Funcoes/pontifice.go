package main //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

import "fmt"


/*func zero(xPtr *int) {
  *xPtr = 100
}
func main() {
  x := 5
  zero(&x)
  fmt.Println(x) // x is 0
}*/

func one(xPtr *int) {
  *xPtr = 1
}
func main() {
  xPtr := new(int)
  one(xPtr)
  fmt.Println(*xPtr) // x is 1
}
