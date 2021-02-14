package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) AssignProjectOwner(w http.ResponseWriter, r *http.Request, id string){

	 body  := map[string]string{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := s.svc.AssignProjectOwner(id, body["employeeId"])

	if res == "not manager" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("employee not a manager to be the project owner")
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
