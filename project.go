package main 

import(
	"fmt"
)

func intopost(infix string)string{
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	pofix := []rune{}
	s := []rune{}

	for _, r := range infix{
		switch{
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '('{
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials [r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]]{
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)
		}
	}

	for len(s) > 0{
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string (pofix)
}

func main(){
	//Answer: ab.c*
	fmt.Println("Infix: ","a.b.c*");
	fmt.Println("Postifx: ",intopost("a.b.c*"))

	//
}