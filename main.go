package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/sirupsen/logrus/hooks/writer"
)
func server() *mux.Router  {
	router := mux.NewRouter()

	router.HandleFunc("/", handleHome).Methods("GET")

	return router
}

func main() {
	router := server()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleHome(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"statusCode" : 200,
		"message" : "Hello World",
	})
}