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
		return fmt.Errorf("cannot resolve linear dependency")
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

	return nil
}
