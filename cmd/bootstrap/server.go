package bootstrap

import (
	"log"
	"login-app/internal/platform/storage/newsql"
	"net/http"
)

type Server struct {
	DB *newsql.ConfigPostgresDB
}

func NewServer(db *newsql.ConfigPostgresDB) Server {
	return Server{DB: db}
}

func (s *Server) Run() {
	s.registerRouter()
}

func (s *Server) registerRouter() {
	router := http.NewServeMux()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	server := http.Server{Addr: ":8080", Handler: router}
	log.Println("Server running on port", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
