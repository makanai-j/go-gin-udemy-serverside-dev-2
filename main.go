package main

import (
	"go-gin-udemy-serverside-dev-2/controllers"
	"go-gin-udemy-serverside-dev-2/domain"
	"go-gin-udemy-serverside-dev-2/services"
	"log"
	"net/http"
)

func main() {
	repo := domain.NewTradeRepoInMem()
	svc  := services.NewTradeService(repo, nil)
	th   := controllers.NewTradeHandler(svc)
	mux  := controllers.NewMux(th)

	addr := ":8080"
	log.Println("listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
