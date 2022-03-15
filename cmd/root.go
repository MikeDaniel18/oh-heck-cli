/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"oh-heck/components"
	"oh-heck/configs"
	"oh-heck/models"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	// "golang.design/x/clipboard"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oh-heck",
	Short: "Helping you to write terminal commands",
	Long: `Oh-heck is an AI tool that helps you to write terminal
commands you've forgotten. It's as simple as writing:

oh-heck "how do I remove all .yaml files from this directory?"

An AI generated response will then be returned which you
can execute in the command line`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flags().Changed("set-testing") {
			result, err := cmd.Flags().GetBool("set-testing")

			isTesting := result && err == nil
			// Set the testing config
			cfg := configs.ReadConfig()
			cfg.IsTesting = isTesting
			cfg.Save()
			return
		} else {
			askQuestion(args, nil)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oh-heck.yaml)")
	rootCmd.Flags().BoolP("set-testing", "t", false, "")
	rootCmd.Flags().MarkHidden("set-testing")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func askQuestion(args []string, placeholder *string) {
	// Check for API Key
	var cfg = configs.ReadConfig()

	if cfg.ApiKey == "" {
		fmt.Println("No API key stored.")
		ShowApiOptions()
		cfg = configs.ReadConfig()
		if cfg.ApiKey == "" {
			return
		}
	}
	// Check for args and proceed
	var question = ""

	var defaultValue = ""
	if placeholder != nil {
		defaultValue = *placeholder
	}

	if len(args) > 0 && len(args[0]) > 3 && len(defaultValue) == 0 {
		question = args[0]
	} else {
		var label = "Question:"
		if len(defaultValue) > 0 {
			label = "Try rewording:"
		}

		result := components.StringInput(label, defaultValue)
		question = result
	}

	if len(question) > 3 {
		makeBashQuestionCall(question)
	}
}

func makeBashQuestionCall(question string) {
	// Ask the question
	completion, errResp := models.AskBashQuestionApi(question)

	if errResp != nil && (errResp.Status == 403 || errResp.Status == 401) {
		InvalidApiKey()
	} else if errResp != nil && errResp.Status > 0 {
		fmt.Println(errResp.ErrorMessage)
	} else if completion != nil && len(completion.Id) > 0 {
		output := strings.TrimSpace(completion.Response)

		promptQuestion := fmt.Sprintf("Output: $ %v?", output)

		accepted := components.YesNoInput(promptQuestion)
		if accepted {
			// Train success
			go models.SetQuestionResponse(*completion, true)
			copyToClipboard(output)
		} else {
			// Train failure
			// go models.SetQuestionResponse(*completion, false) // Default is false so no need to set

			askQuestion(nil, &question)
		}
	} else {
		fmt.Println("Something went wrong")
	}
}

func copyToClipboard(text string) {
	// err := clipboard.Init()
	// if err != nil {
	// 	fmt.Println("Couldn't copy to clipboard")
	// 	return
	// }

	// clipboard.Write(clipboard.FmtText, []byte(text))
	// fmt.Println("Copied to clipboard")

	err := clipboard.WriteAll(text)

	if err != nil {
		fmt.Println("Could not copy to clipboard. On linux make sure either 'xclip' or 'xsel' are installed")
	} else {
		fmt.Println("Copied to clipboard")
	}
}
