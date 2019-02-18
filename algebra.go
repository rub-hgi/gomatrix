package gomatrix

import (
	"math/big"
)

// AddMatrix adds two matrices
//
// This function adds a matrix to the matrix object. The result will be saved
// in the object, whose AddMatrix method was called.
//
// @param *F2 m The matrix to add
//
// @return *F2|nil
func (f *F2) AddMatrix(m *F2) *F2 {
	// if the size is not equal...
	if f.N != m.N || f.M != m.M {
		// ...return an error
		return nil
	}

	// iterate through the rows
	for i := 0; i < f.N; i++ {
		//  xor each row with the relating row of the second matrix
		f.Rows[i].Xor(f.Rows[i], m.Rows[i])
	}

	// return the matrix
	return f
}

// MulMatrix multiplies matrix f with matrix m
//
// This functions multiplies matrix fxm. M could be a Nx1 matrix for a vector.
// If the matrices cannot be multiplied, nil is returned and f is not
// modified. If the multiplication was successful, the result is stored
// in f and returned.
//
// @param *F2 m The matrix that is used for the multiplication
//
// @return *F2
func (f *F2) MulMatrix(m *F2) *F2 {
	// if the dimensions do not fit for a multiplication...
	if f.N != m.M {
		// ...retrun an error
		return nil
	}

	// create the result matrix
	result := NewF2(f.N, m.M)

	// iterate through the rows of f
	for i, row := range f.Rows {
		// iterate through the columns of m
		for j := 0; j < f.M; j++ {
			// get the column from the second matrix
			col := m.GetCol(j)

			// multiply the vectors
			intermediateResult := big.NewInt(0).And(row, col)

			// sum up the values of the vectors
			resultBit := addBits(intermediateResult)

			// set the resulting bit to the result matrix
			result.Rows[i].SetBit(result.Rows[i], j, resultBit)
		}
	}

	// save the result matrix in f
	f.Rows = result.Rows

	// return the result
	return result
}

// addBits sums up all bits of a given number
//
// @param *big.Int number The number to process
//
// @return uint
func addBits(number *big.Int) uint {
	// get the bit length of the number
	bitLen := number.BitLen()

	// initialize the result
	result := uint(0)

	// iterate through the bits
	for i := 0; i < bitLen; i++ {
		result ^= number.Bit(i)
	}

	// return the result
	return result
}
