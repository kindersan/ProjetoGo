package main
import ("fmt")

func remover(array[]int, indice int, ) []int{
	var lenght = len(array)
	var temparray = make([]int, lenght -1)
	
	for i :=0;i< lenght;i++{
		if i < indice{
			temparray[i] = array[i] 	
		}
		if i > indice {
			temparray[i-1] = array[i] 	
		}
	}
	fmt.Println("->",len(temparray))
	return temparray	
}
func main() {
	var scores = []int{52,60,85,90,104,36}
	scores = remover(scores, 3)
	var lenght = len(scores)
	
	for i :=0;i< lenght;i++{
		fmt.Printf("%d\n",scores[i])
	} 
  
}
