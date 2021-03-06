package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mohamedveron/project_management/api"
	"github.com/mohamedveron/project_management/domains"
	"github.com/mohamedveron/project_management/service"
	"net/http"
)


func main() {

	// make a list of employees to do the operation instead of db
	employee1 := domains.Employee{
		ID:         "11aa",
		FirstName:  "Peter",
		LastName:   "Golm",
		Role:       "manager",
		Email:      "peter_golm@visable.com",
		Department: "Engineering",
	}

	employee2 := domains.Employee{
		ID:         "33cc",
		FirstName:  "Andreas",
		LastName:   "Litt",
		Role:       "developer",
		Email:      "Andreas_Litt@visable.com",
		Department: "Engineering",
	}

	employee3 := domains.Employee{
		ID:         "58kk",
		FirstName:  "Nina",
		LastName:   "Wessels",
		Role:       "hr",
		Email:      "Nina_Wessels@visable.com",
		Department: "HR",
	}

	employee4 := domains.Employee{
		ID:         "t3t3",
		FirstName:  "Mohamed",
		LastName:   "Abdel Mohaimen",
		Role:       "developer",
		Email:      "Mohamed_veron@gmail.com",
		Department: "Engineering",
	}

	// make a list of projects to do the operation instead of db
	project1 := domains.Project{
		ID:           "11bb",
		Name:         "project management",
		Owner:        domains.Employee{},
		Progress:     0,
		State:        domains.EnumProjectStatePlanned,
		Participants: []domains.Employee{},
	}

	projectsDB := make(map[string]domains.Project)
	employeesDB := make(map[string]domains.Employee)

	employeesDB["11aa"] = employee1
	employeesDB["33cc"] = employee2
	employeesDB["58kk"] = employee3
	employeesDB["t3t3"] = employee4

	projectsDB["11bb"] = project1

	serviceLayer := service.NewService(projectsDB, employeesDB)
	server := api.NewServer(serviceLayer)

	// prepare router
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Mount("/", api.Handler(server))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	fmt.Println("server starting on port 8080...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}
