package cmd

import (
	"fmt"
	"io"
)

//PrintError prints error message
func PrintError(out io.Writer, msg string) {
	fmt.Fprintln(out, "Error : "+msg)
}

//PrintInfo prints info message
func PrintInfo(out io.Writer, msg string) {
	fmt.Fprintln(out, msg)
}

//PrintStatus prints statis message
func PrintStatus(out io.Writer, msg string) {
	fmt.Fprintln(out, msg)
}
