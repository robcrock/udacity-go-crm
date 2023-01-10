package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customer struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Contacted bool `json:"contacted"`
}

var customerDatabase []customer

func seedCustomers() {
	c1 := customer{
			ID: "0",
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}
	c2 := customer{
			ID: "1",
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}
	c3 := customer{
			ID: "2",
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}

	customerDatabase = append(customerDatabase, c1, c2, c3)
}

// The home route is a client API endpoint, and includes a brief overview of the API (e.g., available endpoints).
func getAPIDescription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	t, _  := template.ParseFiles("index.html")

	t.Execute(w, nil)

}

// Getting a single customer through a /customers/{id} path
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	customerFound := false

	for _, v := range customerDatabase {
		if v.ID == id {
			customerFound = true
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	if !customerFound {
		w.WriteHeader(http.StatusNotFound)
	}

}

// Getting all customers through a the /customers path
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customerDatabase)

}

// Creating a customer through a /customers path
func addCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate Content-Type in the response header
	w.Header().Set("Content-Type", "application/json")

	// Create (but not yet assign values to) for the new entry
	var c customer
	customerAlreadyExists := false


	// Read the HTTP request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Encode the request body
	json.Unmarshal(reqBody, &c)
	c.ID = strconv.Itoa(rand.Intn(1000000))


	for _, v := range customerDatabase {
		if v.ID == c.ID {
			customerAlreadyExists = true
			w.WriteHeader(http.StatusConflict)
		}
	}

	if !customerAlreadyExists {
		customerDatabase = append(customerDatabase, c)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customerDatabase)
	}

}

// Updating a customer through a /customers/{id} path
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate Content-Type in the response header
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	customerFound := false

	// Create (but not yet assign values to) for the new entry
	var c customer

	// Read the HTTP request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Encode the request body
	json.Unmarshal(reqBody, &c)

	for k, v := range customerDatabase {
		if v.ID == id {
			customerFound = true
			customerDatabase = append(customerDatabase[:k], customerDatabase[k+1:]...)
			c.ID = id
			customerDatabase = append(customerDatabase, c)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customerDatabase)
		}
	}

	if !customerFound {
		w.WriteHeader(http.StatusNotFound)
	}

}

// Deleting a customer through a /customers/{id} path
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate Content-Type in the response header
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	customerFound := false


	for k, v := range customerDatabase {
		if v.ID == id {
			customerFound = true
			customerDatabase = append(customerDatabase[:k], customerDatabase[k+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customerDatabase)
			return
		}
	}

	if !customerFound {
		w.WriteHeader(http.StatusNotFound)
	}

}

func main() {
	seedCustomers()

	router := mux.NewRouter()

	router.HandleFunc("/", getAPIDescription).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000")
	http.ListenAndServe(":3000", router)
}
