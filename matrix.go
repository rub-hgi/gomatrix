// Package gomatrix Is a go package for scientific operations with matrices in F2.
package gomatrix

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// F2 represents a matrix with entries that contains 0 or 1
//
// Each row consists of one big Int with arbitrary size. Each column is one bit
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
func (f *F2) T() *F2 {
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

	return f
}

// PartialT partially transpose the matrix
//
// This function partially transposes a matrix. The submatrix that is
// transposed need to be a square matrix.
//
// @param int startRow The row to start
// @param int startCol The column to start
// @param int n        The size of the submatrix
//
// @return error
func (f *F2) PartialT(startRow, startCol, n int) error {
	// verify the given parameters
	if startRow+n > f.N || startCol+n > f.M {
		return fmt.Errorf("Cannot partially transpose a non square matrix")
	}

	// get the submatrix to transpose
	subMatrix := f.GetSubMatrix(
		startRow,
		startCol,
		startRow+n,
		startCol+n,
	)

	// transpose the submatrix
	subMatrix.T()

	// set the transposed submatrix into f
	f.SetSubMatrix(
		subMatrix,
		startRow,
		startCol,
	)

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
func (f *F2) SetToIdentity() *F2 {
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

	return f
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

// PermuteCols permutes the columns of the matrix randomly
//
// This function swaps columns randomly. The swap operation will be repeated
// on every column. After swapping the columns, the permutation
// matrix will be returned.
//
// @return *F2
func (f *F2) PermuteCols() *F2 {
	// initialize the permuation matrix
	permutationMatrix := NewF2(f.N, f.M)

	// set the permutation matrix to the identity matrix of f
	permutationMatrix.SetToIdentity()

	// repeat the swaps
	for i := 0; i < f.M; i++ {
		// get a random column destination index
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(f.M)))

		// check if the column is swapped or not
		if int64(i) == j.Int64() {
			continue
		}

		// swap the columns in f
		f.SwapCols(i, int(j.Int64()))

		// swap the columns in the permutation matrix
		permutationMatrix.SwapCols(i, int(j.Int64()))
	}

	// return the permutation matrix
	return permutationMatrix
}

// GetCol returns the column at index i
//
// This function returns the column as big.Int after the index is verified.
// If an invalid index is used, the function returns nil.
//
// @param int i The index for the column
//
// @return *big.Int
func (f *F2) GetCol(i int) *big.Int {
	// check for input parameters
	if i < 0 || i >= f.M {
		// return nil
		return nil
	}

	// initialize the output big.Int
	output := big.NewInt(0)

	// iterate through the rows
	for j, row := range f.Rows {
		// the the corresponding bit
		output.SetBit(output, j, row.Bit(i))
	}

	// return the result
	return output
}

// GetSubMatrix gets the submatrix with boundaries included
//
// @param int startRow The first row to include
// @param int startCol The first column to include
// @param int stopRow  The last row to include
// @param int stopCol  The last column to include
func (f *F2) GetSubMatrix(startRow, startCol, stopRow, stopCol int) *F2 {
	// create the output matrix
	output := NewF2(stopRow-startRow, stopCol-startCol)

	// calculate the bitlength
	bitLength := stopCol - startCol - 1

	// create the bitmask
	bitMask := big.NewInt(0).Exp(
		big.NewInt(2), big.NewInt(int64(bitLength+1)), nil,
	)

	// decrease the bitmask by one
	bitMask.Sub(bitMask, big.NewInt(1))

	// shift the bitmask to the correct position
	bitMask.Lsh(bitMask, uint(startCol))

	// initialize the rows for the output
	var rows []*big.Int

	// iterate through the given rows
	for i := startRow; i < stopRow; i++ {
		// get the bits with the bitmask
		outputRow := big.NewInt(0).And(f.Rows[i], bitMask)

		// shift the completely to the right
		outputRow.Rsh(outputRow, uint(startCol))

		// save the row
		rows = append(
			rows,
			outputRow,
		)
	}

	// return the matrix with the rows set
	return output.Set(rows)
}

// SetSubMatrix sets the submatrix into the current matrix
//
// @param *F2 m        The submatrix to use
// @param int startRow The first row to replace
// @param int startCol The first column to replace
//
// @return *F2, error
func (f *F2) SetSubMatrix(m *F2, startRow, startCol int) (*F2, error) {
	// verify that the dimensions fit
	if (m.N+startRow) > f.N || (m.M+startCol) > f.M {
		return nil, fmt.Errorf("Submatrix too large")
	}

	// create the bitmask
	bitMask := big.NewInt(0).Exp(
		big.NewInt(2), big.NewInt(int64(m.M)), nil,
	)

	// decrease the bitmask by one
	bitMask.Sub(bitMask, big.NewInt(1))

	// shift the bitmask to the correct position
	bitMask.Lsh(bitMask, uint(startCol))

	// create the output matrix
	output := NewF2(f.N, f.M).Set(f.Rows)

	// iterate through the rows
	for i := startRow; i < (m.N + startRow); i++ {
		// get the bits that are going to be replaced from the current matrix
		subBitMask := big.NewInt(0).And(
			f.Rows[i],
			bitMask,
		)

		// xor the bits with themselves in order to set the to 0
		subBitMask.Xor(
			f.Rows[i],
			subBitMask,
		)

		// xor the shifted bits from the submatrix into the "erased" space
		outputBits := big.NewInt(0).Xor(
			subBitMask,
			big.NewInt(0).Lsh(m.Rows[i-startRow], uint(startCol)),
		)

		// save the row
		output.Rows[i] = outputBits
	}

	// set the rows back into f
	f.Set(output.Rows)

	// return success
	return f, nil
}
