/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"oh-heck/models"
	"os/exec"
	"runtime"

	"oh-heck/components"
	"oh-heck/configs"

	"github.com/spf13/cobra"
)

// apikeyCmd represents the apikey command
var apikeyCmd = &cobra.Command{
	Use:   "apikey",
	Short: "Use to set your API Key",
	Long: `Use to set your API key and store it for future API Calls.
	You can optionally pass through an API key to force set a new one`,
	Run: func(cmd *cobra.Command, args []string) {
		configureApiKeys(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(apikeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apikeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	apikeyCmd.Flags().StringP("key", "k", "", "Pass through a new API Key")
}

func collectUserApiKey() {
	apiKey := components.StringInput("Enter your API Key:", "")

	if len(apiKey) > 0 {
		saveApiKey(apiKey)
	}

}

func ShowApiOptions() {
	hasApiKey := components.YesNoInput("Do you already have an API key?")

	if hasApiKey {
		collectUserApiKey()
	} else {

		email := components.StringInput("Enter your email address and we'll send you a key:", "")
		// validate email

		if len(email) > 3 {
			sendApiKeyToEmail(email)
		}

		// openWebsite := components.YesNoInput("Open website to get one?")

		// if openWebsite {
		// 	openBrowser(configs.GetWebsiteURL())
		// }
	}

	return
}

func sendApiKeyToEmail(email string) {
	_, errResp := models.RequestTrialAccount(email)
	// handle different error responses.
}

func InvalidApiKey() {
	fmt.Println("Your API Key is inactive, make sure you have an active subscription.")
	getNewApiKey := components.YesNoInput("Enter a new API Key?")

	if getNewApiKey {
		ShowApiOptions()
	} else {
		log.Fatal("Oh-heck exited")
	}
}

func configureApiKeys(cmd *cobra.Command, args []string) {
	// Check file

	forceKey, err := cmd.Flags().GetString("key")
	if forceKey == "" || err != nil {
		// fmt.Println("No key passed")
		// Go through normal key process
		apiKey := getSavedApiKey()

		if apiKey != "" {
			// ask if they want to re-enter their api-key?

			reenterApiKey := components.YesNoInput("You already have an API Key. Would you like to set a new one?")

			if reenterApiKey {
				collectUserApiKey()
			} else {
				return
			}

		} else {
			// go through flow
			ShowApiOptions()
		}
	} else {
		// Set entire config as there's only one field
		saveApiKey(forceKey)
	}

}

func openBrowser(url string) {
	var err error
	fmt.Println("Opening URL: ")
	fmt.Printf(url)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
		fmt.Println(err.Error())
	}

}

func saveApiKey(key string) {
	fmt.Println("Saving key, you can now use oh-heck")
	cfg := configs.ReadConfig()
	cfg.ApiKey = key
	cfg.Save()
}

func getSavedApiKey() string {
	cfg := configs.ReadConfig()
	if cfg == nil || cfg.ApiKey == "" {
		fmt.Println("No API Key")
		return ""
	}
	return cfg.ApiKey
}
