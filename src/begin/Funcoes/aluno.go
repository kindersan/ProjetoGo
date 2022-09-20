package main

import "fmt"

type Aluno struct{
	nome string
	idade uint
	sexo string
}
type Endereco struct{
	aluno Aluno
	rua string
	cidade string
	numero uint
	published  bool
}

func main(){
	a := new(Endereco)
	fmt.Println(a)
	
	a.aluno.nome = "Carson"
	a.aluno.idade = 41
	a.aluno.sexo = "M"
	a.rua = "Borges de Medeiros"
	a.cidade = "Alegrete"
	a.numero = 546
	fmt.Println(*a)
}

/*package main

import "fmt"

type Author struct {
  firstName, lastName, Biography string
  photoId    int
}

type BlogPost struct {
  Author  // passando diretamente a estrutura Author como um campo - também chamado de campo anônimo ou tipo incorporado
  title   string
  postId  int 
  published  bool  
}

func main() {
        b := BlogPost{
        Author: Author{"Alex", "Nnakwue", "I am a lazy engineer", 234333},
        title:"understand interface and struct type in Go",
        published:true,
        postId: 12345,
        }

        fmt.Println(b.firstName) // remember the firstName field is present on the Author struct?
        fmt.Println(b)        

}*/
