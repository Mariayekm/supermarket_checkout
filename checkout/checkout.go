package main

import "fmt"

type ICheckout interface {
	Scan(SKU string) (err error)

	GetTotalPrice() (totalPrice int, err error)
}

type checkout struct {
	scannedProducts map[string]int
	total           int
}

func main() {
	fmt.Println("Running checkout program")
}
