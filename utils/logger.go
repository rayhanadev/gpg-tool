package utils

import "fmt"

func PrintVerbose(verbose bool, message string) {
    if verbose {
        fmt.Println(message)
    }
}
