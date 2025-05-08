package api

import (
	"encoding/json"
	"fmt"
	"main/cmd/coordinator"
	"main/internal/task"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Api struct {
	Coordinator *coordinator.Coordinator
}

func New(c *coordinator.Coordinator) *Api {
	return &Api{
		Coordinator: c,
	}
}

func (api *Api) HandleRequests() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/task", api.SubmitTask).Methods("POST")
	r.HandleFunc("/api/v1/task", api.GetAllTasksStatus).Methods("GET")
	r.HandleFunc("/api/v1/task/{id}", api.GetTaskStatus).Methods("GET")
	return r
}

func (api *Api) SubmitTask(w http.ResponseWriter, r *http.Request) {
	var t task.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "failed to process the body", http.StatusBadRequest)
	}
	api.Coordinator.Add(&t)
	json.NewEncoder(w).Encode(
		fmt.Sprintf("successfully added new task: %d", t.Id),
	)
}

func (api *Api) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	status, _ := api.Coordinator.GetStatus(id)
	json.NewEncoder(w).Encode(status)
}

func (api *Api) GetAllTasksStatus(w http.ResponseWriter, r *http.Request) {
	res, _ := api.Coordinator.GetAll()
	json.NewEncoder(w).Encode(res)
}
