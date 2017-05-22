package cmd

import (
	"fmt"
	"io"
)

func PrintError(out io.Writer, msg string) {
	fmt.Fprintln(out, "Error : "+msg)
}

func PrintInfo(out io.Writer, msg string) {
	fmt.Fprintln(out, msg)
}

func PrintStatus(out io.Writer, msg string) {
	fmt.Fprintln(out, msg)
}
