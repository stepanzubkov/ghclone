
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
    "fmt"
    "os"
)

const (
    red = "\x1b[31;1m"
    green = "\x1b[32;1m"
    yellow = "\x1b[33;1m"

    cclear = "\x1b[0m\n"
)

// Print message formatted with fmt.Printf and specified ANSI color sequence
func printFormatted(color string, format string, a ...any) {
    fmt.Printf(color + "%s" + cclear, fmt.Sprintf(format, a...))
}

// Print formatted message with error style (red fg)
func PrintError(format string, a ...any) {
    printFormatted(red, format, a...)
}

// Calls PrintError, then os.Exit(1)
func PrintFatal(format string, a ...any) {
    PrintError(format, a...)
    os.Exit(1)
}

// If err is nil, just returns. Else calls PrintFatal on err.
func CheckIfError(err error) { 
    if err == nil {
        return
    }

    PrintFatal(err.Error())
}

// Print formatted message with success style (green fg)
func PrintSuccess(format string, a ...any) {
    printFormatted(green, format, a...)
}

// Print formatted message with warning style (yellow fg)
func PrintWarning(format string, a ...any) {
    printFormatted(yellow, format, a...)
}
