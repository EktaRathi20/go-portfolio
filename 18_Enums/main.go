package main

import "fmt"

// type OrderStatus int

// const (
// 	Pending OrderStatus = iota
// 	Processing
// 	Shipped
// 	Delivered
// )

type OrderStatus string

const (
	Pending    OrderStatus = "Pending"
	Processing             = "Processing"
	Shipped                = "Shipped"
	Delivered              = "Delivered"
)

func changeStatus(newStatus OrderStatus) {
	fmt.Println("Order status changed to:", newStatus)
}
func main() {
	changeStatus((Pending))
}
