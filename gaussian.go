// Package gomatrix Is a go package for scientific operations with matrices in F2.
package gomatrix

// GaussianElimination converts the matrix to an echelon form
//
// This function applies the gaussian elemination to the matrix in order to
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

				// substract the 1 from all other rows with the pivotBit
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