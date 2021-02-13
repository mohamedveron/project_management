package service

import (
	"github.com/mohamedveron/project_management/domains"
)

func (s *Service) GetProjects() ([]domains.Project, error){

	projectsList := []domains.Project{}

	for _, project := range s.ProjectsDB {

		projectsList = append(projectsList, project)
	}

	return projectsList, nil
}
