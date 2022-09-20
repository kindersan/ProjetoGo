package main
import ("fmt")

func main() {
  arr1 := [4]string{"Rafael", "BMW", "kid", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1),arr1)
  fmt.Println(len(arr2),arr2)
}
