package api

import (
	"encoding/json"
	"github.com/mohamedveron/project_management/domains"
	"net/http"
)

func (s *Server) CreateProject(w http.ResponseWriter, r *http.Request){

	var newProject NewProject

	if err := json.NewDecoder(r.Body).Decode(&newProject); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	project := domains.Project{
		ID:           "",
		Name:         newProject.Name,
		Owner:        domains.Employee{},
		Progress:     newProject.Progress,
		State:        domains.EnumProjectState(newProject.State),
		Participants: nil,
	}

	id, err := s.svc.CreateProject(project)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
