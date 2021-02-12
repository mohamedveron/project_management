package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mohamedveron/project_management/api"
	"github.com/mohamedveron/project_management/service"
	"net/http"
)


func main() {

	serviceLayer := service.NewService()
	server := api.NewServer(serviceLayer)

	// prepare router
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/", api.Handler(server))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "80",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server failed to start", "error", err)
		}
	}()

	fmt.Println("server started...")


	arr := []int{42,43,44,45,46}
	arr1 := []int{47,48,49,50, 51}

	arr2 := arr[2:5]
	arr2 = append(arr2, arr1[0:2]...)

	mm := make(map[string][]string)
	mm["veron"] = []string{"aa", "bb"}
	mm["shy5o"] = []string{"cc", "dd"}

	delete(mm, "shy5o")
	defer fmt.Println("done")


}
