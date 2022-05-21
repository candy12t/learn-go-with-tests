package main

import (
	"log"
	"net/http"

	"github.com/candy12t/app/memory"
	"github.com/candy12t/app/server"
)

func main() {
	store := memory.NewInMemoryPlayerStore()
	srv := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8000", srv))
}
