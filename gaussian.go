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
		for rowCounter := 0; rowCounter < f.N; rowCounter++ {
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
				// substract the 1 from all other rows with the pivotBit
				f.Rows[rr].Xor(f.Rows[rr], f.Rows[pivotBit])
			}
		}
	}
}
