package resolver

import (
	"fmt"
	"math/big"
	"testing"

	"git.noc.ruhr-uni-bochum.de/danieljankowski/gomatrix"

	"github.com/stretchr/testify/assert"
)

func TestLinearDependenciesInGauss(t *testing.T) {
	tests := []struct {
		description    string
		matrix         *gomatrix.F2
		startRow       int
		startCol       int
		stopRow        int
		stopCol        int
		pivotBit       int
		expectedError  bool
		expectedResult *gomatrix.F2
	}{
		{
			description: "simple swap and postprocessing of the row",
			matrix: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(13),
				big.NewInt(0),
				big.NewInt(14),
			}),
			startRow:      0,
			startCol:      1,
			stopRow:       2,
			stopCol:       3,
			pivotBit:      3,
			expectedError: false,
			expectedResult: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(13),
				big.NewInt(9),
				big.NewInt(0),
			}),
		},
		{
			description: "simple swap of columns",
			matrix: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(13),
				big.NewInt(1),
				big.NewInt(1),
			}),
			startRow:      0,
			startCol:      1,
			stopRow:       2,
			stopCol:       3,
			pivotBit:      3,
			expectedError: false,
			expectedResult: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(13),
				big.NewInt(8),
				big.NewInt(8),
			}),
		},
		{
			description: "simple swap of columns + continue",
			matrix: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(5),
				big.NewInt(14),
				big.NewInt(8),
				big.NewInt(8),
			}),
			startRow:      0,
			startCol:      0,
			stopRow:       2,
			stopCol:       2,
			pivotBit:      2,
			expectedError: false,
			expectedResult: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(9),
				big.NewInt(14),
				big.NewInt(4),
				big.NewInt(4),
			}),
		},
		{
			description: "no way to resolve the dependency",
			matrix: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(13),
				big.NewInt(0),
				big.NewInt(0),
			}),
			startRow:      0,
			startCol:      1,
			stopRow:       2,
			stopCol:       3,
			pivotBit:      3,
			expectedError: true,
			expectedResult: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(3),
				big.NewInt(13),
				big.NewInt(8),
				big.NewInt(8),
			}),
		},
		{
			description: "simple swap and postprocessing of the row + continue",
			matrix: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(13),
				big.NewInt(0),
				big.NewInt(10),
			}),
			startRow:      0,
			startCol:      1,
			stopRow:       2,
			stopCol:       3,
			pivotBit:      3,
			expectedError: false,
			expectedResult: gomatrix.NewF2(4, 4).Set([]*big.Int{
				big.NewInt(10),
				big.NewInt(13),
				big.NewInt(0),
				big.NewInt(0),
			}),
		},
	}

	for _, test := range tests {
		fmt.Printf("%s\n", test.description)
		err := LinearDependenciesInGauss(
			test.matrix,
			test.startRow,
			test.startCol,
			test.stopRow,
			test.stopCol,
			test.pivotBit,
		)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if err != nil {
			continue
		}

		test.matrix.PrintSlim()

		assert.Truef(t, test.expectedResult.IsEqual(test.matrix), test.description)
	}
}
