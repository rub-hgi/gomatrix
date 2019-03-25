package resolver

import (
	"fmt"
	"math/big"

	"git.noc.ruhr-uni-bochum.de/danieljankowski/gomatrix"
)

// LinearDependenciesInGauss tries to resolve linear dependencies in the gaussian
// elimination.
//
// This function is used in order to try to resolve linear dependencies while
// using the function PartialGaussianWithLinearChecking as linearCheck-function.
func LinearDependenciesInGauss(
	f *gomatrix.F2,
	startRow int,
	startCol int,
	stopRow int,
	stopCol int,
	pivotBit int,
) error {
	// try to resolve the dependency with a row swap
	err := resolveWithRowSwap(f, startRow, startCol, stopRow, stopCol, pivotBit)

	// if the dependency is resolved...
	if err == nil {
		fmt.Printf("early return\n")
		// ...return success
		return nil
	}

	// try to resolve the dependency with a col swap
	err = resolveWithColSwap(f, startRow, startCol, stopRow, stopCol, pivotBit)

	// return success
	return err
}

// resolveWithRowSwap tries to resolve the linear dependency by swapping the row
// with another one
func resolveWithRowSwap(
	f *gomatrix.F2,
	startRow int,
	startCol int,
	stopRow int,
	stopCol int,
	pivotBit int,
) error {
	// create a bitmask for the row check
	bitmask := big.NewInt(0).SetBit(big.NewInt(0), stopCol-startCol+1, 1)
	bitmask = bitmask.Sub(bitmask, big.NewInt(1))
	bitmask = bitmask.Lsh(bitmask, uint(startCol))

	// initialize the indicator if a valid row is found
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

		// a valid row was swapped into the right place
		foundValidRow = true

		// exit the loop
		break
	}

	// if no valid row was found...
	if !foundValidRow {
		// ...return an error
		return fmt.Errorf("cannot resolve linear dependency")
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
		f.Rows[pivotBit-startCol].Xor(
			f.Rows[pivotBit-startCol],
			f.Rows[startRow+i-startCol],
		)
	}

	// return success
	return nil
}

// resolveWithRowSwap tries to resolve the linear dependency by swapping the column
// with another one
func resolveWithColSwap(
	f *gomatrix.F2,
	startRow int,
	startCol int,
	stopRow int,
	stopCol int,
	pivotBit int,
) error {
	// get the current row as the pivot bit is indexed over the whole matrix
	rowIndex := pivotBit - startCol

	foundValidCol := false

	// iterate through the colummns
	for i := 0; i < f.M; i++ {
		// skip the columns that are used in the gaussian elimination
		if i >= startCol && i <= stopCol {
			continue
		}

		// if the bit is 0...
		if f.Rows[rowIndex].Bit(i) == uint(0) {
			// ...the row is skipped
			continue
		}

		// swap the columns
		f.SwapCols(i, pivotBit)

		// indicate that a correct column was swapped right into the place
		foundValidCol = true

		// break out of the loop
		break
	}

	// if no valid column was found...
	if !foundValidCol {
		// ...return an error
		return fmt.Errorf("cannot resolve linear dependency")
	}

	// return success
	return nil
}
