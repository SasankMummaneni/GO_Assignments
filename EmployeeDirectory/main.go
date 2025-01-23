package main

import (
	"employeeeDirectory/models"
	"employeeeDirectory/repository"
	"employeeeDirectory/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
CRUD

*/

func main() {

	repo := repository.NewEmployeeRepo()

	Execute(repo)

}

func Execute(repo service.EmployeeService) {

	//Create

	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			{

				repo.CreateEmployee(w, r)

			}
		case http.MethodGet:
			{
				path := r.URL.Query().Get("id")
				id, error := strconv.Atoi(path)
				fmt.Println(path)

				if error != nil {
					fmt.Println("Unable to convert to integer")
				}
				repo.GetEmployee(id)
			}
		case http.MethodPut:
			{
				var emp models.Employee
				if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
					fmt.Println("Wrong data is sent")
				}
				repo.UpdateEmployee(emp)

			}
		case http.MethodDelete:
			{
				path := r.URL.Query().Get("id")
				fmt.Println(path)
				id, error := strconv.Atoi(path)

				if error != nil {
					fmt.Println("Unable to convert to integer")
				}
				repo.DeleteEmployee(id)
			}
		default:
			{
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}

		}
	})

	fmt.Println("Starting Server")

	http.ListenAndServe(":8080", nil)

	/*
		//Update

		err := repo.UpdateEmployee(models.Employee{
			EmployeeID:   2,
			EmployeeName: "Bhavani",
			EmployeeAge:  30,
		})

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}

		//GET
		fmt.Println("**********Getting an Employee with ID 2***************")
		val, err := repo.GetEmployee(2)

		fmt.Println(val)

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}

		//Delete

		fmt.Println("**********Deleting an Employee with ID 2***************")
		err = repo.DeleteEmployee(2)

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}
	*/

}
