package main
import ("fmt")

func main() {
	
	x := 5
	fn := func() {
	    fmt.Println("x is", x)
	}
	
	fn()
	x++
	fn()
}

