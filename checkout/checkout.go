package main

import "fmt"

type ICheckout interface {
	Scan(SKU string) (err error)

	GetTotalPrice() (totalPrice int, err error)
}

type myCheckout struct {
	scannedProducts map[string]int
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

// Scan updates the number of products that have been scanned
func (c myCheckout) Scan(sKU string) (err error) {
	if _, ok := c.scannedProducts[sKU]; !ok {
		c.scannedProducts[sKU] = 1
	} else {
		c.scannedProducts[sKU] += 1
	}
	fmt.Println("scanned product")
	return err
}

// GetTotalPrice returns the total cost of all the scanned products
func (c myCheckout) GetTotalPrice() (totalPrice int, err error) {
	for _, k := range c.scannedProducts {
		totalPrice += k
	}
	return totalPrice, err
}

func main() {
	fmt.Println("Running checkout program")
	newCheckout := NewCheckout()
	newCheckout.Scan("A")
	newCheckout.Scan("A")
	newCheckout.Scan("B")
	total, _ := newCheckout.GetTotalPrice()
	fmt.Println("here is total ", total)

}
