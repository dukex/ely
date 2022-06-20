package ely

import (
	"github.com/spf13/cobra"

	"github.com/dukex/ely/pkg/cmd/cli/db"
	"github.com/dukex/ely/pkg/cmd/cli/deploy"
	"github.com/dukex/ely/pkg/cmd/server"
)

func NewCommand(name string) *cobra.Command {
	c := &cobra.Command{
		Use:   name,
		Short: "API from database",
		Long:  "ELY is a tool for create a API service using database functions",
	}

	c.AddCommand(
		db.NewCommand(),
		server.NewCommand(),
		deploy.NewCommand(),
	)

	return c
}
