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
	newCheckout := myCheckout{}
	emptyProducts := make(map[string]int)
	newCheckout.scannedProducts = emptyProducts
	return newCheckout
}

func (c myCheckout) Scan(SKU string) (err error) {
	if _, ok := c.scannedProducts[SKU]; !ok {
		c.scannedProducts[SKU] = 1
	} else {
		c.scannedProducts[SKU] += 1
	}
	fmt.Println("scanned product")
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
	newCheckout.Scan("A")
	newCheckout.Scan("A")

}
