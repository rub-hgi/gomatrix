package gomatrix

import (
	"fmt"
	"math/big"
	"testing"

	_ "github.com/stretchr/testify/assert"
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
	}

	for _, test := range tests {
		test.matrixA.AddMatrix(test.matrixB)

		fmt.Printf("%v\n", test.matrixA)
	}
}
