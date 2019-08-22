package gomatrix

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewF2(t *testing.T) {
	tests := []struct {
		description  string
		n            int
		m            int
		expectedType interface{}
	}{
		{
			description:  "square matrix",
			n:            6,
			m:            6,
			expectedType: &F2{},
		},
	}

	for _, test := range tests {
		mat := NewF2(test.n, test.m)

		assert.IsType(t, test.expectedType, mat)
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		description string
		n           int
		m           int
		data        []*big.Int
		expectedNil bool
	}{
		{
			description: "4x4 matrix",
			n:           2,
			m:           2,
			data:        []*big.Int{big.NewInt(1), big.NewInt(2)},
			expectedNil: false,
		},
		{
			description: "wrong row count",
			n:           3,
			m:           2,
			data:        []*big.Int{big.NewInt(1), big.NewInt(2)},
			expectedNil: true,
		},
		{
			description: "wrong col count",
			n:           2,
			m:           2,
			data:        []*big.Int{big.NewInt(5), big.NewInt(2)},
			expectedNil: true,
		},
	}

	for _, test := range tests {
		m := NewF2(test.n, test.m)
		err := m.Set(test.data)

		assert.Equal(t, test.expectedNil, err == nil)

		if err == nil {
			continue
		}

		for i, row := range m.Rows {
			assert.Zero(t, row.Cmp(test.data[i]))
		}
	}
}

func TestAt(t *testing.T) {
	tests := []struct {
		description    string
		matrix         *F2
		i              int
		j              int
		expectedError  bool
		expectedResult int
	}{
		{
			description:    "success",
			matrix:         NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			i:              0,
			j:              1,
			expectedError:  false,
			expectedResult: 1,
		},
		{
			description:    "success",
			matrix:         NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			i:              1,
			j:              1,
			expectedError:  false,
			expectedResult: 0,
		},
		{
			description:    "invalid j",
			matrix:         NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			i:              1,
			j:              3,
			expectedError:  true,
			expectedResult: 0,
		},
	}

	for _, test := range tests {
		result, err := test.matrix.At(test.i, test.j)

		assert.Equal(t, test.expectedError, err != nil)

		if err != nil {
			continue
		}

		assert.Equal(t, test.expectedResult, result)
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		matrixB        *F2
		expectedResult bool
	}{
		{
			description:    "equal matrices",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			matrixB:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			expectedResult: true,
		},
		{
			description:    "inequal matrices",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			matrixB:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(2)}),
			expectedResult: false,
		},
		{
			description:    "different dimensions",
			matrixA:        NewF2(2, 3).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			matrixB:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			expectedResult: false,
		},
	}

	for _, test := range tests {
		result := test.matrixA.IsEqual(test.matrixB)

		assert.Equal(t, test.expectedResult, result)
	}
}

func TestT(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		expectedMatrix *F2
	}{
		{
			description:    "small matrix",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(0)}),
			expectedMatrix: NewF2(2, 2).Set([]*big.Int{big.NewInt(0), big.NewInt(1)}),
		},
		{
			description:    "3x3 matrix",
			matrixA:        NewF2(3, 3).Set([]*big.Int{big.NewInt(2), big.NewInt(4), big.NewInt(1)}),
			expectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(4), big.NewInt(1), big.NewInt(2)}),
		},
		{
			description:    "3x2 matrix",
			matrixA:        NewF2(3, 2).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(0)}),
			expectedMatrix: NewF2(2, 3).Set([]*big.Int{big.NewInt(3), big.NewInt(1)}),
		},
	}

	for _, test := range tests {
		test.matrixA.T()

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}

func TestPartialT(t *testing.T) {
	tests := []struct {
		description    string
		matrix         *F2
		startRow       int
		startCol       int
		n              int
		expectedMatrix *F2
		expectedError  bool
	}{
		{
			description:    "3x3",
			matrix:         NewF2(3, 3).Set([]*big.Int{big.NewInt(5), big.NewInt(6), big.NewInt(4)}),
			startRow:       1,
			startCol:       1,
			n:              2,
			expectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(5), big.NewInt(2), big.NewInt(6)}),
			expectedError:  false,
		},
		{
			description:    "n too big for selected startRow and startCol",
			matrix:         NewF2(3, 3).Set([]*big.Int{big.NewInt(5), big.NewInt(6), big.NewInt(4)}),
			startRow:       1,
			startCol:       1,
			n:              3,
			expectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(5), big.NewInt(2), big.NewInt(6)}),
			expectedError:  true,
		},
	}

	for _, test := range tests {
		err := test.matrix.PartialT(
			test.startRow,
			test.startCol,
			test.n,
		)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if err != nil {
			continue
		}

		assert.Truef(t, test.expectedMatrix.IsEqual(test.matrix), test.description)
	}
}

func TestSetToIdentity(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		expectedMatrix *F2
	}{
		{
			description:    "2x2 matrix",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(3), big.NewInt(1)}),
			expectedMatrix: NewF2(2, 2).Set([]*big.Int{big.NewInt(1), big.NewInt(2)}),
		},
		{
			description:    "3x2 matrix",
			matrixA:        NewF2(3, 2).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(2)}),
			expectedMatrix: NewF2(3, 2).Set([]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(0)}),
		},
	}

	for _, test := range tests {
		test.matrixA.SetToIdentity()

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}

func TestSwapRows(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		i              int
		j              int
		expectedMatrix *F2
		expectedError  bool
	}{
		{
			description:    "2x2 matrix",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(3), big.NewInt(1)}),
			i:              1,
			j:              0,
			expectedMatrix: NewF2(2, 2).Set([]*big.Int{big.NewInt(1), big.NewInt(3)}),
			expectedError:  false,
		},
		{
			description:    "invalid index",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(3), big.NewInt(1)}),
			i:              2,
			j:              0,
			expectedMatrix: NewF2(2, 2),
			expectedError:  true,
		},
	}

	for _, test := range tests {
		err := test.matrixA.SwapRows(test.i, test.j)

		assert.Equal(t, test.expectedError, err != nil)

		if err != nil {
			continue
		}

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}

func TestSwapCols(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		i              int
		j              int
		expectedMatrix *F2
		expectedError  bool
	}{
		{
			description:    "3x3 matrix",
			matrixA:        NewF2(3, 3).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(5)}),
			i:              2,
			j:              0,
			expectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(6), big.NewInt(4), big.NewInt(5)}),
			expectedError:  false,
		},
		{
			description:    "invalid index",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(3), big.NewInt(1)}),
			i:              2,
			j:              0,
			expectedMatrix: NewF2(2, 2),
			expectedError:  true,
		},
	}

	for _, test := range tests {
		err := test.matrixA.SwapCols(test.i, test.j)

		assert.Equal(t, test.expectedError, err != nil)

		if err != nil {
			continue
		}

		assert.True(t, test.matrixA.IsEqual(test.expectedMatrix))
	}
}

func TestPermuteCols(t *testing.T) {
	tests := []struct {
		description       string
		maxRepeats        int
		matrixA           *F2
		notExpectedMatrix *F2
	}{
		{
			description:       "3x3 matrix",
			maxRepeats:        9,
			matrixA:           NewF2(3, 3).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(5)}),
			notExpectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(5)}),
		},
		{
			description:       "3x3 matrix",
			maxRepeats:        9,
			matrixA:           NewF2(3, 3).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(5)}),
			notExpectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(5)}),
		},
		{
			description:       "3x3 matrix",
			maxRepeats:        9,
			matrixA:           NewF2(3, 3).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(5)}),
			notExpectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(5)}),
		},
	}

	for _, test := range tests {
		repeats := 0

		var permMat *F2

		savedMat := NewF2(
			test.matrixA.N,
			test.matrixA.M,
		).Set(test.matrixA.Rows)

		for test.matrixA.IsEqual(test.notExpectedMatrix) && repeats < test.maxRepeats {
			repeats++

			permMat = test.matrixA.PermuteCols()
		}

		assert.False(t, test.matrixA.IsEqual(test.notExpectedMatrix))
		assert.NotNil(t, permMat)

		test.matrixA.MulMatrix(permMat.T())

		assert.True(t, test.matrixA.IsEqual(savedMat))
	}
}

func TestGetCol(t *testing.T) {
	tests := []struct {
		description    string
		matrixA        *F2
		colIndex       int
		expectedResult *big.Int
	}{
		{
			description:    "first column of a 2x2 matrix",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			colIndex:       0,
			expectedResult: big.NewInt(2),
		},
		{
			description:    "second column of a 2x2 matrix",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			colIndex:       1,
			expectedResult: big.NewInt(1),
		},
		{
			description:    "invalid parameter",
			matrixA:        NewF2(2, 2).Set([]*big.Int{big.NewInt(2), big.NewInt(1)}),
			colIndex:       2,
			expectedResult: nil,
		},
	}

	for _, test := range tests {
		result := test.matrixA.GetCol(test.colIndex)

		assert.Equal(t, test.expectedResult, result)
	}
}

func TestGetSubMatrix(t *testing.T) {
	tests := []struct {
		description    string
		matrix         *F2
		startRow       int
		startCol       int
		stopRow        int
		stopCol        int
		expectedMatrix *F2
	}{
		{
			matrix: NewF2(3, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(9),
				big.NewInt(6),
			}),
			startRow: 0,
			startCol: 1,
			stopRow:  2,
			stopCol:  4,
			expectedMatrix: NewF2(2, 3).Set([]*big.Int{
				big.NewInt(5),
				big.NewInt(4),
			}),
		},
		{
			matrix: NewF2(3, 3).Set([]*big.Int{
				big.NewInt(5),
				big.NewInt(6),
				big.NewInt(4),
			}),
			startRow: 1,
			startCol: 1,
			stopRow:  3,
			stopCol:  3,
			expectedMatrix: NewF2(2, 2).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(2),
			}),
		},
	}

	for _, test := range tests {
		result := test.matrix.GetSubMatrix(
			test.startRow,
			test.startCol,
			test.stopRow,
			test.stopCol,
		)

		assert.True(t, test.expectedMatrix.IsEqual(result))
	}
}

func TestSetSubMatrix(t *testing.T) {
	tests := []struct {
		description    string
		matrix         *F2
		submatrix      *F2
		startRow       int
		startCol       int
		expectedError  bool
		expectedMatrix *F2
	}{
		{
			description:    "1x6",
			matrix:         NewF2(1, 6).Set([]*big.Int{big.NewInt(22)}),
			submatrix:      NewF2(1, 3).Set([]*big.Int{big.NewInt(1)}),
			startRow:       0,
			startCol:       3,
			expectedError:  false,
			expectedMatrix: NewF2(1, 6).Set([]*big.Int{big.NewInt(14)}),
		},
		{
			description:    "3x3",
			matrix:         NewF2(3, 3).Set([]*big.Int{big.NewInt(5), big.NewInt(6), big.NewInt(4)}),
			submatrix:      NewF2(2, 2).Set([]*big.Int{big.NewInt(3), big.NewInt(0)}),
			startRow:       0,
			startCol:       1,
			expectedError:  false,
			expectedMatrix: NewF2(3, 3).Set([]*big.Int{big.NewInt(7), big.NewInt(0), big.NewInt(4)}),
		},
		{
			description:    "submatrix too big",
			matrix:         NewF2(3, 3).Set([]*big.Int{big.NewInt(5), big.NewInt(6), big.NewInt(4)}),
			submatrix:      NewF2(2, 4).Set([]*big.Int{big.NewInt(3), big.NewInt(0)}),
			startRow:       1,
			startCol:       1,
			expectedError:  true,
			expectedMatrix: nil,
		},
	}

	for _, test := range tests {
		result, err := test.matrix.SetSubMatrix(
			test.submatrix,
			test.startRow,
			test.startCol,
		)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if err != nil {
			continue
		}

		assert.Truef(t, test.expectedMatrix.IsEqual(result), test.description)
	}
}
