package service

import "github.com/mohamedveron/project_management/domains"

type Service struct {
	ProjectsDB	map[string]domains.Project
}

func NewService(projectsDB	map[string]domains.Project) *Service {

	return &Service{
		ProjectsDB: projectsDB,
	}
}
