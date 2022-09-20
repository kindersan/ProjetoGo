package main
import ("fmt")

func procurar(array []int, valor int, ) []int{
	var lenght =len(array)
	
	for i :=0;i< lenght;i++{
		if array[i] == valor{
			return i
		}	
	}
	return -1	
}
func main() {
	var scores = []int{52,60,85,90,104,36}
	var indice = procurar(scores, 75)
	
	if indice > 0{
		fmt.Printf("Valor encontrado:%d O indice é:%d ",valor,indice)
	}else{
		fmt.Printf("O valor não foi encontrado ",valor)
	} 
}
