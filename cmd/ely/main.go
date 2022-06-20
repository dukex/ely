package main

import (
	"os"
	"path/filepath"

	"github.com/dukex/ely/pkg/cmd"
	"github.com/dukex/ely/pkg/cmd/ely"
)

func main() {
	baseName := filepath.Base(os.Args[0])

	err := ely.NewCommand(baseName).Execute()
	cmd.CheckError(err)
}
