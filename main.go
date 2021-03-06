package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	swag "github.com/swaggo/http-swagger"

	_ "example.com/swag-test/docs"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {

	return &Server{
		router: mux.NewRouter(),
	}
}

func (self *Server) respond(w http.ResponseWriter, req *http.Request, data string) {

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, data)
}

// HandleGet godoc
// @Summary Responds by saying "Hello world!"
// @Tags messages
// @Produce json
// @Success 200 {string} string "Hello world!"
// @Router /sayhello [get]
func (self *Server) HandleGet() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		self.respond(w, req, `{"message":"Hello world!"}`)
	}

}

func (self *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	self.router.ServeHTTP(w, req)
}

func (self *Server) Routes() {

	self.router.HandleFunc("/sayhello", self.HandleGet()).Methods("GET")
	self.router.PathPrefix("/swagger").Handler(swag.WrapHandler)
}

// @title Swagger Example API by Jeff
// @version 1.0
// @description This is a sample server that says "Hello world!".
// @BasePath /v2
func main() {

	server := NewServer()
	server.Routes()

	log.Fatal(http.ListenAndServe(":8080", server))
}
