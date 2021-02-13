package service

import "github.com/mohamedveron/project_management/domains"

func (s *Service) findProjectById(id string) *domains.Project{

	if project, ok := s.ProjectsDB[id]; ok {
		return &project
	}

	return nil
}

func (s *Service) findEmployeeById(id string) *domains.Employee{

	if employee, ok := s.EmployeesDB[id]; ok {
		return &employee
	}

	return nil
}
