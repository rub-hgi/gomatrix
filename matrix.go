package gomatrix

import (
	"math/big"
)

// F2 represents a matrix with entries that contains 0 or 1
//
// Each row consists of one big Int with arbitray size. Each column is one bit
// at 2**(column_index) of the big Int in each row.
type F2 struct {
	N    int
	M    int
	Rows []*big.Int
}

// NewF2 creates a new matrix in F_2
//
// @param int n The count of rows
// @param int m The count of columns
//
// @return *F2
func NewF2(n, m int) *F2 {
	// initialize rows array
	var rows []*big.Int

	// set each row to 0
	for i := 0; i < n; i++ {
		rows = append(
			rows,
			big.NewInt(0),
		)
	}

	// return the matrix
	return &F2{
		N:    n,
		M:    m,
		Rows: rows,
	}
}

// Set sets data from the data array
//
// @param []*big.Int data The data to insert into the matrix
//
// @return *F2|nil
func (f *F2) Set(data []*big.Int) *F2 {
	// if the size is different...
	if len(data) != f.N {
		// ...return an error
		return nil
	}

	// iterate through all given rows
	for i, datum := range data {
		// if the size is different...
		if datum.BitLen() > f.M {
			// ...return an error
			return nil
		}

		// set the depending row
		f.Rows[i] = new(big.Int).Set(datum)
	}

	// return success
	return f
}
