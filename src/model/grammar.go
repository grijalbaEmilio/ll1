package model

import (
	"fmt"
	"strings"

	aux "github.com/grijalbaEmilio/ll1/src/helpers"
)

type Grammar struct {
	Productions  map[string][]string
	NonTerminals []string
	Terminals    []string
	StartSymbol  string
	follows      map[string][]string
}

func (g *Grammar) GetFollow(symbol string) []string {
	return g.follows[symbol]
}

// producciones de un sólo no terminal
func (g *Grammar) First(symbol string) ([]string, error) {

	if !aux.Contains(g.NonTerminals, symbol) {
		return nil, fmt.Errorf("el símbolo \"" + symbol + "\" no pertenece a los NO terminales!")
	}

	productions := g.Productions[symbol]
	firsts := []string{}
	for _, production := range productions {
		symbolsSplit := strings.Split(production, " ")
		y1 := symbolsSplit[0]

		// evitamos recursión infinita
		if y1 == symbol {
			return nil, fmt.Errorf("la gramática tienerecursión por izquierda en la producción %s", symbol)
		}

		// Si y1 es un terminal, entonces agregamos y1 a prim(x)
		if aux.Contains(g.Terminals, y1) {
			firsts = append(firsts, y1)
			continue
		}

		// Si y1 es un no terminal, entonces agregamos prin(y1) a prim(x)
		if aux.Contains(g.NonTerminals, y1) {
			recursiveY1, _ := g.First(y1)
			firsts = append(firsts, recursiveY1...)
			continue
		}

		// Si y1 es λ, entonces agregamos prim(y2) a prim(x)
		if y1 == "λ" && len(symbolsSplit) > 1 {
			y2 := symbolsSplit[1]
			if aux.Contains(g.Terminals, y2) {
				firsts = append(firsts, y2)
			} else {
				recursiveY2, _ := g.First(y2)
				firsts = append(firsts, recursiveY2...)
			}
		}

		// Si y1 es λ y no existe y2, entonces agregamos λ a prim(x)
		if y1 == "λ" && len(symbolsSplit) == 1 {
			firsts = append(firsts, "λ")
		}
	}

	// Eliminamos los primeros repetidos y retornamos el resultado
	return aux.Unique(firsts), nil
}

func (g *Grammar) Follow(symbol string) ([]string, error) {

	if !aux.Contains(g.NonTerminals, symbol) {
		return nil, fmt.Errorf("el símbolo \"" + symbol + "\" no pertenece a los NO terminales!")
	}

	// Inicializamos los siguientes
	if g.follows == nil {
		g.follows = map[string][]string{}
		for _, nonTerminal := range g.NonTerminals {
			if nonTerminal == g.StartSymbol {
				g.follows[nonTerminal] = []string{"$"}
				continue
			}

			g.follows[nonTerminal] = []string{}
		}
	}

	// Recorremos las producciones de la gramática
	for keyNonTerminal, productionList := range g.Productions {
		// se recorre la separación por |
		for _, production := range productionList {
			symbols := strings.Split(production, " ")
			for i, symbol := range symbols {
				if !aux.Contains(g.NonTerminals, symbol) {
					continue
				}

				// si existe siguiente
				if i+1 < len(symbols) {
					beta := symbols[i+1]

					// si el siguiente es unterminal se agrega siguientes
					if aux.Contains(g.Terminals, beta) {
						g.follows[symbol] = append(g.follows[symbol], beta)
					} else if aux.Contains(g.NonTerminals, beta) {
						betaFirsts, err := g.First(beta)

						if err != nil {
							return nil, fmt.Errorf("no se pueden hallar siguientes -> %v", err)
						}

						if aux.Contains(betaFirsts, "λ") {
							betaFirstsNoLamnda := aux.RemoveElement(betaFirsts, "λ")
							g.follows[symbol] = append(g.follows[symbol], g.follows[keyNonTerminal]...)
							g.follows[symbol] = append(g.follows[symbol], betaFirstsNoLamnda...)
						}
					}

				} else {
					// si siguiente es λ
					g.follows[symbol] = append(g.follows[symbol], g.follows[keyNonTerminal]...)
				}

			}
		}
	}

	// Eliminamos los siguientes repetidos y retornamos el resultado
	for nonTerminal, followList := range g.follows {
		g.follows[nonTerminal] = aux.Unique(followList)
	}

	return g.follows[symbol], nil
}

func (g *Grammar) Predictions(symbol string) ([][]string, error) {
	if !aux.Contains(g.NonTerminals, symbol) {
		return nil, fmt.Errorf("el símbolo \"" + symbol + "\" no pertenece a los NO terminales!")
	}

	// predictions := []string{}
	firsts, _ := g.FirstForPrediction(symbol)

	for i, first := range firsts {
		if aux.Contains(first, "λ") {

			firsts[i], _ = g.Follow(symbol)
		}
	}

	return firsts, nil
}

func (g *Grammar) FirstForPrediction(symbol string) ([][]string, error) {

	if !aux.Contains(g.NonTerminals, symbol) {
		return nil, fmt.Errorf("el símbolo \"" + symbol + "\" no pertenece a los NO terminales!")
	}

	productions := g.Productions[symbol]
	firsts := [][]string{}
	for _, production := range productions {
		symbolsSplit := strings.Split(production, " ")
		y1 := symbolsSplit[0]

		// evitamos recursión infinita
		if y1 == symbol {
			continue
		}

		// Si y1 es un terminal, entonces agregamos y1 a prim(x)
		if aux.Contains(g.Terminals, y1) {
			firsts = append(firsts, []string{y1})
			continue
		}

		// Si y1 es un no terminal, entonces agregamos prin(y1) a prim(x)
		if aux.Contains(g.NonTerminals, y1) {
			recursiveY1, _ := g.FirstForPrediction(y1)
			firsts = append(firsts, recursiveY1...)
			continue
		}

		// Si y1 es λ, entonces agregamos prin(beta) a prim(x)
		if y1 == "λ" && len(symbolsSplit) > 1 {
			beta := symbolsSplit[1]
			if aux.Contains(g.Terminals, beta) {
				firsts = append(firsts, []string{beta})
			} else {
				recursivebeta, _ := g.FirstForPrediction(beta)
				firsts = append(firsts, recursivebeta...)
			}
		}

		// Si y1 hasta yk tiene λ, entonces agregamos λ a prim(x)
		if y1 == "λ" && len(symbolsSplit) == 1 {
			firsts = append(firsts, []string{"λ"})
		}
	}

	// Eliminamos los primeros repetidos y retornamos el resultado
	return firsts, nil
}

func (g *Grammar) CheckGrammar() error {
	for _, productionList := range g.Productions {
		for _, production := range productionList {
			symbolsSplit := strings.Split(production, " ")
			for _, prod := range symbolsSplit {
				if !aux.Contains(g.Terminals, prod) && !aux.Contains(g.NonTerminals, prod) && prod != "λ" {
					return fmt.Errorf("%s no pertenece a la gramática !", prod)
				}
			}
		}
	}

	return nil
}

/*

func validateProductions(prods map[string][]string, grammar model.Grammar) bool {
	// Iteramos sobre todas las producciones definidas en el mapa
	for nonTerminal, productions := range prods {
		// Verificamos que el no terminal pertenezca a la gramática
		if !contains(grammar.NonTerminals, nonTerminal) {
			return false
		}

		// Iteramos sobre cada producción del no terminal
		for _, production := range productions {
			// Separamos los símbolos de la producción por espacio
			symbolsSplit := strings.Split(production, " ")

			// Iteramos sobre cada símbolo de la producción
			for _, symbol := range symbolsSplit {
				// Verificamos que el símbolo pertenezca a la gramática
				if !contains(grammar.NonTerminals, symbol) && !contains(grammar.Terminals, symbol) {
					return false
				}
			}
		}
	}

	return true
}

// Función auxiliar que retorna true si un elemento está presente en un arreglo de strings
func contains(arr []string, elem string) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}
*/
