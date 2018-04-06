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

}

func main(){
    nfa := poregtonfa("ab.c*|")
    fmt.Println(nfa)
}

