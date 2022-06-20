package deploy

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/dukex/ely/pkg/cmd"
	"github.com/dukex/ely/pkg/database"
	"github.com/spf13/cobra"
	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy functions to database",
		Long:  "Deploy functions to database",
		Run: func(c *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)

			functionsPath, err := c.Flags().GetString("functions-dir")
			cmd.CheckError(err)

			files, err := ioutil.ReadDir(functionsPath)
			cmd.CheckError(err)

			db, err := database.NewDatabase()
			cmd.CheckError(err)

			for _, file := range files {
				if !file.IsDir() {
					log.Println("Deploying", file.Name())
					filePath := filepath.Join(functionsPath, file.Name())

					statement, err := os.ReadFile(filePath)
					cmd.CheckError(err)

					query := string(statement)
					_, err = db.Exec(query)
					cmd.CheckError(err)
				}
			}
		},
	}

	c.PersistentFlags().StringP("functions-dir", "f", "./functions", "directory with funtions to deploy")

	return c
}
