package main

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

func poregtonfa(pofix string) *nfa { 
    nfastack := []*nfa{}   

    for _, r := range pofix {
        switch r {
        case '.':          
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]           

            frag1.accept.edge1 = frag2.initial

            nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

        case '|':   
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            accept := state{}
            initial := state{edge1: frag1.initial, edge2: frag2.initial}
            frag1.accept.edge1 = &accept
            frag2.accept.edge1 = &accept


            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})         
        case '*':            
        default:            
        }
        }
    return nfastack[0]
}

func main(){
    nfa := poregtonfa("ab.c*|")
    fmt.Println(nfa)
}

