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
    "net/http"
)


/*
Requests for user repos. If user does not exist - print error and exit.
If user exists - return decoded repos slice.
*/
func GetUserRepos(username string) []any {
    response, err := http.Get("https://api.github.com/users/" + username + "/repos?per_page=100")
    CheckIfError(err)
    if response.StatusCode == 404 {
        Error("User not found!")
    }
    repos := decodeJsonResponse(response)
    return repos
}


// Checks if access token is valid via github api and returns true if access token is valid
func CheckAccessToken(cfg *Config) bool {
    response := makeApiRequest("user/repos", cfg)
    return response.StatusCode == 200
}
