package cmd

import (
	"context"
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		if err != context.Canceled {
			fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
		}
		os.Exit(1)
	}
}

func Exit(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
