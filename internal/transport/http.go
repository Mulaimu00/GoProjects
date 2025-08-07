package transport

import (
	"encoding/json"
	"log"
	"my-first-api/internal/todo"
	"net/http"
	"strconv"
)

type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(todoSvc.GetAll())
		if err != nil {
			log.Println(err)
			return
		}

	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		todoSvc.Add(t.Item)
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("DELETE /todo/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := request.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if id < 0 || id >= (todoSvc.Len()) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		err = todoSvc.Remove(id)
		if err != nil {
			log.Printf("failed to remove todo at index %d: %v", id, err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}

	})
	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
