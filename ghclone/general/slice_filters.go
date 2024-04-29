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
package general


/* Remove duplicates from slice */
func RemoveDuplicateValues(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}
 
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}


/* Filters elements from `slice` only with indexes presented in `indexes` slice 
Works properly only with valid indexes */
func FilterByIndexes[T any](slice []T, indexes []int) []T {
    filtered_slice := []T{}
    for _, index := range indexes {
        filtered_slice = append(filtered_slice, slice[index])
    }
    return filtered_slice
}
