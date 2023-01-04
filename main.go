package main

import (
	"fmt"

	"github.com/google/uuid"
)

type customer struct {
	ID uuid.UUID
	Name string
	Role string
	Email string
	Phone string
	Contacted bool
}

var customers []customer

func seedCustomers() {

	c1 := customer{
			ID: uuid.New(),
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}
	c2 := customer{
			ID: uuid.New(),
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}
	c3 := customer{
			ID: uuid.New(),
			Name: "Robert",
			Role: "customer",
			Email: "r@g.com",
			Phone: "(555) 555-5555",
			Contacted: false,
		}

	customers = append(customers, c1, c2, c3)

}

func main() {
	seedCustomers()

	fmt.Println(customers)
}
