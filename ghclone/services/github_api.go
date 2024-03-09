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
	"encoding/json"
	"net/http"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)



func DecodeJsonResponse(response *http.Response) []any {
    var result []any
    decoder := json.NewDecoder(response.Body)
    decoder.Decode(&result)
    return result
}


func CloneRepositories(repos []any, directory string, ssh bool) error {
    for _, value := range repos {
        repo := value.(map[string]any)

        var clone_url string
        var err error
        if (ssh) {
            clone_url = repo["ssh_url"].(string)
            public_keys, err_ := GetPublicKeys()
            if err_ != nil {
                return err_
            }
            _, err = git.PlainClone(path.Join(directory, repo["name"].(string)), false, &git.CloneOptions{
                Auth: public_keys,
                URL: clone_url,
                Progress: os.Stdout,
            })
        } else {
            clone_url = repo["clone_url"].(string)
            _, err = git.PlainClone(path.Join(directory, repo["name"].(string)), false, &git.CloneOptions{
                URL: clone_url,
                Progress: os.Stdout,
            })
        }

        if err != nil {
            return err
        }
    }
    return nil
}

func GetPublicKeys() (*ssh.PublicKeys, error) {
    homedir, _ := os.UserHomeDir()
    private_key_file := path.Join(homedir, ".ssh/id_ed25519")
    _, err := os.Stat(private_key_file)
	if err != nil {
		return nil, errors.New("Ssh id file is not found!")
	}

	public_keys, err := ssh.NewPublicKeysFromFile("git", private_key_file, "")
	if err != nil {
		return nil, err
	}
    return public_keys, nil
}
