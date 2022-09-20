package main
import ("fmt"
	"math/rand"
)

func main() {
	
	x := 5
	fn := func() {
	    fmt.Println("x is", x)
	}
	
	fn()
	x++
	fn()
	fmt.Println("My favorite number is", rand.Intn(10))
}


