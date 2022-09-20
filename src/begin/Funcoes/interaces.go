package main

import "fmt"

// interface definition
type Publisher interface {
     Publish()  error
}

type blogPost struct {
  author  string
  title   string
  postId  int  
}

// method with a value receiver
func (b blogPost) Publish() error {
   fmt. Printf("The title on %s has been published by %s, with postId %d\n" , b.title, b.author, b.postId)
   return nil
}

 func test(){

  b := blogPost{"Alex","understanding structs and interface types",12345}

  fmt.Println(b.Publish())

   d := &b   // pointer receiver for the struct type

   b.author = "Chinedu"


   fmt.Println(d.Publish())

}


func main() {

        var p Publisher

        fmt.Println("->",p)

        p = blogPost{"Alex","understanding structs and interface types",12345}

        fmt.Println("-#",p.Publish())

        test()  // call the test function 

}
