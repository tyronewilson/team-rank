package ask

import (
	"github.com/AlecAivazis/survey/v2"
	"spanchallenge/cmd/answer"
)

// OutputType is the type of output the user wants to get
func OutputType() (string, error) {
	var outputType string
	prompt := &survey.Select{
		Message: "What type of output do you want?",
		Options: answer.OutputDestinations,
	}
	err := survey.AskOne(prompt, &outputType)
	return outputType, err
}

func OutputFilename() (string, error) {
	var filename string
	prompt := &survey.Input{
		Message: "What is the filename you want to output to?",
	}
	err := survey.AskOne(prompt, &filename)
	return filename, err
}
