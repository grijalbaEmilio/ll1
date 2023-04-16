package main

import (
	"fmt"

	"github.com/grijalbaEmilio/ll1/src/controller"
	"github.com/grijalbaEmilio/ll1/src/model"
)

func main() {
	// Creamos una gramática de ejemplo

	/* prods := map[string][]string{
		"S": {"T E n"},
		"E": {"+ T E", "- T E", "λ"},
		"T": {"F Y"},
		"Y": {"* F Y", "/ F Y", "λ"},
		"F": {"( S )", "num", "id"},
	} */

	prods := map[string][]string{
		"S": {"T E"},
		"E": {"+ T E", "λ"},
		"T": {"F Y"},
		"Y": {"* F Y", "λ"},
		"F": {"( S )", "num"},
	}

	grammar := model.Grammar{
		Productions:  prods,
		NonTerminals: []string{"S", "E", "T", "Y", "F"},
		Terminals:    []string{"*", "+", "(", ")", "num"},
		StartSymbol:  "S",
	}

	err := grammar.CheckGrammar()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Obtenemos los primeros de cada símbolo no terminal
	first, err := controller.Firsts(&grammar)
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, value := range first {
		fmt.Printf("First(%s) = %s\n", key, value)
	}
	fmt.Println()

	// Obtenemos los siguientes de cada símbolo no terminal
	follows, err := controller.Follows(&grammar)

	if err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range follows {
		fmt.Printf("Follow(%s) = %s\n", key, value)
	}
	fmt.Println()

	// conjunto de prediccion
	predicitions, err := controller.Predictions(&grammar)

	if err != nil {
		fmt.Println(err)
		return
	}
	for key, value := range predicitions {

		fmt.Println("prediction de ", key, ": ", value)
	}

}
