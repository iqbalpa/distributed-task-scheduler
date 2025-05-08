package main

import (
	"main/cmd/coordinator"
	"main/internal/api"
	worker "main/internal/workerpool"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	coord := coordinator.New()
	api := api.New(coord)

	router := mux.NewRouter()
	taskRouter := api.HandleRequests()
	router.PathPrefix("/api/v1/task").Handler(taskRouter)

	wp := worker.New(3)
	wp.Start(coord)

	http.ListenAndServe(":8000", router)
}