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
	"fmt"
	"ghclone/general"
	"os"
	"strconv"
	"strings"
)

func SelectFromList(length int) []int {
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    CheckIfError(err)
    line = strings.TrimSpace(line)

    indexes := strings.Split(line, " ")
    return general.RemoveDuplicateValues(strSliceToIntSlice(length, indexes))
}


func parseIndexesRange(length int, indexes_range string) []int {
    range_borders := strings.Split(indexes_range, "-")
    if len(range_borders) != 2 {
        PrintFatal("Invalid indexes range: '%s'", indexes_range)
    }
    int_range_borders := strSliceToIntSlice(length, range_borders)

    if int_range_borders[1] < int_range_borders[0] {
        temp := int_range_borders[1]
        int_range_borders[1] = int_range_borders[0]
        int_range_borders[0] = temp
    }

    var expanded_range []int
    for i:=int_range_borders[0]; i<=int_range_borders[1]; i++ {
        expanded_range = append(expanded_range, i)
    }
    return expanded_range
}

// TODO: Choose other name...
func strSliceToIntSlice(length int, str_slice []string) []int {
    var int_indexes []int
    for _, index := range str_slice {
        if strings.Contains(index, "-") {
            int_indexes = append(int_indexes, parseIndexesRange(length, index)...)
            continue
        }
        int_index, err := strconv.ParseInt(index, 10, 32)
        if err != nil {
            PrintFatal("Index '%s' is not integer!", index)
        }
        if int_index < 0 || int(int_index) > length-1 {
            PrintFatal("Index %s is out of range", index)
        }
        int_indexes = append(int_indexes, int(int_index))
    }
    return int_indexes
}


func InputYesNo(label string, defaultYes bool) bool {
    fmt.Print(label)
    var answer string
    fmt.Scanln(&answer)
    answer = strings.ToLower(strings.TrimSuffix(answer, "\n"))
    if answer == "" {
        return defaultYes
    }
    if answer == "y" || answer == "yes" {
        return true
    }
    return false
}
