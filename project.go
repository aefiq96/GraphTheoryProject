package main 

import(
	"fmt"
)

func intopost(infix string)string{
	specials := map[rune]int{'*': 10, '.':9, '|':8}

	pofix := []rune{}
	s := []rune{}

	return string (pofix)
}

func main(){
	//Answer: ab.c*
	fmt.Println("Infix: ","a.b.c*");
	fmt.Println("Postifx: ",intopost("a.b.c*"))

	//
}