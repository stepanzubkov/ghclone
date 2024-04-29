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


/* Checks if Github User with `username` exists and exits application with error if not*/
func CheckIfUserExists(username string) {
    response, err := http.Get("https://api.github.com/users/" + username + "/repos")
    CheckIfError(err)
    if response.StatusCode == 404 {
        Error("User not found!")
    }
}
