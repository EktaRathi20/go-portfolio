package main

import "fmt"

// Different payment gateways -
type razorpay struct{}

func (r *razorpay) pay(amount float32) {
	fmt.Println("Payment of", amount, "processed through Razorpay")
}

type stripe struct{}

func (s *stripe) pays(amount float32) {
	fmt.Println("Payment of", amount, "processed through Stripe")
}

// Interface
type paymentGateway interface {
	pay(amount float32)
}

// Payment struct
type payment struct {
	gateway paymentGateway
}

func (p *payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

func main() {
	razorpayGw := razorpay{}
	payment1 := payment{
		gateway: &razorpayGw,
	}
	payment1.makePayment(100.50)

	stripeGw := stripe{}
	payment2 := payment{
		gateway: &stripeGw,
	}
	payment2.makePayment(100.50)

}
