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

func TestPartialGaussianElimination(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		startRow       int
		startCol       int
		stopRow        int
		stopCol        int
		expectedMatrix *F2
	}{
		{
			matrixA: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(7),
				big.NewInt(4),
				big.NewInt(1),
			}),
			startRow: 0,
			stopRow:  2,
			startCol: 1,
			stopCol:  3,
			expectedMatrix: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(9),
				big.NewInt(1),
			}),
		},
		{
			matrixA: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(4),
				big.NewInt(10),
				big.NewInt(7),
				big.NewInt(1),
			}),
			startRow: 0,
			stopRow:  2,
			startCol: 1,
			stopCol:  3,
			expectedMatrix: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(9),
				big.NewInt(1),
			}),
		},
	}

	for _, test := range tests {
		test.matrixA.PartialGaussianElimination(
			test.startRow,
			test.startCol,
			test.stopRow,
			test.stopCol,
		)

		test.matrixA.PrettyPrint()

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}

func TestCheckGaussian(t *testing.T) {
	tests := []struct {
		description    string
		matrix         *F2
		startRow       int
		startCol       int
		n              int
		expectedResult bool
	}{
		{
			description: "3x3 identity matrix",
			matrix: NewF2(3, 3).Set([]*big.Int{
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(4),
			}),
			startRow:       0,
			startCol:       0,
			n:              3,
			expectedResult: true,
		},
		{
			description: "3x3 matrix",
			matrix: NewF2(3, 3).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(2),
				big.NewInt(4),
			}),
			startRow:       0,
			startCol:       0,
			n:              3,
			expectedResult: false,
		},
		{
			description: "3x3 matrix with lower right identity matrix",
			matrix: NewF2(3, 3).Set([]*big.Int{
				big.NewInt(2),
				big.NewInt(2),
				big.NewInt(4),
			}),
			startRow:       1,
			startCol:       1,
			n:              2,
			expectedResult: true,
		},
		{
			description: "3x3 matrix with upper left identity matrix",
			matrix: NewF2(3, 3).Set([]*big.Int{
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(7),
			}),
			startRow:       0,
			startCol:       0,
			n:              2,
			expectedResult: true,
		},
	}

	for _, test := range tests {
		test.matrix.PrettyPrint()
		result := test.matrix.CheckGaussian(
			test.startRow,
			test.startCol,
			test.n,
		)

		assert.Equalf(t, test.expectedResult, result, test.description)
	}
}
