package db

import (
	"os"

	"github.com/dukex/ely/pkg/cmd"
	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func NewCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "db",
		Short: "Manage ELY database",
		Long:  "Manage ELY database",
	}

	c.AddCommand(&cobra.Command{
		Use: "setup",
		Run: func(c *cobra.Command, args []string) {
			databaseURL := os.Getenv("DATABASE_URL")
			sourceURL := "github://dukex/ely/migrations#main"

			m, err := migrate.New(sourceURL, databaseURL)

			cmd.CheckError(err)

			m.Up()
		},
	},
	)

	return c
}
