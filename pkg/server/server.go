/*
  Copyright 2019 KompiTech GmbH

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/KompiTech/fabric-sdk-server-go/pkg/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

// Server HTTP server implementation
type Server struct {
	router *mux.Router
	url    string
}

// NewServer returns new *Server instance
func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
		url:    viper.GetString("server.address") + ":" + viper.GetString("server.port"),
	}
}

// RegisterHTTPHandlers registers handler functions for HTTP server
func (s *Server) RegisterHTTPHandlers() {
	s.router.HandleFunc("/api/v1/chaincode/query", s.httpHandleCCQuery)
	s.router.HandleFunc("/api/v1/chaincode/invoke", s.httpHandleCCInvoke)
}

// Start starts listening HTTP
func (s *Server) Start() error {
	s.RegisterHTTPHandlers()
	log.Print(fmt.Sprintf("Listening HTTP on: %s", s.url))

	handler := CORSWrap(s.router)
	return http.ListenAndServe(s.url, handler)
}

type chaincodeHTTPRequest struct {
	User      string   `schema:"user"`
	Channel   string   `schema:"channel"`
	Chaincode string   `schema:"chaincode"`
	Fcn       string   `schema:"fcn"`
	Args      []string `schema:"args"`
}

func (s *Server) httpHandleCCQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400: Bad Request"))
		return
	}

	ccResponse, err := callChaincode(w, r, "query")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(ccResponse.Payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

}

func (s *Server) httpHandleCCInvoke(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400: Bad Request"))
		return
	}

	ccResponse, err := callChaincode(w, r, "invoke")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(ccResponse.Payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func callChaincode(w http.ResponseWriter, r *http.Request, method string) (*channel.Response, error) {
	var (
		response channel.Response
		req      chaincodeHTTPRequest
	)

	decoder := schema.NewDecoder()
	if err := decoder.Decode(&req, r.URL.Query()); err != nil {
		return nil, err
	}
	if req.User == "" {
		req.User = viper.GetString("fabric.user")
	}

	argBytes := convertArgs(req.Args)

	configpath := viper.GetString("fabric.configpath")
	client, close, err := fabric.NewChannelClient(req.Channel, configpath, req.User)
	if err != nil {
		return nil, err
	}
	defer close()

	ccReq := channel.Request{ChaincodeID: req.Chaincode, Fcn: req.Fcn, Args: argBytes}

	switch method {
	case "query":
		response, err = client.Query(ccReq)
		if err != nil {
			return nil, err
		}
	case "invoke":
		response, err = client.Execute(ccReq)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Unrecognized chaincode request method")
	}

	return &response, nil
}

func convertArgs(args []string) [][]byte {
	argBytes := [][]byte{}
	for _, arg := range args {
		argBytes = append(argBytes, []byte(arg))
	}
	return argBytes
}

// CORSWrap wraps http.Handler with CORS
func CORSWrap(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8081", "http://localhost:8081", "https://127.0.0.1:8081", "https://localhost:8081"},
		AllowedMethods:   []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		MaxAge:           3200,
		AllowCredentials: true,
	})
	return c.Handler(handler)
}
