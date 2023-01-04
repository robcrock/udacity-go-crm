package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

// Getting a single customer through a /customers/{id} path
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := mux.Vars(r)["id"]

	for k := range customers {
		if fmt.Sprint(k) == id {
			json.NewEncoder(w).Encode(customers[k])
		}
	}

}

// Getting all customers through a the /customers path
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)

}

// Creating a customer through a /customers path
func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var newEntry = customer{}

	// Read the HTTP request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Encode the request body
	json.Unmarshal(reqBody, &newEntry)

	append(customers, newEntry)

	json.NewEncoder(w).Encode(customers)

}

// Updating a customer through a /customers/{id} path
// Deleting a customer through a /customers/{id} path

func main() {
	seedCustomers()

	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", getCustomer).Methods("GET")

	fmt.Println("Server is starting on port 3000")
	http.ListenAndServe(":3000", router)
}
