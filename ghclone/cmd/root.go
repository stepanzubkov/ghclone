/*
Copyright © 2023 Stepan Zubkov stepanzubkov@florgon.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"os"
    "fmt"

	"github.com/spf13/cobra"
    "net/http"

    "ghclone/services"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghclone",
	Short: "Clone multiple repositories from github.",
	Long: `ghclone can clone multiple repositories from your github account or other.
Repositories can be filtered.`,
	Run: func(cmd *cobra.Command, args []string) {
        dir, err := cmd.Flags().GetString("dir")
        if err != nil {
            fmt.Println(err)
            return
        }
        if len(args) != 1 {
            fmt.Println("Only one argument is allowed!")
            return
        }
        var github_username string = args[0]
        response, err := http.Get("https://api.github.com/users/" + github_username + "/repos")
        if err != nil {
            fmt.Println(err)
            return
        }
        if response.StatusCode == 404 {
            fmt.Println("User not found!")
            return
        }
        repos := services.DecodeJsonResponse(response)
        fmt.Printf("Found %d repositories. Continue (y/n)? ", len(repos))
        var answer string
        fmt.Scanln(&answer)
        if answer != "y" {
            return
        }
        if dir == "" {
            dir, err = os.Getwd()
            if err != nil {
                fmt.Println(err)
                return
            }
        }
        services.CloneRepositories(repos, dir)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ghclone.yaml)")
    rootCmd.Flags().BoolP("all", "a", true, "Clones all user's repositories.")
    rootCmd.Flags().StringP("dir", "d", "", "Specify a directory (default to current working directory)")

}


