package main

import (
	"fmt"
	"time"
)

// Structs are composite data types that group together variables (fields) under a single name.

type order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time
}

// struct don't have it's constructor like classes in other languages
// but you can create a function that initializes the struct
// and returns an instance of the struct
func newOrder(id int, amount float64, status string) *order {
	order := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}

	return &order
}

// use pointer when you want to modify the struct else if you just want to read the struct you can use it without pointer
// struct automatically de-refrences the pointer when accessing fields
func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
	fmt.Println("Order status changed to:", o.status)
}

func (o order) getAmout() float64 {
	return o.amount
}

// Another way to create a struct is to use an anonymous struct
func createStruct() {
	newStruct := struct {
		name string
		age  int
	}{
		name: "John Doe",
		age:  30,
	}

	fmt.Println("New Struct:", newStruct)
}

func main() {
	// if you don't set any fields in the struct, it will be set to zero value
	order := order{
		id:     1,
		amount: 100.50,
		status: "Pending",
		// createdAt: time.Now(),
	}

	order.createdAt = time.Now()
	order.changeStatus("Shipped")
	fmt.Println("Order Amount:", order.getAmout())
	fmt.Println("Order:", order)

	order2 := newOrder(2, 200.75, "Processing")

	fmt.Println("Order 2:", order2)

	createStruct()
}
