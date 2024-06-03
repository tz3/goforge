package template

type GorillaRouteTemplate struct {
}

func (c GorillaRouteTemplate) Main() []byte {
	return MainTemplate()
}
func (c GorillaRouteTemplate) Server() []byte {
	return MakeHTTPServer()
}
func (c GorillaRouteTemplate) Routes() []byte {
	return MakeGorillaRoutes()
}

func MakeGorillaRoutes() []byte {
	return []byte(`

package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.helloWorldHandler)

	return r
}

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

`)
}
