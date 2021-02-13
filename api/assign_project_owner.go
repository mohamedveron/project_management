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

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
