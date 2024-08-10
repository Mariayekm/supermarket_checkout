package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetTotalPrice(t *testing.T) {
	testcases := []struct {
		description string
		input       int
		output      int
	}{
		{
			description: "Positive - first test",
			input:       1,
			output:      1,
		},
	}

	testCheckout := NewCheckout()
	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := testCheckout.GetTotalPrice()
			assert.Equal(t, result, tc.output)
			require.Nil(t, err)
		})
	}

}
