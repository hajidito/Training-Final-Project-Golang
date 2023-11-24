package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct{
	ID int
	Name string
	Age int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Dio", Age: 12, Division: "Surya"},
	{ID: 1, Name: "Ijal", Age: 12, Division: "Pro Mild"},
	{ID: 1, Name: "Irfan", Age: 12, Division: "Relaxa"},
}

var PORT = ":8080"

func mainY(){
	http.HandleFunc("/employees",GetEmployees)
	http.HandleFunc("/employee",CreateEmployee)
	fmt.Println("application is listening on port:", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err !=  nil {
		panic(err.Error())
	}
}

func GetEmployees(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}
	http.Error(w,"invalid method",http.StatusBadRequest)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")
		convertAge, err := strconv.Atoi(age)
		if err != nil{
			http.Error(w, "invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID: len(employees)+1, 
			Name: name, 
			Age: convertAge, 
			Division: division,
		}
		employees = append(employees, newEmployee)
		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w,"invalid method",http.StatusBadRequest)
}