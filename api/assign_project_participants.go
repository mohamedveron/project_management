package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) AssignProjectParticipants(w http.ResponseWriter, r *http.Request, id string){
	body  := map[string][]string{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := s.svc.AssignProjectParticipants(id, body["employeeIds"])

	if res == "different department" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("can't assign employee from different department of the owner in the project")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
