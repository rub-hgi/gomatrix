package gomatrix

import (
	"fmt"
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

// At returns the value at index i, j
//
// @param int i The row index
// @param int j The column index
//
// @return int, error|nil
func (f *F2) At(i, j int) (int, error) {
	// check if the indice are in the matrix
	if i >= f.N || j >= f.M {
		return -1, fmt.Errorf("Index out of bounds")
	}

	// create the bitmask for the value selection
	bitmask := big.NewInt(0).SetBit(big.NewInt(0), j, 1)

	// get the bit from the bitmask
	selectedBit := big.NewInt(0).And(f.Rows[i], bitmask)

	// compare the selected bit with 0. If it's 0, the result is 0. If it's not
	// zero, the result is 1 or -1. As the And only produces results greater or
	// equal 0, the result can
	result := selectedBit.Cmp(big.NewInt(0))

	// return the result
	return result, nil
}

// IsEqual checks the equality of the matrix objects
//
// This function compares the values of the matrix that is given with the matrix
// whose function is called.
//
// @param *F2 m The matrix to compare with
//
// @return bool
func (f *F2) IsEqual(m *F2) bool {
	// compare the sizes
	if f.N != m.N || f.M != m.M {
		return false
	}

	// iterate through the rows
	for i, row := range f.Rows {
		// if the rew is not equal...
		if row.Cmp(m.Rows[i]) != 0 {
			// ...return false
			return false
		}
	}

	// size and values are equal
	return true
}

// T transposes matrix f
//
// @return error
func (f *F2) T() error {
	// create the result matrix
	var resultRows []*big.Int

	// initialize the result matrix with 0
	for i := 0; i < f.M; i++ {
		resultRows = append(
			resultRows,
			big.NewInt(0),
		)
	}

	// iterate through the rows
	for i, row := range f.Rows {
		// transpose the row to a column
		transposeRowToColumn(row, resultRows, i)
	}

	// save the result matrix
	f.Rows = resultRows

	// save the dimensions
	f.N, f.M = f.M, f.N

	// return success
	return nil
}

// transposeRowToColumn transpose the row into the specified column
//
// This function creates a column at the columnIndex in rows by
// the given row.
//
// @param *big.Int   row         The row to transpose
// @param []*big.Int rows        The result matrix
// @param int        columnIndex The index of the column to create
func transposeRowToColumn(row *big.Int, rows []*big.Int, columnIndex int) {
	// iterate through the rows
	for i := range rows {
		// get the bit from the row
		bit := row.Bit(i)

		// set the bit to the column index
		rows[i].SetBit(rows[i], columnIndex, bit)
	}
}

// SetToIdentity sets the matrix to the identity matrix
//
// This function sets the identity matrix into f. If f is a non square matrix,
// the remaining rows/columns will be set to 0.
func (f *F2) SetToIdentity() {
	// iterate through the rows
	for i := 0; i < f.N; i++ {
		// create a row with 0 only
		f.Rows[i] = big.NewInt(0)

		// if the column/row counter is greater than the specified dimension...
		if i >= f.M {
			// ...skip the 1 value
			continue
		}

		// set the i'th bit for the identity matrix
		f.Rows[i].SetBit(f.Rows[i], i, 1)
	}
}

// SwapRows swaps the row at index i with the row at index j
//
// @param int i The index of the first row to swap
// @param int j The index of the second row to swap
//
// @return error
func (f *F2) SwapRows(i, j int) error {
	// check for input parameters
	if i >= f.N || j >= f.N || i < 0 || j < 0 {
		return fmt.Errorf("Index does not exist")
	}
	// swap the rows
	f.Rows[i], f.Rows[j] = f.Rows[j], f.Rows[i]

	// return success
	return nil
}

// SwapCols swaps the columns at index i with the row at index j
//
// @param int i The index of the first columns to swap
// @param int j The index of the second columns to swap
//
// @return error
func (f *F2) SwapCols(i, j int) error {
	// check for input parameters
	if i >= f.M || j >= f.M || i < 0 || j < 0 {
		return fmt.Errorf("Index does not exist")
	}

	// iterate through the rows
	for _, row := range f.Rows {
		// get the bit with the given index
		bitI := row.Bit(i)
		bitJ := row.Bit(j)

		// set the swapped bits
		row.SetBit(row, i, bitJ)
		row.SetBit(row, j, bitI)
	}

	// return success
	return nil
}
