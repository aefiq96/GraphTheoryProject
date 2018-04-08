package main 

import(
	"bufio"
	"os"
	"strings"
	"fmt"
	nfa "./src"
)
//intopost function
func intopost(infix string)string{
	specials := map[rune]int{'*': 10, '.': 9, '|': 8, '+':9, '?':9}

	pofix := []rune{}
	s := []rune{}
	//for loop with switch cases to append
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
		//if it doesn't find the needs then you will get the default
		default:
			pofix = append(pofix, r)
		}
	}

	for len(s) > 0{
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}
	//simple returning string
	return string (pofix)
}
// function to read input from users
func readInput() (string, error){
	reader := bufio.NewReader(os.Stdin)
	str,err := reader.ReadString('\n')

	return strings.TrimSpace(str),err
}
//main method
func main(){
	//declaring variables
	var i int 
	//list of options for user to enter
	fmt.Print("Please Select 1\n")
	fmt.Print("select 1 for Infix To PostFix\n")
	fmt.Print("select 2 for PostFix\n")

	fmt.Scanln(&i)
	//switch case for the menu
	switch i{
	case 1:
		fmt.Print("Enter infix: ");
		readInfix, err := readInput()

		if err != nil{
			return
		}
		fmt.Println("Infix: ",readInfix)
		fmt.Println("Postfix: ",intopost(readInfix))
		 
		post := intopost(readInfix)

		fmt.Print("Enter the string you would like to match: ");
		readStr, err := readInput()

		fmt.Println(nfa.Pomatch(post,readStr))	

	case 2:
		fmt.Print("Enter postfix: ");
		readInfix, err := readInput()

		if err != nil{
			return
		}
				
		fmt.Print("Enter the string you would like to match: ");
		readStr, err := readInput()

		fmt.Println(nfa.Pomatch(readInfix,readStr))	
	}

}

