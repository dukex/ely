package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/dukex/ely/pkg/cmd"
	"github.com/dukex/ely/pkg/config"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "server",
		Short: "Run the ELY server",
		Long:  "Run the ELY server",
		Run: func(c *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)

			port, err := c.Flags().GetInt("port")
			cmd.CheckError(err)

			configPath, err := c.Flags().GetString("config")
			cmd.CheckError(err)

			configuration := config.Load(configPath)

			err = newServer(port, configuration)
			cmd.CheckError(err)
		},
	}

	c.PersistentFlags().IntP("port", "p", 9001, "server port")
	c.PersistentFlags().StringP("config", "c", "./ely.yml", "config YAML file")

	return c
}

func defaultHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("not found"))
	})
}

func newServer(port int, c config.Config) error {
	mux := http.NewServeMux()

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	cmd.CheckError(err)

	err = db.Ping()
	cmd.CheckError(err)

	for _, endpoint := range c.Endpoints {
		mux.Handle(endpoint.Path, logRequest(handleEndpoint(endpoint, db)))
	}

	mux.Handle("/", logRequest(defaultHandler()))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("Server starring at %d", port)

	return server.ListenAndServe()
}
