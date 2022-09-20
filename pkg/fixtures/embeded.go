package fixtures

import (
	// we use embed here to provide ag hoc fixtures for testing and don't want to have to remember to do it in
	// every other test package
	_ "embed"
)

// ProvidedExampleInput is an example input for the provided example.
//go:embed provided-example-input.csv
var ProvidedExampleInput []byte

// ProvidedExampleOutput is the expected output for the provided example input
//go:embed provided-example-output.csv
var ProvidedExampleOutput []byte
