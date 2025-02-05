package main

import "fmt"

type State struct{
	transitionsFunction map[rune]*State
	finalState bool
}

func newState(finalState bool) *State{
	return &State{
		transitionsFunction: make(map[rune]*State),
		finalState: finalState,
	}
}

func (s *State) acceptCaracter(word string) bool{
	
	current := s;
	
	for _, c := range word{
		nextState, exists := current.transitionsFunction[c];

		if !exists{
			return false
		}
		
		current = nextState;
	}

	return current.finalState;
}

func main() {
	q0 := newState(false);
	q1 := newState(false);
	q2 := newState(false);
	q3 := newState(false);
	q4 := newState(false);
	q5 := newState(true);

	q0.transitionsFunction['c'] = q1;
	q1.transitionsFunction['a'] = q2;
	q1.transitionsFunction['o'] = q2;
	q2.transitionsFunction['r'] = q3;
	q3.transitionsFunction['r'] = q4;
	q4.transitionsFunction['o'] = q5;
	

	chains := []string{"carro","casa", "corro"};


	for _, word := range chains{
		accept := q0.acceptCaracter(word);
		if accept {
			fmt.Printf("Cadeia %s aceita\n", word);
		}else{
			fmt.Printf("Cadeia %s rejeitada\n", word);
		}
	}
}
