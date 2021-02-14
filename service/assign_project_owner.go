package service

import (
	"errors"
	"github.com/mohamedveron/project_management/domains"
)

func (s *Service) AssignProjectOwner(projectId string, ownerId string)(string , error){

	project := s.findProjectById(projectId)

	if project == nil {
		return "not exist", errors.New("project not exist")
	}

	owner := s.findEmployeeById(ownerId)

	if owner == nil {
		return "not exist", errors.New("employee not exist")
	}

	if owner.Role != "manager" {
		return "not manager", errors.New("employee not manager to be the project owner")
	}

	updatedProject := domains.Project{
		ID:           project.ID,
		Name:         project.Name,
		Owner:       *owner,
		Progress:     project.Progress,
		State:        project.State,
		Participants: project.Participants,
	}

	s.ProjectsDB[projectId] = updatedProject

	return "updated", nil
}
