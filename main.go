package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	cli "github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Endpoint struct {
	Path     string `yaml:"path"`
	Function string `yaml:"function"`
}

type ElyConfigFile struct {
	Endpoints []Endpoint `yaml:"endpoint"`
}

func handleEndpoint(endpoint Endpoint, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, endpoint.Path)

		var (
			status  int
			body    string
			headers string
		)
		err := db.QueryRow("SELECT status, body, headers FROM "+endpoint.Function+"()").Scan(&status, &body, &headers)

		check(err)

		var headersResult map[string]string
		json.Unmarshal([]byte(headers), &headersResult)

		for header, value := range headersResult {
			fmt.Println(header)
			w.Header().Set(header, value)
		}

		w.WriteHeader(status)

		io.WriteString(w, body)
	}
}

func main() {
	app := &cli.App{
		Name:  "ely",
		Usage: "the project description here",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "ely.yml",
				Usage:   "Load configuration from `FILE`",
			},
		},
		Action: func(c *cli.Context) error {
			configFile := c.String("config")

			configData, err := os.ReadFile(configFile)

			check(err)

			fmt.Println(configData)

			t := ElyConfigFile{}

			err = yaml.Unmarshal([]byte(configData), &t)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			db, err := sql.Open("postgres", "postgresql://localhost:5432/ely?sslmode=disable")
			check(err)

			err = db.Ping()
			check(err)

			for _, endpoint := range t.Endpoints {
				http.HandleFunc(endpoint.Path, handleEndpoint(endpoint, db))
			}

			err = http.ListenAndServe(":3333", nil)

			check(err)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
