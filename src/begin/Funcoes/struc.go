package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []string{"true", "true"}
	fmt.Println(r)

	s := []struct {
		i int
		b string
	}{
		{2, "true"},
		{3, "false"},
		
	}
	fmt.Println(s)
}

