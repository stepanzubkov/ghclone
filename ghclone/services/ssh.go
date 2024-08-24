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
    "os"
    "path"
    "strings"
	"path/filepath"
	"io/fs"
    "github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func getPublicKeys() *ssh.PublicKeys {
    private_key_file := findSshKeyFile()
    _, err := os.Stat(private_key_file)
    if err != nil {
        PrintFatal("Ssh private key file is not found!")
    }

    public_keys, err := ssh.NewPublicKeysFromFile("git", private_key_file, "")
    CheckIfError(err)
    return public_keys
}

func findSshKeyFile() string {
    homedir, _ := os.UserHomeDir()
    files := findFiles(path.Join(homedir, ".ssh"), func(s string) bool {
        return strings.HasPrefix(filepath.Base(s), "id_") && filepath.Ext(s) == ""
    })
    if len(files) != 1 {
        PrintFatal("Can't choose ssh private key file!")
    }
    return files[0]
}

func findFiles(root string, fn func(string)bool) []string {
   var files []string
   filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
      if fn(s) {
         files = append(files, s)
      }
      return nil
   })
   return files
}
