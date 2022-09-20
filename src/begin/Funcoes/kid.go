package main

import "fmt"


	/*arr := [5]float64{1,2,3,4,5}
	x := arr[1:4]
	fmt.Println("cascading panics",x)
	fmt.Println("cascading Madi XXL tamanho",len(x))
	fmt.Println("cascading capacidade",cap(x))
	/*slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	slice1 := []int{1,2,3}
  	slice2 := make([]int, 2)
  	copy(slice2, slice1)
  	fmt.Println(slice1, slice2)


  elements := make(map[string]string)
  elements["H"] = "Hydrogen"
  elements["He"] = "Helium"
  elements["Li"] = "Lithium"
  elements["Be"] = "Beryllium"
  elements["B"] = "Boron"
  elements["C"] = "Carbon"
  elements["N"] = "Nitrogen"
  elements["O"] = "Oxygen"
  elements["F"] = "Fluorine"
  elements["Ne"] = "Neon"
  elements["GI"] = "Carson"

  fmt.Println(elements["GI"])
  fmt.Println(len(elements))
  
	/*if name, ok := elements["GI"]; ok {
	fmt.Println(name, ok)
	}
	name, val := elements["Ne"]
        fmt.Println(name, val)
        ####################
func main() {        
	fmt.Println(f1())
	x, y := f()
	fmt.Println(x, y)
}
	func f1() int {
	  return f2()
	}
	func f2() int {
	  return 1
	}

	func f() (int, int) {
	  return 5, 6
	}
	####################

	func main() {
	  add := func(x, y int) int {
	    return x + y
	  }
	  fmt.Println(add(3,4))
	}
	#########################

	#########################
func main() {
  x := 0
  increment := func() int {
    x++
    return x
  }
  fmt.Println(increment())
  fmt.Println(increment())
} 
	######################### 
*/

func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
func main() {
  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4
}

