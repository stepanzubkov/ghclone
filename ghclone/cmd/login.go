/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"ghclone/general"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to github account via API token",
	Long: `This command is used to login to github account via API token.
It's used for cloning private repositories.`,
	Run: LoginCommand,
}


func LoginCommand(cmd *cobra.Command, args []string) {
    var apiToken string
    for {
        apiToken = general.GetPasswordInput("Paste github API auth token: ")
        if strings.HasPrefix(apiToken, "ghp_") {
            break
        }
        fmt.Println("Your input doesn't look like a github token! It should has prefix 'ghp_'")
    }
    fmt.Println(apiToken)
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
