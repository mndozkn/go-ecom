package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mndozkn/go-ecom/service/user"
)

type APISERVER struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APISERVER {
	return &APISERVER{
		addr: addr,
		db:   db,
	}
}

func (s *APISERVER) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
