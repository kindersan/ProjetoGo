/*package main

import "fmt"

func main() {
	//declarando um map
	var m = map[string]int{"Arroz": 1, "Feijão": 2}
	fmt.Println(m)

	//declarando map com make
	m1 := make(map[string]int)

	m1["Arroz"] = 1
	m1["Feijão"] = 2

	fmt.Println(m1)

	//verificar tamanho de um map
	fmt.Println(len(m1))
	

	//deletando elemento de um map
	delete(m1, "Feijão")
	fmt.Println(m1)
}*/

/*package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}*/


package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
	"Carson": Vertex{
		97121792, -00145603,
	},
}

func main() {
	fmt.Println(m)
	v, ok := m["carson"]
	fmt.Println("The value:", v, "Present?", ok)
}

