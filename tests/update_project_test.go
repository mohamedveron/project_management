package tests

import (
	"github.com/mohamedveron/project_management/domains"
	"github.com/mohamedveron/project_management/service"
	"testing"
)

func TestUpdateProject(t *testing.T){

	projectsDB, employeesDB := PrepareDataForTest()

	project := domains.Project{
		ID:           "11bb",
		Name:         "project1",
		Owner:        projectsDB["11bb"].Owner,
		Progress:     20,
		State:        "active",
		Participants: projectsDB["11bb"].Participants,
	}

	service := service.NewService(projectsDB, employeesDB)

	res , err := service.UpdateProject(project, "11bb")

	if err != nil || res != "updated"{
		t.Error("can't update project")
	}
}
