package main

import "fmt"
/*https://www.tutorialkart.com/golang-tutorial/
//https://pt.b-ok.lat/    (libre)
https://pkg.go.dev/std   (Biblioteca Padrão)
*/
func main() {https://pkg.go.dev/std
	/*for i :=10;i>=3;i--{
		for j :=0;j<3-i-1;j++{
			fmt.Println(i)	
		}

	}
	var scores = []int{52,60,85,90,104,36}
	var lenght = len(scores)
	var tempArray = make([]int, lenght +1)
	
	tempArray = scores 
		for i :=0;i< lenght+1;i++{
			fmt.Println(i,"-",tempArray[i])
		} */
	
	var TwoDArray [8][8]int
	 TwoDArray[3][6] = 18
	 TwoDArray[7][4] = 3
	 fmt.Println(TwoDArray)
	
	/*
	var arr = [] int{5,6,7,8,9}
	fmt.Println("arr[]: ",len(arr),"-",cap(arr))
	var slic1 = arr[:]
	var sliceu = append(slic1, 12)
	fmt.Println("sliceu",sliceu)
	fmt.Println("sliceu: capacidade-> ",cap(sliceu))
	fmt.Println("sliceu: tamanho-> ",len(sliceu))
	fmt.Println("slic1: tamanho-> ",len(slic1))
	fmt.Println("slic1: capacidade->",cap(slic1))
	fmt.Println("slice1",slic1)
	
	var slic2 = arr[1:5]
	fmt.Println("slice2",slic2)
	var slic3 = append(slic2, 12)
	fmt.Println("slic2: tamanho-> ",len(slic2))
	fmt.Println("slic3: capacidade->",cap(slic2))
	fmt.Println("slic3: tamanho-> ",len(slic3))
	fmt.Println("slic3: capacidade-> ",cap(slic3))
	fmt.Println("slice3",slic3)*/
}

