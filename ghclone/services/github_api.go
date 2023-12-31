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
	"encoding/json"
	"net/http"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
)



func DecodeJsonResponse(response *http.Response) []any {
    var result []any
    decoder := json.NewDecoder(response.Body)
    decoder.Decode(&result)
    return result
}


func CloneRepositories(repos []any, directory string) error {
    for _, value := range repos {
        repo := value.(map[string]any)
        clone_url := repo["clone_url"]
        _, err := git.PlainClone(path.Join(directory, repo["name"].(string)), false, &git.CloneOptions{
            URL: clone_url.(string),
            Progress: os.Stdout,
        })
        if err != nil {
            return err
        }
    }
    return nil
}
