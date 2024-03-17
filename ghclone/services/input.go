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
	"bufio"
	"os"
	"strings"
    "strconv"
)

func SelectFromList(length int) []int {
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    CheckIfError(err)
    line = strings.TrimSpace(line)

    indexes := strings.Split(line, " ")
    var int_indexes []int
    for _, index := range indexes {
        int_index, err := strconv.ParseInt(index, 10, 32)
        if err != nil {
            Error("Index '%s' is not integer!", index)
        }
        if int_index < 0 || int(int_index) > length-1 {
            Error("Index %s is out of range", index)
        }
        int_indexes = append(int_indexes, int(int_index))
    }
    return int_indexes
}
