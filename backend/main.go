package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Text string `json:"text"`
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got an API request!\n")
	message := Message{Text: "Hello from Go API!"}
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Responed with %s", message.Text)
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/api/message", getMessage)
	println("Server is listening at port 8000")

	err := http.ListenAndServe(":8000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting srever: %s\n", err)
		os.Exit(1)
	}
}
