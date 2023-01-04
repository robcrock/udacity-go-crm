package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Customer struct {
	ID uuid.UUID
	Name string
	Role string
	Email string
	Phone string
	Contacted bool
}

type CustomerDatabase struct {
	customerList []Customer
}

func main() {

	cd := CustomerDatabase {
		customerList: []Customer {
			{
					ID: uuid.New(),
					Name: "Robert",
					Role: "customer",
					Email: "r@g.com",
					Phone: "(555) 555-5555",
					Contacted: false,
			},
			{
					ID: uuid.New(),
					Name: "Robert",
					Role: "customer",
					Email: "r@g.com",
					Phone: "(555) 555-5555",
					Contacted: false,
			},
			{
					ID: uuid.New(),
					Name: "Robert",
					Role: "customer",
					Email: "r@g.com",
					Phone: "(555) 555-5555",
					Contacted: false,
			},
		},
	}

	fmt.Println(cd)
}
