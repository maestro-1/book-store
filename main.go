package main

import (
	"context"
	"encoding/json"
	"go-book-sales-api/book_sales_api"
	"go-book-sales-api/config"
	"go-book-sales-api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)


type apiError struct {
	Err string
	Status int
}

func (e apiError) Error() string{
	return e.Err
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err :=  f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJson(w, e.Status, e)
				return
			}
			writeJson(w, http.StatusInternalServerError, apiError{ Err: "internal server error", Status: http.StatusInternalServerError })
		}
	}
}


func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// ping server to return 200 status
func pingServer(w http.ResponseWriter, r *http.Request) error {
	return writeJson(w, http.StatusOK, "ok")	
}


type APIServer struct {
	addr string
	db *book_sales_api.Queries
}

func NewApiServer(addr string, db *book_sales_api.Queries) *APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) run() error {
	router := mux.NewRouter()
	subroute := router.PathPrefix("/api/v1").Subrouter()
	
	subroute.HandleFunc("/", makeHttpHandler(pingServer))
	bookHander := handlers.NewBookHandler(s.db)

	subroute.HandleFunc("/buyers", makeHttpHandler(bookHander.SellBook))
	

	log.Println("listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}


func main() {
	ctx := context.Background()
	database := config.DatabaseUri

	conn, err := pgx.Connect(ctx, database)
	if err != nil {
		log.Fatal("failed")
	}
	defer conn.Close(ctx)
	queries := book_sales_api.New(conn)

	NewApiServer(":8008", queries).run()
}