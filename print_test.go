// Package gomatrix Is a go package for scientific operations with matrices in F2.
package gomatrix

import (
	"math/big"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	tests := []struct {
		matrix *F2
	}{
		{
			matrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(5), big.NewInt(3), big.NewInt(2)}),
		},
	}

	for _, test := range tests {
		test.matrix.PrettyPrint()
		test.matrix.PrintLaTex()
		test.matrix.PrintCSV()
	}
}
