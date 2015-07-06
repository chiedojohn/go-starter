package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	msg := os.Getenv("MESSAGE")
	w.Write([]byte(fmt.Sprintf(msg)))
}

func Bye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Bye World!")))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/bye", Bye)

	http.ListenAndServe(":8080", r)
}
