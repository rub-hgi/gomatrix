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
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			expectedMatrix: NewF2(2, 2).Set([]*big.Int{big.NewInt(1), big.NewInt(3)}),
		},
	}

	for _, test := range tests {
		test.matrixA.GaussianElimination()

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}
