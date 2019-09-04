package resolver

import (
	"fmt"

	"git.noc.ruhr-uni-bochum.de/danieljankowski/gomatrix"
)

// LinearDependenciesInGauss tries to resolve linear dependencies in the gaussian
// elimination.
//
// This function is used in order to try to resolve linear dependencies while
// using the function PartialGaussianWithLinearChecking as linearCheck-function.
func LinearDependenciesInGauss(
	f *gomatrix.F2,
	gaussMatrix *gomatrix.F2,
	permutationMatrix *gomatrix.F2,
	startRow int,
	startCol int,
	stopRow int,
	stopCol int,
	pivotBit int,
) (*gomatrix.F2, *gomatrix.F2, error) {
	// resolve the linear dependency
	permutationMatrix, err := resolveWithOptimizedAlgorithm(
		f,
		permutationMatrix,
		startRow,
		startCol,
		stopRow,
		stopCol,
		pivotBit,
	)

	// if an error occured...
	if err != nil {
		// ...return it
		return nil, nil, err
	}

	// apply the previous operations on the new row, with iterating through
	// the columns 'til the pivot bit is reached
	for i := startCol; i < pivotBit; i++ {
		// if the column is zero...
		if f.Rows[pivotBit-startCol].Bit(i) == uint(0) {
			// ...skip to the next column
			continue
		}

		// remove the 1 with a xor operation with the relating row
		f.Rows[startRow+pivotBit-startCol].Xor(
			f.Rows[startRow+pivotBit-startCol],
			f.Rows[startRow+i-startCol],
		)

		gaussMatrix.Rows[startRow+pivotBit-startCol].Xor(
			gaussMatrix.Rows[startRow+pivotBit-startCol],
			gaussMatrix.Rows[startRow+i-startCol],
		)
	}

	// return success
	return gaussMatrix, permutationMatrix, nil
}

// resolveWithOptimizedAlgorithm tries to resolve the dependency with finding
// an appropriate value that can be swapped right into the correct position
// without destroying the already processed rows and columns.
func resolveWithOptimizedAlgorithm(
	f *gomatrix.F2,
	permutationMatrix *gomatrix.F2,
	startRow int,
	startCol int,
	stopRow int,
	stopCol int,
	pivotBit int,
) (*gomatrix.F2, error) {
	// iterate through the rows
	for rowIndex := 0; rowIndex < f.N; rowIndex++ {
		// if the rowindex points on to the already processed rows...
		if rowIndex >= startRow && rowIndex <= startRow+pivotBit-startCol {
			// ...skip it
			continue
		}

		// iterate through the columns
		for colIndex := 0; colIndex < f.M; colIndex++ {
			// if the colindex points on to the already processed rows...
			if colIndex >= startCol && colIndex < pivotBit {
				// ...skip it
				continue
			}

			// get the value at the current index
			bit, err := f.At(rowIndex, colIndex)

			// if an error occured or the bit is 0...
			if err != nil || bit == 0 {
				// ...skip it
				continue
			}

			// swap the value into the right place
			f.SwapRows(rowIndex, startRow+pivotBit-startCol)
			f.SwapCols(colIndex, pivotBit)

			// swap the rows in the permutation matrix
			permutationMatrix.SwapRows(rowIndex, startRow+pivotBit-startCol)
			permutationMatrix.SwapCols(colIndex, pivotBit)

			// return success
			return permutationMatrix, nil
		}
	}

	return nil, fmt.Errorf("cannot resolve dependency")
}
