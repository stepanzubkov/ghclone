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
    "errors"

    "github.com/spf13/cobra"
    "ghclone/models"
)

func ParseRootCmdArgs(cmd *cobra.Command, args []string) (*models.RootArgs, error) {
    dir, err := cmd.Flags().GetString("dir")
    if err != nil {
        return nil, err
    }
    latest, err := cmd.Flags().GetBool("latest")
    if err != nil {
        return nil, err
    }

    ssh, err := cmd.Flags().GetBool("ssh")
    if err != nil {
        return nil, err
    }

    if len(args) != 1 {
        return nil, errors.New("Only one argument is allowed!")
    }
    var github_username string = args[0]

    var root_args models.RootArgs = models.RootArgs{Name: github_username, Dir: dir, Latest: latest, Ssh: ssh}
    return &root_args, nil
}
