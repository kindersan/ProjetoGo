package main
import ("fmt")

func append(array []int, valor int, ) []int{
	var lenght = len(array)
	var temparray = make([]int, lenght +1)
	
	for i :=0;i< lenght;i++{
		temparray[i] = array[i] 	
	}
	
	temparray[lenght] = valor
	
	return temparray	
}
func main() {
	var scores = []int{52,60,85,90,104,36}
	scores = append(scores, 75)
	var lenght = len(scores)
	
	for i :=0;i< lenght;i++{
		fmt.Printf("%d\n",scores[i])
	} 
  
}
