package service

import (
	"errors"
	"github.com/mohamedveron/project_management/domains"
)

func (s *Service) UpdateProject(project domains.Project, id string) (string, error){

	proj := s.findProjectById(id)

	if proj == nil {
		return "not exist", errors.New("project not exist")
	}

	updatedProject := domains.Project{
		ID:           proj.ID,
		Name:         project.Name,
		Owner:        proj.Owner,
		Progress:     project.Progress,
		State:        project.State,
		Participants: proj.Participants,
	}

    s.ProjectsDB[id] = updatedProject

	return "updated", nil
}

