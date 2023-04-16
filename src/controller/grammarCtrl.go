package controller

import (
	"fmt"

	"github.com/grijalbaEmilio/ll1/src/model"
)

func First(grammar *model.Grammar) (map[string][]string, error) {
	firsts := map[string][]string{}
	for _, symbol := range grammar.NonTerminals {
		first, err := grammar.First(symbol)

		if err != nil {
			return nil, fmt.Errorf("no se pueden encontrar primeros -> %v", err)
		}

		firsts[symbol] = append(firsts[symbol], first...)
	}

	return firsts, nil
}

func Follow(grammar *model.Grammar) (map[string][]string, error) {
	follows := map[string][]string{}
	for _, symbol := range grammar.NonTerminals {
		follow, err := grammar.Follow(symbol)
		if err != nil {
			return nil, fmt.Errorf("no se puede encontrar siguintes -> %v", err)
		}
		follows[symbol] = append(follows[symbol], follow...)
	}

	return follows, nil
}

func Predictions(grammar *model.Grammar) (map[string][][]string, error) {
	predicitions := map[string][][]string{}
	for _, noTeminal := range grammar.NonTerminals {
		predicition, _ := grammar.Predictions(noTeminal)
		predicitions[noTeminal] = predicition
	}

	return predicitions, nil
}
