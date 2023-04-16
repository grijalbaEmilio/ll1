# ll1

This is a module for find `firsts` `follows` and `set of predictions`

## how use?

1. import `model` and `controller`

```go
import (
	"github.com/grijalbaEmilio/ll1/src/controller"
	"github.com/grijalbaEmilio/ll1/src/model"
)
```

2. instance a grammar

```go
	grammar := model.Grammar{
		Productions: map[string][]string{
			"S": {"T E n"},
			"E": {"+ T E", "- T E", "λ"},
			"T": {"F Y"},
			"Y": {"* F Y", "/ F Y", "λ"},
			"F": {"( S )", "num", "id"},
		},
		NonTerminals: []string{"S", "E", "T", "Y", "F"},
		Terminals:    []string{"*", "+", "(", ")", "num"},
		StartSymbol:  "S",
	}
```

2. use functions `controller.First`, `controller.Follow` or `controller.Predictions`

```go
	first, err := controller.First(&grammar)
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, value := range first {
		fmt.Printf("First(%s) = %s\n", key, value)
	} 
```

