package tests

import (
	"github.com/mohamedveron/project_management/domains"
	"github.com/mohamedveron/project_management/service"
	"testing"
)

func TestCreateProject(t *testing.T){

	projectsDB, employeesDB := PrepareDataForTest()

	project := domains.Project{
		ID:           "",
		Name:         "project1",
		Owner:        domains.Employee{},
		Progress:     0,
		State:        "planned",
		Participants: nil,
	}

	service := service.NewService(projectsDB, employeesDB)

	_ , err := service.CreateProject(project)

	if err != nil {
		t.Error("can't create project")
	}
}
