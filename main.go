package main

import (
	"main/internal/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	api := api.Api{}
	api.Init()

	router := mux.NewRouter()
	taskRouter := api.HandleRequests()
	router.PathPrefix("/api/v1/task").Handler(taskRouter)
	

	http.ListenAndServe(":8000", router)
}