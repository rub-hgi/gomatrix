package gomatrix

import (
	_ "fmt"
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
			assert.Equal(t, 0, row.Cmp(test.data[i]))
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
