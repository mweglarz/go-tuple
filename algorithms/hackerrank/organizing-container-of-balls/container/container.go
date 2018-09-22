package container

import (
	"math"
)

type matrix struct {
	rows int
	cols int
}

func (A *matrix) Rows() int { return A.rows }

func (A *matrix) Cols() int { return A.cols }

type DenseMatrix struct {
	matrix
	// flattened matrix data. elements[i*step+j] is row i, col j
	elements []float64
	// actual offset between rows
	step int
}

func (A *DenseMatrix) Det() float64 {
	B := A.Copy()
	P := B.LUInPlace()
	return product(B.DiagonalCopy()) * P.Det()
}

func (A *DenseMatrix) Copy() *DenseMatrix {
	B := new(DenseMatrix)
	B.rows = A.rows
	B.cols = A.cols
	B.step = A.cols
	B.elements = make([]float64, B.rows*B.cols)
	for row := 0; row < B.rows; row++ {
		copy(B.rowSlice(row), A.rowSlice(row))
	}
	return B
}

func (A *DenseMatrix) DiagonalCopy() []float64 {
	span := A.rows
	if A.cols < span {
		span = A.cols
	}
	diag := make([]float64, span)
	for i := 0; i < span; i++ {
		diag[i] = A.Get(i, i)
	}
	return diag
}

/*
Overwrites A with [L\U] and returns P, st PLU=A. L is considered to
have 1s in the diagonal.
*/
func (A *DenseMatrix) LUInPlace() (P *PivotMatrix) {
	m := A.Rows()
	n := A.Cols()
	LUcolj := make([]float64, m)
	LUrowi := make([]float64, n)
	piv := make([]int, m)
	for i := 0; i < m; i++ {
		piv[i] = i
	}
	pivsign := float64(1.0)

	for j := 0; j < n; j++ {
		A.BufferCol(j, LUcolj)
		for i := 0; i < m; i++ {
			A.BufferRow(i, LUrowi)
			kmax := i
			if j < i {
				kmax = j
			}
			s := float64(0)
			for k := 0; k < kmax; k++ {
				s += LUrowi[k] * LUcolj[k]
			}
			LUcolj[i] -= s
			LUrowi[j] = LUcolj[i]
			A.Set(i, j, LUrowi[j])
		}

		p := j
		for i := j + 1; i < m; i++ {
			if math.Abs(LUcolj[i]) > math.Abs(LUcolj[p]) {
				p = i
			}
		}
		if p != j {
			A.SwapRows(p, j)
			k := piv[p]
			piv[p] = piv[j]
			piv[j] = k
			pivsign = -pivsign
		}

		if j < m && A.Get(j, j) != 0 {
			for i := j + 1; i < m; i++ {
				A.Set(i, j, A.Get(i, j)/A.Get(j, j))
			}
		}
	}

	P = MakePivotMatrix(piv, pivsign)

	return
}

func (A *DenseMatrix) rowSlice(row int) []float64 {
	return A.elements[row*A.step : row*A.step+A.cols]
}

func (A *DenseMatrix) BufferCol(j int, buf []float64) {
	for i := 0; i < A.rows; i++ {
		buf[i] = A.Get(i, j)
	}
}

func (A *DenseMatrix) BufferRow(i int, buf []float64) {
	for j := 0; j < A.cols; j++ {
		buf[j] = A.Get(i, j)
	}
}

func (m *DenseMatrix) SwapRows(r1 int, r2 int) {
	index1 := r1 * m.step
	index2 := r2 * m.step
	for j := 0; j < m.cols; j++ {
		m.elements[index1], m.elements[index2] = m.elements[index2], m.elements[index1]
		index1++
		index2++
	}
}

func (A *DenseMatrix) Set(i int, j int, v float64) {
	/*
		i = i % A.rows
		if i < 0 {
			i = A.rows - i
		}
		j = j % A.cols
		if j < 0 {
			j = A.cols - j
		}
	*/
	// reslicing like this does efficient range checks, perhaps
	A.elements[i*A.step : i*A.step+A.cols][j] = v
	//A.elements[i*A.step+j] = v
}

func (A *DenseMatrix) Get(i int, j int) (v float64) {
	/*
		i = i % A.rows
		if i < 0 {
			i = A.rows - i
		}
		j = j % A.cols
		if j < 0 {
			j = A.cols - j
		}
	*/

	// reslicing like this does efficient range checks, perhaps
	v = A.elements[i*A.step : i*A.step+A.cols][j]
	//v = A.elements[i*A.step+j]
	return
}

type PivotMatrix struct {
	matrix
	pivots    []int
	pivotSign float64
}

func (P *PivotMatrix) Det() float64 { return P.pivotSign }

func MakePivotMatrix(pivots []int, pivotSign float64) *PivotMatrix {
	n := len(pivots)
	P := new(PivotMatrix)
	P.rows = n
	P.cols = n
	P.pivots = pivots
	P.pivotSign = pivotSign
	return P
}

func product(a []float64) float64 {
	p := float64(1)
	for _, v := range a {
		p *= v
	}
	return p
}

func Zeros(rows, cols int) *DenseMatrix {
	A := new(DenseMatrix)
	A.elements = make([]float64, rows*cols)
	A.rows = rows
	A.cols = cols
	A.step = cols
	return A
}

