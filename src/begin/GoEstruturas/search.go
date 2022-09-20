package main
import ("fmt")

func main() {
	var scores = []int{52,60,85,90,104,36}
	var length = len(scores)
	var procurarValor = 36
	var posicao = procuraLinear(scores, length, procurarValor)
	fmt.Printf("%dPosicao:  %d\n",procurarValor, posicao)
	fmt.Println("-------------------------")
}

func  procuraLinear(array []int, length int, searchValor int)int{
	var low =0       // indice mais baixos
	var high= length // indice mais altos
	var mind=0       //indice do meio
	for {
		if low >= high {
			break
		}
		mind = (low + high)/2
		if array[mind] == searchValor {
			return mind
		}else if array[mind] < searchValor {
			low = mind +1
		}else if array[mind] > searchValor {
			high = mind -1
	        }
	 }
	return -1
} 
