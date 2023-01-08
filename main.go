package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customer struct {
	ID int
	Name string
	Role string
	Email string
	Phone string
	Contacted bool
}

var customers []customer

func seedCustomers() {
	c1 := customer{
			ID: 0,
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}
	c2 := customer{
			ID: 1,
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}
	c3 := customer{
			ID: 2,
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}

	customers = append(customers, c1, c2, c3)
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

	for k := range customers {
		if fmt.Sprint(k) == id {
			customerFound = true
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customers[k])
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

	json.NewEncoder(w).Encode(customers)

}

// Creating a customer through a /customers path
func addCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate Content-Type in the response header
	w.Header().Set("Content-Type", "application/json")

	// Create (but not yet assign values to) for the new entry
	c := customer{}
	customerAlreadyExists := false


	// Read the HTTP request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Encode the request body
	json.Unmarshal(reqBody, &c)

	for _, v := range customers {
		if v.ID == c.ID {
			customerAlreadyExists = true
			w.WriteHeader(http.StatusConflict)
		}
	}

	if !customerAlreadyExists {
		customers = append(customers, c)
		w.WriteHeader(http.StatusCreated)
	}

	// Regardless of successful resource creation or not, return the current state of the "dictionary" map
	json.NewEncoder(w).Encode(customers)

}

// Updating a customer through a /customers/{id} path
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate Content-Type in the response header
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	customerFound := false

	// Create (but not yet assign values to) for the new entry
	c := customer{}

	// Read the HTTP request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Encode the request body
	json.Unmarshal(reqBody, &c)

	for k := range customers {
		if k == id {
			customerFound = true
			w.WriteHeader(http.StatusOK)
			customers[k] = c
		}
	}

	if !customerFound {
		w.WriteHeader(http.StatusNotFound)
	}

	// Regardless of successful resource creation or not, return the current state of the "dictionary" map
	json.NewEncoder(w).Encode(customers)

}

// Deleting a customer through a /customers/{id} path
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate Content-Type in the response header
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	customerFound := false
	nullReplacement := customer{
			ID: 0,
			Name: "",
			Role: "",
			Email: "",
			Phone: "",
			Contacted: false,
    }


	for _, v := range customers {
		if v.ID == id {
			customerFound = true
			w.WriteHeader(http.StatusOK)
			// Remove the element at index i from a.
			customers[id] = customers[len(customers)-1] // Copy last element to index i.
			customers[len(customers)-1] = nullReplacement // Erase last element (write zero value).
			customers = customers[:len(customers)-1] // Truncate slice.
			json.NewEncoder(w).Encode(customers)
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
	router.HandleFunc("/customer/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000")
	http.ListenAndServe(":3000", router)
}
