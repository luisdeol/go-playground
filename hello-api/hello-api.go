package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

var employees []Employee;

func main(){
	router := mux.NewRouter()

	employees = append(employees, Employee{ID: "1", Name: "Luis", Company: &Company{ID: "1", Name: "MP"}})
	employees = append(employees, Employee{ID: "2", Name: "Joel", Company: &Company{ID: "2", Name: "MS"}})
	employees = append(employees, Employee{ID: "3", Name: "FonFon", Company: &Company{ID: "3", Name: "GL"}})

	router.HandleFunc("/employees", GetEmployees).Methods("GET")

	http.ListenAndServe(":8080", router)
}

type Employee struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Company *Company `json:"company,omitempty"`
}

type Company struct { 
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees);
}