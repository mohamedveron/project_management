package api

import (
	"net/http"
	"encoding/json"
)

func (s *Server) GetProjects(w http.ResponseWriter, r *http.Request){

	projects , err := s.svc.GetProjects()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
