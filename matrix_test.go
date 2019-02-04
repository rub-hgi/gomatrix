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
