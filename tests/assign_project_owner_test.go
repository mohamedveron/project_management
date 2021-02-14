package tests

import (
	"github.com/mohamedveron/project_management/service"
	"testing"
)

func TestAssignProjectOwner(t *testing.T){

	projectsDB, employeesDB := PrepareDataForTest()


	service := service.NewService(projectsDB, employeesDB)

	res , err := service.AssignProjectOwner("11bb" ,"11aa")

	if err != nil || res != "updated"{
		t.Error("employee not manager to be the project owner")
	}
}

func TestAssignProjectOwnerButNotManager(t *testing.T){

	projectsDB, employeesDB := PrepareDataForTest()


	service := service.NewService(projectsDB, employeesDB)

	res , err := service.AssignProjectOwner("11bb" ,"t3t3")

	if err == nil || res == "updated"{
		t.Error("employee not manager to be the project owner")
	}
}
