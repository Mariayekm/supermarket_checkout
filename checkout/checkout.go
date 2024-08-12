package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type ICheckout interface {
	Scan(SKU string) (err error)

	GetTotalPrice() (totalPrice int, err error)
}

type myCheckout struct {
	scannedProducts map[string]int
	skus            map[string]SKU
}

type SKU struct {
	unitPrice       int
	specialQuantity *int
	specialPrice    *int
}

type shopConf struct {
	Items []struct {
		SKUName      string  `yaml:"SKU"`
		UnitPrice    int     `yaml:"unitPrice"`
		SpecialPrice *string `yaml:"specialPrice"`
	} `yaml:"items"`
}

// Make sure myCheckout implements ICheckout
var _ ICheckout = myCheckout{}

// Create a checkout instance
func NewCheckout(s shopConf) (myCheckout, error) {
	newCheckout := myCheckout{}

	emptyProducts := make(map[string]int)
	newCheckout.scannedProducts = emptyProducts

	emptyInventory := make(map[string]SKU)
	newCheckout.skus = emptyInventory

	for _, item := range s.Items {
		if err := newCheckout.registerSKU(item.SKUName, item.UnitPrice, item.SpecialPrice); err != nil {
			return newCheckout, err
		}
	}
	return newCheckout, nil
}

// Scan updates the number of products that have been scanned
func (c myCheckout) Scan(SKU string) (err error) {
	// check sku is in inventory?
	if _, ok := c.scannedProducts[SKU]; !ok {
		c.scannedProducts[SKU] = 1
	} else {
		c.scannedProducts[SKU] += 1
	}
	fmt.Println("scanned product")
	return err
}

// GetTotalPrice returns the total cost of all the scanned products
func (c myCheckout) GetTotalPrice() (totalPrice int, err error) {
	for productName, productQuantity := range c.scannedProducts {
		sKU := c.skus[productName]
		if sKU.specialQuantity == nil {
			totalPrice += (sKU.unitPrice * productQuantity)
		} else {
			discountedItems := productQuantity / (*sKU.specialQuantity)
			fullPriceItems := productQuantity % (*sKU.specialQuantity)
			totalPrice += ((int(*sKU.specialPrice) * discountedItems) + (sKU.unitPrice * fullPriceItems))
		}
	}
	return totalPrice, err
}

func (s *shopConf) loadShopConf() *shopConf {

	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, s)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return s
}

// registerSKU adds a new SKU to the checkout.
// The offer parameter is expected in the format "x for y"
// otherwise an error is returned.
func (c myCheckout) registerSKU(name string, price int, offer *string) (err error) {
	newSKU := SKU{
		unitPrice: price,
	}
	if offer != nil {
		if !strings.Contains(*offer, "for") {
			err = fmt.Errorf("invalid offer")
			return err
		} else {
			tempStr := strings.ReplaceAll(*offer, " ", "")
			processedOffer := strings.Split(tempStr, "for")
			if len(processedOffer) != 2 {
				err = fmt.Errorf("invalid offer")
				return err
			}
			if quantity, err := strconv.Atoi(processedOffer[0]); err != nil {
				err = fmt.Errorf("invalid offer")
				return err
			} else {
				newSKU.specialQuantity = &quantity
			}
			if offerPrice, err := strconv.Atoi(processedOffer[1]); err != nil {
				err = fmt.Errorf("invalid offer")
				return err
			} else {
				newSKU.specialPrice = &offerPrice
			}
		}
	}

	c.skus[name] = newSKU
	// fmt.Println("new sku: ", c.skus)
	return err
}

func main() {
	fmt.Println("Running checkout program")
	var shop shopConf
	shop.loadShopConf()
	newCheckout, err := NewCheckout(shop)
	if err != nil {
		err = fmt.Errorf("failed to create new checkout")
		panic(err)
	}
	newCheckout.Scan("A")
	newCheckout.Scan("A")
	newCheckout.Scan("A")
	newCheckout.Scan("B")
	total, _ := newCheckout.GetTotalPrice()
	fmt.Println("here is total ", total)
}
