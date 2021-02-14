package service

import (
	"errors"
	"github.com/mohamedveron/project_management/domains"
)

func (s *Service) AssignProjectParticipants(projectId string, participantsIds []string)(string , error){

	project := s.findProjectById(projectId)

	if project == nil {
		return "not exist", errors.New("project not exist")
	}

	participants := []domains.Employee{}

	if project.Participants != nil && len(project.Participants) > 0 {
		participants = project.Participants
	}

	for _, id := range participantsIds {

		owner := s.findEmployeeById(id)

		if owner == nil {
			return "not exist", errors.New("employee not exist")
		}

		if owner.Department != project.Owner.Department {
			return "different department", errors.New("employee from different department of the owner")
		}

		participants = append(participants, *owner)
	}

	updatedProject := domains.Project{

		ID:           project.ID,
		Name:         project.Name,
		Owner:        project.Owner,
		Progress:     project.Progress,
		State:        project.State,
		Participants: participants,
	}

	s.ProjectsDB[projectId] = updatedProject

	return "added", nil
}
