package tests

import (
	"github.com/mohamedveron/project_management/domains"
	"github.com/mohamedveron/project_management/service"
	"testing"
)

	var projectsDB = make(map[string]domains.Project)
	var employeesDB = make(map[string]domains.Employee)


	var employee1 = domains.Employee{
		ID:         "11aa",
		FirstName:  "Peter",
		LastName:   "Golm",
		Role:       "manager",
		Email:      "peter_golm@visable.com",
		Department: "Engineering",
	}

	var employee2 = domains.Employee{
		ID:         "33cc",
		FirstName:  "Andreas",
		LastName:   "Litt",
		Role:       "developer",
		Email:      "Andreas_Litt@visable.com",
		Department: "Engineering",
	}

	var employee3 = domains.Employee{
		ID:         "58kk",
		FirstName:  "Nina",
		LastName:   "Wessels",
		Role:       "hr",
		Email:      "Nina_Wessels@visable.com",
		Department: "HR",
	}

	var employee4 = domains.Employee{
		ID:         "t3t3",
		FirstName:  "Mohamed",
		LastName:   "Abdel Mohaimen",
		Role:       "developer",
		Email:      "Mohamed_veron@gmail.com",
		Department: "Engineering",
	}

	// make a list of projects to do the operation instead of db
	var project1 = domains.Project{
		ID:           "11bb",
		Name:         "project management",
		Owner:        domains.Employee{},
		Progress:     0,
		State:        domains.EnumProjectStatePlanned,
		Participants: []domains.Employee{},
	}


func TestCreateProject(t *testing.T){

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
