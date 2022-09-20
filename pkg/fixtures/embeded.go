package fixtures

import (
	_ "embed"
)

//go:embed provided-example-input.csv
var ProvidedExampleInput []byte

//go:embed provided-example-output.csv
var ProvidedExampleOutput []byte
