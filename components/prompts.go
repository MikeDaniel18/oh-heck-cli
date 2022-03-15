package components

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/manifoldco/promptui"
)

func YesNoInput(question string) bool {
	result := false
	prompt := &survey.Confirm{
		Message: question,
	}

	err := survey.AskOne(prompt, &result)

	surveyExitError(err)

	return result
}

func StringInput(text string, initial string) string {

	if len(initial) > 0 {
		// FIXME: Falls on its face when spanning multiple lines. Needs a fix
		prompt := promptui.Prompt{
			Label:     text,
			Default:   initial,
			AllowEdit: true,
		}

		result, _ := prompt.Run()
		return result
	} else {
		answer := ""

		// FIXME: Currently doesn't have a way of setting a default input string
		prompt := &survey.Input{
			Message: text,
		}
		err := survey.AskOne(prompt, &answer, nil)
		surveyExitError(err)

		return answer
	}
}

func surveyExitError(err error) {
	if err != nil { // NOTE: This is a fix for ctrl + c not working with Survey
		if err == terminal.InterruptErr {
			log.Fatal("Oh-heck exited")
		}
	}
}
