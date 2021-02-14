package service

import "github.com/mohamedveron/project_management/domains"

func (s *Service) GetEmployees() ([]domains.Employee, error){

	employeesList := []domains.Employee{}

	for _, employees := range s.EmployeesDB {

		employeesList = append(employeesList, employees)
	}

	return employeesList, nil
}
