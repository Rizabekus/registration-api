package app

import (
	"log"
	"net/http"

	"github.com/Rizabekus/registration-api/internal/handlers"
	"github.com/Rizabekus/registration-api/pkg/loggers"
	"github.com/gorilla/mux"
)

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()

	r.HandleFunc("/register", h.Register).Methods("POST")

	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/modify", h.Modify).Methods("POST")

	loggers.InfoLog.Println("Started the server at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
