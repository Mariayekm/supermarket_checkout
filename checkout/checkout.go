package main

import "fmt"

type ICheckout interface {
	Scan(SKU string) (err error)

	GetTotalPrice() (totalPrice int, err error)
}

type myCheckout struct {
	scannedProducts map[string]int
	total           int
}

// Make sure myCheckout implements ICheckout
var _ ICheckout = myCheckout{}

// Create a checkout instance
func NewCheckout() myCheckout {
	return myCheckout{}
}

func (c myCheckout) Scan(SKU string) (err error) {
	return err
}

func (c myCheckout) GetTotalPrice() (totalPrice int, err error) {
	return totalPrice, err
}

func main() {
	fmt.Println("Running checkout program")
	newCheckout := NewCheckout()
	total, _ := newCheckout.GetTotalPrice()
	fmt.Println("here is total ", total)
}
