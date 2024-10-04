package main

import (
	"errors"
	"fmt"
	"go-backend/database"
	"go-backend/handlers"
	"net/http"
	"os"
)

func main() {
	database.Init()
	http.HandleFunc("/message", handlers.GetMessage)
	println("Server is listening at port 5000")

	err := http.ListenAndServe(":5000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting srever: %s\n", err)
		os.Exit(1)
	}
}
