package main

import "fmt"

type customer struct {
	id    int
	name  string
	email string
}
type order struct {
	id        int
	amount    float64
	status    string
	createdAt string
	customer  customer // Embedding another struct
}

func main() {

	order1 := order{
		id:        1,
		amount:    100.50,
		status:    "Pending",
		createdAt: "2023-10-01",
		customer: customer{
			id:    1,
			name:  "John Doe",
			email: "john@gmail.com",
		},
	}

	fmt.Println("Order ID:", order1)
}
