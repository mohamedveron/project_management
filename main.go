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
	
	employee1 := domains.Employee{
		ID:         "11",
		FirstName:  "Peter",
		LastName:   "Golm",
		Role:       "senior developer",
		Email:      "peter_golm@visable.com",
		Department: "engineering",
	}

	employee2 := domains.Employee{
		ID:         "33",
		FirstName:  "Andreas",
		LastName:   "Litt",
		Role:       "developer",
		Email:      "Andreas_Litt@visable.com",
		Department: "engineering",
	}

	project1 := domains.Project{
		ID:           "1",
		Name:         "project management",
		Owner:        employee1,
		Progress:     0,
		State:        domains.EnumProjectStatePlanned,
		Participants: []domains.Employee{employee2},
	}

	 projectsDB := make(map[string]domains.Project)

	projectsDB["1"] = project1

	serviceLayer := service.NewService(projectsDB)
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

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server failed to start", "error", err)
		}

	fmt.Println("server started...")


	arr := []int{42,43,44,45,46}
	arr1 := []int{47,48,49,50, 51}

	arr2 := arr[2:5]
	arr2 = append(arr2, arr1[0:2]...)

	mm := make(map[string][]string)
	mm["veron"] = []string{"aa", "bb"}
	mm["shy5o"] = []string{"cc", "dd"}

	delete(mm, "shy5o")


}
