// Package gomatrix Is a go package for scientific operations with matrices in F2.
package gomatrix

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGaussianElimination(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		expectedMatrix *F2
	}{
		{
			matrixA:        NewF2(3, 3).Set([]*big.Int{big.NewInt(5), big.NewInt(3), big.NewInt(2)}),
			expectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(4)}),
		},
		{
			matrixA:        NewF2(3, 3).Set([]*big.Int{big.NewInt(2), big.NewInt(5), big.NewInt(3)}),
			expectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(4)}),
		},
	}

	for _, test := range tests {
		test.matrixA.GaussianElimination()

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}
