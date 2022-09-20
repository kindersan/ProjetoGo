package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {
        b := new(blogPost)

        fmt.Println(b) // zero value

        b.author= "Alex"
        b.title= "understand interface and struct type in Go"
        b.postId= 12345

        fmt.Println(*b)   // desreferenciar o ponteiro    

}

/*package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {
        var b blogPost // ou := {"Alex", "understand struct and interface type", 12345}
        b.author= "Alex"
        b.title="understand structs and interface types"
        b.postId=12345

        fmt.Println(b)  

        b.author = "Chinedu"  // já que tudo é passado por valor por padrão em Go, podemos atualizar este campo após inicializar - veja os tipos de ponteiro mais tarde

        fmt.Println("Updated Author's name is: ", b.author)           
}*/
