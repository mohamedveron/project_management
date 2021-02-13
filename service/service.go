package service

import "github.com/mohamedveron/project_management/domains"

type Service struct {
	ProjectsDB	map[string]domains.Project
	EmployeesDB map[string]domains.Employee
}

func NewService(projectsDB	map[string]domains.Project, employeesDB map[string]domains.Employee) *Service {

	return &Service{
		ProjectsDB: projectsDB,
		EmployeesDB: employeesDB,
	}
}
