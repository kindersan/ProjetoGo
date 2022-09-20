package main

import "fmt"

type Author struct {
  firstName, lastName, Biography string
  photoId    int
}

type blogPost struct {
  author  Author // campo de estrutura aninhada
  title   string
  postId  int 
  published  bool  
}

func main() {
        b := new(blogPost)

        fmt.Println(b)

        b.author.firstName= "Alex"
        b.author.lastName= "Nnakwue"
        b.author.Biography = "I am a lazy engineer"
        b.author.photoId = 234333
        b.published=true
        b.title= "understand interface and struct type in Go"
        b.postId= 12345

        fmt.Println(*b)        

}

