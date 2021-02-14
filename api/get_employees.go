package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetEmployees(w http.ResponseWriter, r *http.Request){
	employees , err := s.svc.GetEmployees()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}
