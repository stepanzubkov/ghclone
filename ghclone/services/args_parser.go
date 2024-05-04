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
package services
import (
    "github.com/spf13/cobra"
    "ghclone/models"
)

func ParseRootCmdArgs(cmd *cobra.Command, args []string) *models.RootArgs {
    dir, err := cmd.Flags().GetString("dir")
    CheckIfError(err)
    latest, err := cmd.Flags().GetBool("latest")
    CheckIfError(err)
    choose, err := cmd.Flags().GetBool("choose")
    CheckIfError(err)

    ssh, err := cmd.Flags().GetBool("ssh")
    CheckIfError(err)

    if len(args) != 1 {
        Error("Only one argument is allowed!")
    }
    var github_username string = args[0]

    var root_args models.RootArgs = models.RootArgs{
        Name: github_username,
        Dir: dir,
        Latest: latest,
        Ssh: ssh,
        Choose: choose,
    }
    return &root_args
}

func ValidateRootCmdArgs(root_args *models.RootArgs) {
    if root_args.Latest && root_args.Choose {
        Error("Pass --latest or --choose flag, not both!")
    }
}
