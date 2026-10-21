package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Обработчик для главного маршрута "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from my DevOps Server on Ubuntu!")
	})

	// Запускаем HTTP-сервер на порту 8080
	fmt.Println("Server is running on http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
