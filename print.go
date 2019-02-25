// Package gomatrix Is a go package for scientific operations with matrices in F2.
package gomatrix

import (
	"fmt"
)

// PrettyPrint prints the matrix to stdout
func (f *F2) PrettyPrint() {
	f.printWithSeperators(" ", "\n")
}

// PrintLaTex prints the matrix as latex code
func (f *F2) PrintLaTex() {
	fmt.Printf("\\begin{bmatrix}\n")
	f.printWithSeperators(" & ", "\\\n")
	fmt.Printf("\\end{bmatrix}\n")
}

// PrintCSV prints the matrix as csv
func (f *F2) PrintCSV() {
	f.printWithSeperators(", ", "\n")
}

// printWithSeperators prints the matrix with custom seperators
//
// @param string valSep  The seperator for the single values
// @param string lineSep The line seperator
func (f *F2) printWithSeperators(valSep, lineSep string) {
	for _, row := range f.Rows {
		for i := 0; i < f.M; i++ {
			if i == f.M-1 {
				fmt.Printf("%d ", row.Bit(i))
				continue
			}
			fmt.Printf("%d%s", row.Bit(i), valSep)
		}
		fmt.Print(lineSep)
	}
}
