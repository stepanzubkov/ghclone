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
    "time"
)



func GetLatestRepository(repos []any) (map[string]any, error) {
    var latest_repo map[string]any
    var latest_created_at = time.Unix(0, 0)
    for _, value := range repos {
        repo := value.(map[string]any)
        repo_created_at, err := time.Parse("2006-01-02T15:04:05Z", repo["created_at"].(string))
        if err != nil {
            return nil, err
        }
        if repo_created_at.Compare(latest_created_at) == 1 {
            latest_repo = repo
            latest_created_at = repo_created_at
        }
    }
    return latest_repo, nil

}
