package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-backend/database"
	"net/http"
	"os"
)

type Message struct {
	Text string `json:"text"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

func main() {
	database.ConnectDb()

	http.HandleFunc("/message", getMessage)
	println("Server is listening at port 5000")

	err := http.ListenAndServe(":5000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting srever: %s\n", err)
		os.Exit(1)
	}
}
