package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer interface {
	Start()
	Stop()
	AddHttpHandler(path string, methods string, handler http.Handler)
	AddStaticHandler(path string, handler http.Handler)
}

type httpServer struct {
	server *http.Server
	router *mux.Router
}

func NewHttpServer() HttpServer {
	muxRouter := mux.NewRouter().StrictSlash(true)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: muxRouter,
	}

	return &httpServer{
		server: server,
		router: muxRouter,
	}
}

func (h *httpServer) Start() {
	fmt.Println("serving at port: 8080")
	h.server.ListenAndServe()
}

func (h *httpServer) Stop() {
	h.server.Shutdown(context.TODO())
}

func (h *httpServer) AddHttpHandler(path string, methods string, handler http.Handler) {
	h.router.HandleFunc(path, handler.ServeHTTP).Methods(methods)
}

func (h *httpServer) AddStaticHandler(path string, handler http.Handler) {
	h.router.PathPrefix(path).Handler(handler)
}
