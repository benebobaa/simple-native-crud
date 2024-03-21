package main

import (
	"fmt"
	"net/http"
)

func main() {
	db := NewDB("postgres", "user=root dbname=tutorial password=root sslmode=disable")
	router := NewRouter(db)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("listening to port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
