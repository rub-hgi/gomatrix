package gomatrix

import (
	"math/big"
)

// GaussianElimination converts the matrix to an echelon form
//
// This function applies the gaussian elimination to the matrix in order to
// create an echelon form.
func (f *F2) GaussianElimination() {
	// iterate through all possible pivot bits
	for pivotBit := 0; pivotBit < f.M; pivotBit++ {
		// iterate through the rows
		for rowCounter := pivotBit; rowCounter < f.N; rowCounter++ {
			// if the pivotbit of this row is 0...
			if f.Rows[rowCounter].Bit(pivotBit) == uint(0) {
				// ...check the next row
				continue
			}

			// if the row with a valid pivot bit is not the first row...
			if pivotBit != rowCounter {
				// ...swap it with first one
				f.SwapRows(pivotBit, rowCounter)
			}

			// iterate through all other rows except the first one
			for rr := pivotBit + 1; rr < f.N; rr++ {
				if f.Rows[rr].Bit(pivotBit) == uint(0) {
					continue
				}

				// subtract the 1 from all other rows with the pivotBit
				f.Rows[rr].Xor(f.Rows[rr], f.Rows[pivotBit])
			}
		}
	}

	// do the same thing backwards to get the identity matrix
	f.diagonalize()
}

// diagonalize Diagonalizes the matrix after creating the triangular matrix
//
// This function removes the top right 1 entries in a matrix with the
// echelon form.
func (f *F2) diagonalize() {
	// iterate backwards through the pivot bits
	for pivotBit := f.M - 1; pivotBit >= 0; pivotBit-- {
		// choose each row from the top row to the one with the pivot bit
		for rowCounter := 0; rowCounter < pivotBit; rowCounter++ {
			// if the bit in the same position at the other row is 0...
			if f.Rows[rowCounter].Bit(pivotBit) == uint(0) {
				// ...continue to the next row
				continue
			}

			// eliminate the 1
			f.Rows[rowCounter].Xor(f.Rows[rowCounter], f.Rows[pivotBit])
		}
	}
}

// PartialGaussianElimination performs a gaussian elimination on a part of the matrix
func (f *F2) PartialGaussianElimination(startRow, startCol, stopRow, stopCol int) {
	// iterate through all possible pivot bits
	for pivotBit := startCol; pivotBit <= stopCol; pivotBit++ {
		// iterate through the rows
		for rowCounter := startRow + pivotBit - startCol; rowCounter <= stopRow; rowCounter++ {
			// if the pivotbit of this row is 0...
			if f.Rows[rowCounter].Bit(pivotBit) == uint(0) {
				// ...check the next row
				continue
			}

			// if the row with a valid pivot bit is not the first row...
			if startRow+pivotBit-startCol != rowCounter {
				// ...swap it with first one
				f.SwapRows(startRow+pivotBit-startCol, rowCounter)
			}

			// iterate through all other rows except the first one
			for rr := startRow + pivotBit - startCol + 1; rr <= stopRow; rr++ {
				if f.Rows[rr].Bit(pivotBit) == uint(0) {
					continue
				}

				// subtract the 1 from all other rows with the pivotBit
				f.Rows[rr].Xor(
					f.Rows[rr],
					f.Rows[startRow+pivotBit-startCol],
				)
			}

			break
		}
	}

	// do the same thing backwards to get the identity matrix
	f.partialDiagonalize(startRow, startCol, stopRow, stopCol, nil)
}

func (f *F2) partialDiagonalize(startRow, startCol, stopRow, stopCol int, gaussMatrix *F2) *F2 {
	// iterate backwards through the pivot bits
	for pivotBit := stopCol; pivotBit >= startCol; pivotBit-- {
		// choose each row from the top row to the one with the pivot bit
		for rowCounter := startRow; rowCounter < stopRow; rowCounter++ {
			// prevent xor with the row itself
			if rowCounter == startRow+pivotBit-startCol {
				continue
			}

			// if the bit in the same position at the other row is 0...
			if f.Rows[rowCounter].Bit(pivotBit) == uint(0) {
				// ...continue to the next row
				continue
			}

			// eliminate the 1
			f.Rows[rowCounter].Xor(
				f.Rows[rowCounter],
				f.Rows[startRow+pivotBit-startCol],
			)

			if gaussMatrix == nil {
				continue
			}

			// eliminate the 1
			gaussMatrix.Rows[rowCounter].Xor(
				gaussMatrix.Rows[rowCounter],
				gaussMatrix.Rows[startRow+pivotBit-startCol],
			)
		}
	}

	return gaussMatrix
}

// PartialGaussianWithLinearChecking performs a partial gaussian elimination
//
// This function performs a gaussian elimination on the matrix and calls the
// check callback after each iteration in order to verify that linear
// dependencies in the code could be resolved easily. The function returns
// the permutation matrix in addition to the error. For the linearCheck
// callback take a look at the resolver package.
func (f *F2) PartialGaussianWithLinearChecking(
	startRow int,
	startCol int,
	stopRow int,
	stopCol int,
	linearCheck func(*F2, *F2, *F2, int, int, int, int, int) (*F2, *F2, error),
) (*F2, *F2, error) {
	// initialize the permutation matrix
	gaussMatrix := NewF2(f.N, f.N).SetToIdentity()
	permutationMatrix := NewF2(f.M, f.M).SetToIdentity()

	// initialize the error vector
	var err error

	// iterate through all possible pivot bits
	for pivotBit := startCol; pivotBit <= stopCol; pivotBit++ {
		// intialize the pivotbit indicator
		foundPivotBit := false

		// iterate through the rows
		for rowCounter := startRow + pivotBit - startCol; rowCounter <= stopRow; rowCounter++ {
			// if the pivotbit of this row is 0...
			if f.Rows[rowCounter].Bit(pivotBit) == uint(0) {
				// ...check the next row
				continue
			}

			// if the row with a valid pivot bit is not the first row...
			if startRow+pivotBit-startCol != rowCounter {
				// ...swap it with first one
				f.SwapRows(startRow+pivotBit-startCol, rowCounter)
				gaussMatrix.SwapRows(startRow+pivotBit-startCol, rowCounter)
			}

			// iterate through all other rows except the first one
			for rr := startRow + pivotBit - startCol + 1; rr <= stopRow; rr++ {
				if f.Rows[rr].Bit(pivotBit) == uint(0) {
					continue
				}

				// subtract the 1 from all other rows with the pivotBit
				f.Rows[rr].Xor(
					f.Rows[rr],
					f.Rows[startRow+pivotBit-startCol],
				)
				gaussMatrix.Rows[rr].Xor(
					gaussMatrix.Rows[rr],
					gaussMatrix.Rows[startRow+pivotBit-startCol],
				)
			}

			// indicate the pivotbit is found
			foundPivotBit = true

			// break out of the loop
			break
		}

		// if a pivot bit was found...
		if foundPivotBit {
			// ...skip to the next row
			continue
		}

		// detect linear dependencies and try to resolve them
		gaussMatrix, permutationMatrix, err = linearCheck(
			f,
			gaussMatrix,
			permutationMatrix,
			startRow,
			startCol,
			stopRow,
			stopCol,
			pivotBit,
		)

		// check the error
		if err != nil {
			return nil, nil, err
		}

		// process the same row again
		pivotBit--
	}

	// do the same thing backwards to get the identity matrix
	gaussMatrix = f.partialDiagonalize(startRow, startCol, stopRow, stopCol, gaussMatrix)

	return gaussMatrix, permutationMatrix, nil
}

// CheckGaussian checks if the given range in the matrix is the identity matrix
//
// @param int startRow The row where the check starts
// @param int startCol The column where the check starts
// @param int n        The size of the submatrix to check
//
// @return bool
func (f *F2) CheckGaussian(startRow, startCol, n int) bool {
	counter := 0

	// create the bitmask for the bits to check
	bitmask := big.NewInt(0).SetBit(
		big.NewInt(0),
		n,
		1,
	)

	bitmask = bitmask.Sub(bitmask, big.NewInt(1))
	bitmask = bitmask.Lsh(bitmask, uint(startCol))

	// iterate through the rows
	for i := startRow; i < startRow+n; i++ {
		// get the row
		row := f.Rows[i]

		// calculate the expected result if it is in echelon form
		expectedBits := big.NewInt(0).SetBit(
			big.NewInt(0),
			startCol+counter,
			1,
		)

		// get the bits to check
		bitsToCheck := big.NewInt(0).And(
			row,
			bitmask,
		)

		// xor the bits to check with the bitmask
		shouldBeZero := big.NewInt(0).Xor(
			bitsToCheck,
			expectedBits,
		)

		// if the xor'ed result is not zero...
		if shouldBeZero.Cmp(big.NewInt(0)) != 0 {
			// ...the check failed
			return false
		}

		// increase the counter
		counter++
	}

	return true
}
