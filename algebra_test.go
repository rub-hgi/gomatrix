package gomatrix

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixAdd(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		matrixB        *F2
		expectedMatrix *F2
	}{
		{
			description:    "success",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			matrixB:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(2)}),
			expectedMatrix: NewF2(2, 2).Set([]*big.Int{big.NewInt(0), big.NewInt(3)}),
		},
		{
			description:    "success",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			matrixB:        NewF2(2, 3).Set([]*big.Int{big.NewInt(2), big.NewInt(2)}),
			expectedMatrix: nil,
		},
	}

	for _, test := range tests {
		test.matrixA.AddMatrix(test.matrixB)

		assert.IsType(t, test.expectedMatrix, test.matrixA)

		if test.expectedMatrix == nil {
			continue
		}

		assert.Equal(t, true, test.expectedMatrix.IsEqual(test.matrixA))
	}
}
