package tests

import (
	"github.com/mohamedveron/project_management/service"
	"testing"
)

func TestAssignProjectParticipants(t *testing.T){

	projectsDB, employeesDB := PrepareDataForTest()


	service := service.NewService(projectsDB, employeesDB)

	_ , err := service.AssignProjectOwner("11bb" ,"11aa")

	participants := []string{"33cc", "t3t3"}

	res , err := service.AssignProjectParticipants("11bb" , participants)

	if err != nil || res != "added"{
		t.Error("employee from different department of the owner")
	}
}

func TestAssignProjectParticipantsWithDifferentDepartment(t *testing.T){

	projectsDB, employeesDB := PrepareDataForTest()


	service := service.NewService(projectsDB, employeesDB)

	_ , err := service.AssignProjectOwner("11bb" ,"11aa")

	participants := []string{"58kk", "t3t3"}

	res , err := service.AssignProjectParticipants("11bb" , participants)

	if err == nil || res == "added"{
		t.Error("employee from different department of the owner")
	}
}
