package interfaces

import "fmt"

type paymentr interface {
	Pay(amount float64) error
	Refund(amount float64) error
}

type payment struct {
	gateway paymentr
}

func (p *payment) Pay(amount float64) error {
	return p.gateway.Pay(amount)
}

func (p *payment) Refund(amount float64) error {
	return p.gateway.Refund(amount)
}
type razorpay struct {}

func (r *razorpay) Pay(amount float64) error {
	fmt.Println("payment succesful of amount:-", amount , "from razorpay")
	return nil
}

func (r *razorpay) Refund(amount float64) error {
	fmt.Println("refund succesful of amount:-", amount)
	return nil
}

type stripe struct {}

func (s *stripe) Pay(amount float64) error {
	fmt.Println("payment succesful of amount:-", amount, "from stripe")
	return nil
}

func (s *stripe) Refund(amount float64) error {
	fmt.Println("refund succesful of amount:-", amount)
	return nil
}

func InterfaceExample() {
	newPayment := &payment{gateway: &stripe{}}
	newPayment2 := &payment{gateway: &razorpay{}}
	newPayment.Pay(122)
	newPayment2.Pay(122)
}
