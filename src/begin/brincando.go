package main
import ("fmt"
	//"math/rand"
)

func remover(array []string, indice int ) []string{
	var lenght = len(array)
	var temparray = make([]string, lenght -1)
	
	for i :=0;i< lenght;i++{
		if i < indice{
			temparray[i] = array[i] 	
		}
		if i > indice {
			temparray[i-1] = array[i] 	
		}
	}
	//fmt.Println("->",len(temparray))
	//temparray[lenght] = valor
	return temparray	
}
func main() {
	var scores = []string{"Dino","Didi","Madibu","kidi","leonardo","chico linguiça"}
	
	scores = remover(scores, 3)
	var lenght = len(scores)
	
	for i :=0;i< lenght;i++{
		fmt.Println("Boy -> ",i+1,"=", scores[i])//,"<-->",rand.Intn(10))
	} 
  
}
