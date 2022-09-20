package main
import ("fmt")

func main() {
	var scores = []int{52,60,85,90,104,36}
	slice := scores[3:]
	fmt.Println(scores)
	fmt.Println("slice: capacidade-> ",cap(slice))
	fmt.Println("slice: tamanho-> ",len(slice))
	fmt.Println(slice)
	ponteiro := append(slice,10,11,12,13,14,15,16,17)
	
  	fmt.Println("Ponteiro-> ",ponteiro)
	fmt.Println("Ponteiro: capacidade-> ",cap(ponteiro))
	fmt.Println("Ponteiro: tamanho-> ",len(ponteiro))
	
 
}
