package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ICheckout interface {
	Scan(SKU string) (err error)

	GetTotalPrice() (totalPrice int, err error)
}

type myCheckout struct {
	scannedProducts map[string]int
	prices          []SKU
}

type SKU struct {
	sKUname         string
	normalPrice     int
	specialQuantity int
	specialPrice    float32
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

// GetTotalPrice returns the total cost of all the scanned products
func (c myCheckout) registerSKU(name string, price int, offer string) (err error) {
	newSKU := SKU{
		sKUname:     name,
		normalPrice: price,
	}
	if !strings.Contains(offer, "for") {
		err = fmt.Errorf("invalid offer")
		return err
	} else {
		tempStr := strings.ReplaceAll(offer, " ", "")
		processedOffer := strings.Split(tempStr, "for")
		if len(processedOffer) != 2 {
			err = fmt.Errorf("invalid offer")
			return err
		}
		if quantity, err := strconv.Atoi(processedOffer[0]); err != nil {
			err = fmt.Errorf("invalid offer")
			return err
		} else {
			newSKU.specialQuantity = quantity
		}
		if offerPrice, err := strconv.ParseFloat(processedOffer[1], 32); err != nil {
			err = fmt.Errorf("invalid offer")
			return err
		} else {
			newSKU.specialPrice = float32(offerPrice)
		}
	}
	fmt.Println("new sku: ", newSKU)
	return err
}

func main() {
	fmt.Println("Running checkout program")
	newCheckout := NewCheckout()
	newCheckout.Scan("A")
	newCheckout.Scan("A")
	newCheckout.Scan("B")
	total, _ := newCheckout.GetTotalPrice()
	fmt.Println("here is total ", total)
	newCheckout.registerSKU("A", 10, "3 for 2")
}
