package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dukex/ely/pkg/cmd"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var port int

	c := &cobra.Command{
		Use:   "server",
		Short: "Run the ELY server",
		Long:  "Run the ELY server",
		Run: func(c *cobra.Command, args []string) {
			log.SetOutput(os.Stdout)

			err := newServer(port)

			cmd.CheckError(err)
		},
	}

	c.PersistentFlags().IntVarP(&port, "port", "p", 9001, "server port")

	return c
}

func newServer(port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("Server starring at %d", port)

	return server.ListenAndServe()
}
