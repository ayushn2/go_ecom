package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ayushn2/go_ecom.git/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer{
	return &APIServer{
		addr: address,
		db: db,
	}
}

func (s *APIServer) Run() error{
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)

	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on",s.addr)
	return http.ListenAndServe(s.addr,router)
}