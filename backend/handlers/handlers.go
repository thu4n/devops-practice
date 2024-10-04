package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader((http.StatusOK))
		return
	}

	fmt.Printf("\nGot an API request!\n")
	message := Message{Text: "Hello from Go API!"}
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Responed with %s", message.Text)
	json.NewEncoder(w).Encode(message)
}
