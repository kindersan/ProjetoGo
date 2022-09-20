package main
import ("fmt")

func sort(array []int, length int, ) []int{
	
	for i :=0;i < length-1;i++{
		var swap = false
		for j :=0;j<length-1;j++{
		
			if array[j] > array[j+1]{ //troca
				
				var flag = array[j]
				array[j] = array[j+1]
				array[j+1] = flag
				swap = true
			}
			
		}
		if swap == false{
			break
		}
	
	}
	return array
}
func main() {
	var scores = []int{90,100,120}
	
	var lenght = len(scores)
	sort(scores, lenght)
	
	for i :=0;i < lenght;i++{
		fmt.Printf("%d-%d\n",i, scores[i])
	}  
  
}
