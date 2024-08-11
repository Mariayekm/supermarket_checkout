package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetTotalPrice(t *testing.T) {
	testcases := []struct {
		description    string
		inputSKUs      map[string]SKU
		productsBought map[string]int
		expectedOut    int
	}{
		{
			description: "Positive - first test",
			inputSKUs: map[string]SKU{
				"A": {10, nil, nil},
				"B": {10, makeIntPtr(5), makeIntPtr(40)},
				"C": {10, nil, nil},
				"D": {20, makeIntPtr(2), makeIntPtr(25)},
			},
			productsBought: map[string]int{
				"A": 3,
				"B": 5,
				"C": 1,
				"D": 5,
			},
			// (3*10 + 40 + 10 + (2*25)+20)
			expectedOut: 150,
		},
	}

	testCheckout := NewCheckout()
	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			testCheckout.skus = tc.inputSKUs
			testCheckout.scannedProducts = tc.productsBought
			result, err := testCheckout.GetTotalPrice()
			assert.Equal(t, tc.expectedOut, result)
			require.Nil(t, err)
		})
	}

}

func makeIntPtr(v int) *int {
	return &v
}
func makeStrPtr(v string) *string {
	return &v
}
