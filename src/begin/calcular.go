package main
 
import (
    "fmt"
)
 
func f(s string) {
    panic(s)      // throws panic [lança pânico]
}
 
func main() {
        // defer makes the function run at the end [defer faz a função ser executada no final]
    defer func() {      // recovers panic
        if e := recover(); e != nil {
                    fmt.Println("Recovered from panic [Recuperado do pânico]")
            }
    }()
     
    f("1") // throws panic 
     
    // output:
    // Recovered from panic
}


/*package main


import (
    "fmt"
    "errors"
)

func main(){

	resultado,err := calcular(10,5)
	
	if err != nil {
		fmt.Println("Erro:",err)	
	}
	fmt.Println(resultado)
	
}


func calcular(a, b int) (int, error){
	if a == 0 {
		return 0, errors.New("Denominador não pode ser zero")
	}
	return a / b, nil
}


/*package main
 
import (
    "fmt"
    "errors"
)
 
func e(v int) (int, error) {
    if v == 0 {
        return 0, errors.New("Zero cannot be used")
    } else {
        return 2*v, nil
    }
}
 
func main() {
    v, err := e(0)
     
    if err != nil {
        fmt.Println(err, v)      // Zero cannot be used 0
    }   
}*/
