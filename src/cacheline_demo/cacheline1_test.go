package cacheline_demo

import (
	"testing"
)

//https://learnku.com/go/t/45683

//go test -v -bench="MatrixCombination$" -benchmem -count=12 -benchtime 1s -timeout 2s -run=none
//go test -v -bench=. -benchmem -run=none
//go test -v -bench="StructureFalseSharing$" -benchmem -run=none

//var matrixLength int
//
//func init() {
//	flag.IntVar(&matrixLength, "matrixLength", 1000, "print test log")
//	flag.Parse()
//}

var matrixLength = 1000

func createMatrix(matrixLength int) [][]int64 {
	array := make([][]int64, matrixLength)
	for i := 0; i < matrixLength; i++ {
		array[i] = make([]int64, matrixLength)
	}
	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			array[i][j] = 0
		}
	}
	return array
}

func BenchmarkMatrixCombination(b *testing.B) {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)

	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i++ {
			for j := 0; j < matrixLength; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[i][j]
			}
		}
	}
}

func BenchmarkMatrixReversedCombination(b *testing.B) {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)

	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i++ {
			for j := 0; j < matrixLength; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[j][i]
			}
		}
	}
}

func BenchmarkMatrixReversedCombinationPerBlock(b *testing.B) {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)
	blockSize := 8

	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i += blockSize {
			for j := 0; j < matrixLength; j += blockSize {
				for ii := i; ii < i+blockSize; ii++ {
					for jj := j; jj < j+blockSize; jj++ {
						matrixA[ii][jj] = matrixA[ii][jj] + matrixB[jj][ii]
					}
				}
			}
		}
	}
}
