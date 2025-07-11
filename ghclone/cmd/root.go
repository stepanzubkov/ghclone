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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"ghclone/general"
	"ghclone/services"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghclone",
	Short: "Clone multiple repositories from github.",
	Long: `ghclone can clone multiple repositories from your github account or other.
Repositories can be filtered.`,
	Run: MainCommand,
    TraverseChildren: true,
}
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func MainCommand(cmd *cobra.Command, args []string) {
    cfg := services.ParseConfig()
    root_args := services.ParseRootCmdArgs(cmd, args, cfg)
    services.ValidateRootCmdArgs(root_args)

    repos := services.GetUserRepos(root_args.Name, cfg)

    if root_args.Latest {
        repos = FilterLatestRepo(repos)
    }
    if root_args.Choose {
        repos = FilterChooseRepos(repos)
    }

    continue_ := services.InputYesNo(fmt.Sprintf("Found %d repositories. Continue (Y/n)? ", len(repos)), true)
    if !continue_ {
        return
    }
    var err error
    if root_args.Dir == "" {
        root_args.Dir, err = os.Getwd()
        services.CheckIfError(err)
    }

    services.CloneRepositories(repos, root_args.Dir, root_args.Ssh)
}

// TODO: Move these filters
func FilterLatestRepo(repos []any) []any {
    latest_repo := services.GetLatestRepository(repos)
    repos = []any{latest_repo}
    return repos
}

func FilterChooseRepos(repos []any) []any {
    privateRepos := false
    for index, value := range repos {
        repo := value.(map[string]any)
        if repo["visibility"].(string) == "private" {
            privateRepos = true
            fmt.Printf("(%v) %v\n", index, services.Red(repo["name"].(string)))
        } else {
            fmt.Printf("(%v) %v\n", index, repo["name"].(string))
        }
    }
    if privateRepos {
        fmt.Printf("Private repositories marked as %v.\n", services.Red("red"))
    }
    fmt.Print("\nChoose one or multiple repos from list (0, 1, 1 2 3, 1-10 for example): ")
    chosen_indexes := services.SelectFromList(len(repos))
    repos = general.FilterByIndexes(repos, chosen_indexes)
    return repos
}


func init() {
	// Here you will define your flags nd configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ghclone.yaml)")
    rootCmd.Flags().BoolP("all", "a", true, "Clones all user's repositories.")
    rootCmd.Flags().StringP("dir", "d", "", "Specify a directory to clone (defaults to current working directory)")
    rootCmd.Flags().BoolP("latest", "l", false, "Clone 1 latest repository.")
    rootCmd.Flags().BoolP("ssh", "s", false, "Clone via ssh")
    rootCmd.Flags().BoolP("choose", "c", false, "Choose multiple repos to clone from list")
}


