package gomatrix

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMatrix(t *testing.T) {
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
			description:    "invalid addition",
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

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}

func TestMulMatrix(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		matrixB        *F2
		expectedMatrix *F2
	}{
		{
			description:    "success",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			matrixB:        NewF2(2, 2).Set([]*big.Int{big.NewInt(1), big.NewInt(3)}),
			expectedMatrix: NewF2(2, 2).Set([]*big.Int{big.NewInt(3), big.NewInt(1)}),
		},
		{
			description:    "invalid multiplication",
			matrixA:        NewF2(2, 4).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			matrixB:        NewF2(2, 3).Set([]*big.Int{big.NewInt(1), big.NewInt(3)}),
			expectedMatrix: nil,
		},
	}

	for _, test := range tests {
		test.matrixA.MulMatrix(test.matrixB)

		assert.IsType(t, test.expectedMatrix, test.matrixA)

		if test.expectedMatrix == nil {
			continue
		}

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}

func TestAddBits(t *testing.T) {
	tests := []struct {
		description    string
		number         *big.Int
		expectedResult uint
	}{
		{
			description:    "two bits 1",
			number:         big.NewInt(3),
			expectedResult: 0,
		},
		{
			description:    "one bit 1",
			number:         big.NewInt(4),
			expectedResult: 1,
		},
	}

	for _, test := range tests {
		result := addBits(test.number)

		assert.Equal(t, test.expectedResult, result)
	}
}

func TestPartialXor(t *testing.T) {
	tests := []struct {
		description    string
		x              *big.Int
		y              *big.Int
		startCol       int
		stopCol        int
		expectedResult *big.Int
	}{
		{
			x:              big.NewInt(10),
			y:              big.NewInt(6),
			startCol:       0,
			stopCol:        4,
			expectedResult: big.NewInt(12),
		},
		{
			x:              big.NewInt(6),
			y:              big.NewInt(10),
			startCol:       1,
			stopCol:        2,
			expectedResult: big.NewInt(4),
		},
	}

	for _, test := range tests {
		result := PartialXor(
			test.x,
			test.y,
			test.startCol,
			test.stopCol,
		)

		assert.Equal(t, 0, test.expectedResult.Cmp(result))
	}
}
