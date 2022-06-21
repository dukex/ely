package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/dukex/ely/pkg/cmd"
	"github.com/dukex/ely/pkg/config"
)

type Request struct {
	Method        string            `json:"method"`
	URL           *url.URL          `json:"url"`
	Headers       map[string]string `json:"headers"`
	Proto         string            `json:"proto"`
	ContentLength int64             `json:"content_length"`
	Params        map[string]string `json:"params"`
}

func handleEndpoint(endpoint config.Endpoint, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		cmd.CheckError(err)

		request := Request{
			Method:        r.Method,
			URL:           r.URL,
			Proto:         r.Proto,
			ContentLength: r.ContentLength,
		}

		request.Params = make(map[string]string)
		request.Headers = make(map[string]string)

		for k, v := range r.Header {
			request.Headers[k] = v[0]
		}
		for k, v := range r.Form {
			request.Params[k] = v[0]
		}

		requestParams, err := json.Marshal(request)

		cmd.CheckError(err)

		fn := fmt.Sprintf("boot('%s'::text, '%s'::jsonb)", endpoint.Function, requestParams)
		query := fmt.Sprintf(" SELECT status, body, headers FROM %s", fn)

		var (
			status  int
			body    string
			headers string
		)

		err = db.QueryRow(query).Scan(&status, &body, &headers)

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
