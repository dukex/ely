package server

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dukex/ely/pkg/cmd"
	"github.com/dukex/ely/pkg/config"
)

func handleEndpoint(endpoint config.Endpoint, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			status  int
			body    string
			headers string
		)
		err := db.QueryRow("SELECT status, body, headers FROM "+endpoint.Function+"('{}')").Scan(&status, &body, &headers)

		cmd.CheckError(err)

		var headersResult map[string]string
		json.Unmarshal([]byte(headers), &headersResult)

		for header, value := range headersResult {
			w.Header().Set(header, value)
		}
		w.WriteHeader(status)
		io.WriteString(w, body)
	})
}
