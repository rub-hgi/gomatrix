// Package gomatrix Is a go package for scientific operations with matrices in F2.
package gomatrix

import (
	"fmt"
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
		{
			matrixA: NewF2(3, 4).Set([]*big.Int{
				big.NewInt(4),
				big.NewInt(10),
				big.NewInt(7),
			}),
			startRow: 0,
			stopRow:  2,
			startCol: 1,
			stopCol:  3,
			expectedMatrix: NewF2(3, 4).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(9),
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

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}

func TestPartialGaussianWithLinearChecking(t *testing.T) {
	tests := []struct {
		description    string
		matrix         *F2
		startRow       int
		startCol       int
		stopRow        int
		stopCol        int
		linearCheck    func(*F2, *F2, int, int, int, int, int) (*F2, error)
		expectedResult *F2
		expectedError  bool
	}{
		{
			description: "4x4 with one swap",
			matrix: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(13),
				big.NewInt(12),
				big.NewInt(14),
			}),
			startRow: 0,
			startCol: 1,
			stopRow:  2,
			stopCol:  3,
			linearCheck: func(f, permMatrix *F2, startRow, startCol, stopRow, stopCol, pivotBit int) (*F2, error) {
				// create a bitmask for the row check
				bitmask := big.NewInt(0).SetBit(big.NewInt(0), stopCol-startCol+1, 1)
				bitmask = bitmask.Sub(bitmask, big.NewInt(1))
				bitmask = bitmask.Lsh(bitmask, uint(startCol))

				foundValidRow := false

				// iterate through the rows
				for index, row := range f.Rows {
					// skip all rows, that are processed by the gaussian elimination
					if index >= startRow && index <= stopRow {
						continue
					}

					// get the bits to check
					bitsToCheck := big.NewInt(0).And(
						bitmask,
						row,
					)

					// if the bits are 0...
					if bitsToCheck.Cmp(big.NewInt(0)) == 0 {
						// ...skip the row
						continue
					}

					// swap the rows
					f.SwapRows(pivotBit-1, index)

					foundValidRow = true

					// exit the loop
					break
				}

				if !foundValidRow {
					return nil, fmt.Errorf("cannot resolv linear dependency")
				}

				for i := startCol; i < pivotBit; i++ {
					if f.Rows[pivotBit-startCol].Bit(i) == uint(0) {
						continue
					}

					fmt.Printf("%d xor %d\n", pivotBit-startCol, startRow+i-startCol)

					f.Rows[pivotBit-startCol].Xor(
						f.Rows[pivotBit-startCol],
						f.Rows[startRow+i-startCol],
					)
				}

				return permMatrix, nil
			},
			expectedResult: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(9),
				big.NewInt(1),
			}),
			expectedError: false,
		},
		{
			description: "4x4 with error",
			matrix: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(13),
				big.NewInt(12),
				big.NewInt(14),
			}),
			startRow: 0,
			startCol: 1,
			stopRow:  2,
			stopCol:  3,
			linearCheck: func(f, permMatrix *F2, startRow, startCol, stopRow, stopCol, pivotBit int) (*F2, error) {
				return nil, fmt.Errorf("testfoo")
			},
			expectedResult: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(9),
				big.NewInt(1),
			}),
			expectedError: true,
		},
		{
			description: "4x4 with error",
			matrix: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(13),
				big.NewInt(10),
				big.NewInt(12),
				big.NewInt(14),
			}),
			startRow: 0,
			startCol: 1,
			stopRow:  2,
			stopCol:  3,
			linearCheck: func(f, permMatrix *F2, startRow, startCol, stopRow, stopCol, pivotBit int) (*F2, error) {
				return nil, fmt.Errorf("testfoo")
			},
			expectedResult: NewF2(4, 4).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(9),
				big.NewInt(1),
			}),
			expectedError: true,
		},
	}

	for _, test := range tests {
		_, err := test.matrix.PartialGaussianWithLinearChecking(
			test.startRow,
			test.startCol,
			test.stopRow,
			test.stopCol,
			test.linearCheck,
		)

		test.matrix.PrettyPrint()

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if err != nil {
			continue
		}

		assert.Truef(t, test.expectedResult.IsEqual(test.matrix), test.description)
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
