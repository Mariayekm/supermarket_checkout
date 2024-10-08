package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

func (s *shopConf) loadShopConf() *shopConf {

	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Fatalf("failed to read config file: %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, s)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return s
}

// Create a checkout instance based on the config file
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
	if _, ok := c.skus[SKU]; !ok {
		err = fmt.Errorf("product does not exit")
		return err
	}
	if _, ok := c.scannedProducts[SKU]; !ok {
		c.scannedProducts[SKU] = 1
	} else {
		c.scannedProducts[SKU] += 1
	}
	fmt.Println("Scanned product: ", SKU)
	return err
}

// GetTotalPrice returns the total cost of all the scanned products
func (c myCheckout) GetTotalPrice() (totalPrice int, err error) {
	if len(c.scannedProducts) < 1 {
		err = fmt.Errorf("cannot get total price: no products were scanned")
		return totalPrice, err
	}
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
	return err
}

func main() {
	fmt.Println("Running checkout program")
	var shop shopConf
	shop.loadShopConf()
	newCheckout, err := NewCheckout(shop)
	if err != nil {
		err = fmt.Errorf("failed to create new checkout: %v", err)
		log.Fatalf(err.Error())
	}

	fmt.Println("Start scanning products")
	reader := bufio.NewReader(os.Stdin)
	scannedInput, err := reader.ReadString('\n')
	processedInput := strings.ReplaceAll(scannedInput, " ", "")
	for _, product := range processedInput {
		if product == '\n' {
			break
		}
		newCheckout.Scan(string(product))
	}
	if total, err := newCheckout.GetTotalPrice(); err != nil {
		log.Fatalf(err.Error())
	} else {
		fmt.Printf("Total cost is %dp\n", total)
	}
}
