package gomatrix

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
