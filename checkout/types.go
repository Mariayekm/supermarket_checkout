package main

type ICheckout interface {
	Scan(SKU string) (err error)

	GetTotalPrice() (totalPrice int, err error)
}

type myCheckout struct {
	scannedProducts map[string]int
	skus            map[string]SKU
}

// Make sure myCheckout implements ICheckout
var _ ICheckout = myCheckout{}

type SKU struct {
	unitPrice       int
	specialQuantity *int
	specialPrice    *int
}

// shopConf is used for loading the inventory for the shop
type shopConf struct {
	Items []struct {
		SKUName      string  `yaml:"SKU"`
		UnitPrice    int     `yaml:"unitPrice"`
		SpecialPrice *string `yaml:"specialPrice"`
	} `yaml:"items"`
}
