	package main
	import ("fmt")

	func insert(array []int, lenght int, temparray []int, valor int, insertIndex int ) []int{
		
		for i :=0;i< lenght;i++{
			if i < insertIndex {
				temparray[i] = array[i]
			}else{
				temparray[i+1] = array[i]
			}
		
		}
		
		temparray[insertIndex] = valor
		return temparray
	}
	
	func main() {
		var scores = []int{52,60,85,90,104,36}
		var lenght = len(scores)
		var tempArray = make([]int, lenght +1)
		
		insert(scores, lenght, tempArray, 1100, 0)
		scores = tempArray
		for i :=0;i< lenght+1;i++{
			fmt.Printf("%d\n",scores[i])
		} 
	  
	}
