package service

import (
	"github.com/mohamedveron/project_management/domains"
	"math/rand"
)

func (s *Service) CreateProject(project domains.Project) (string, error){

	// generate new id
	id := randGeneratePassword(4)

	project.ID = id

	s.ProjectsDB[id] = project

	return id, nil
}



func randGeneratePassword(n int) string {

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}