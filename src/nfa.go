package nfa

import (
    "fmt"
)
type state struct{
    symbol rune
    edge1 *state
    edge2 *state
}

type nfa struct{
    initial *state
    accept  *state
}
//postfix regular expression function
func poregtonfa(pofix string) *nfa { 
//creating a stack
nfaStack := []*nfa{}
	for _, r := range pofix {
		switch r {
		case '.':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}

			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '*':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) 
         case '+':
            frag := nfaStack[len(nfaStack)-1]
            nfaStack = nfaStack[:len(nfaStack)-1]

            accept := state{}
            initial := state{edge1: frag.initial, edge2: &accept}

            frag.accept.edge1 = &initial
            frag.accept.edge2 = &accept

            nfaStack = append(nfaStack, &nfa{initial: frag.initial, accept: &accept}) 
        case '?':
            frag := nfaStack[len(nfaStack)-1]
            nfaStack = nfaStack[:len(nfaStack)-1]

            initial := state{edge1: frag.initial, edge2: frag.accept}

            nfaStack = append(nfaStack, &nfa{initial: &initial, accept: frag.accept}) 
        default:    
            accept := state{}
            initial := state{symbol: r,edge1: &accept}

            nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})        
		}
	}
    //if there is an error then this println will show up
	if len(nfaStack) != 1 {
		fmt.Println("error", len(nfaStack), nfaStack)
	}
	return nfaStack[0]
}
//add state function
func addState (l []*state, s *state, a *state) []*state{
    l = append(l, s)

    if s != a && s.symbol == 0 {
        l = addState(l, s.edge1, a)
        //other programming language it would be null but in GO it's nil
        if s.edge2 != nil{
            l = addState(l, s.edge2, a)
        }
    }
    return l
}
//postfix match function to match the strings
func Pomatch (po string, s string)bool{
    ismatch := false
    ponfa := poregtonfa(po)

    current := []*state{}
    next := []*state{}

    current = addState(current[:], ponfa.initial, ponfa.accept)
//for loop within a for loop
    for _, r := range s{
        for _, c := range current{
            if c.symbol == r{
                next = addState(next[:], c.edge1, ponfa.accept)
            }
        }
        current, next = next, []*state{}
    }
//for loop to see if the boolean is true
    for _, c := range current{
        if c == ponfa.accept{
            ismatch = true
            break
        }
    }
    //returning the ismatch boolean
    return ismatch
}



