package container

import (
	"math"
	"testing"
)

func TestDet1(t *testing.T) {
	m := Zeros(2, 2)
	m.Set(0, 0, 1)
	m.Set(0, 1, 1)
	m.Set(1, 0, 1)
	m.Set(1, 1, 1)

	mDet := m.Det()
	if math.Abs(mDet) > 10e-9 {
		t.Errorf("Expecting determinant to be zero, got %.2f", mDet)
	}
}

func TestDet2(t *testing.T) {
	m := Zeros(2, 2)
	m.Set(0, 0, 0)
	m.Set(0, 1, 2)
	m.Set(1, 0, 1)
	m.Set(1, 1, 1)

	mDet := m.Det()
	if math.Abs(mDet)-2.0 > 10e-9 {
		t.Errorf("Expecting determinant to be zero, got %.2f", mDet)
	}
}
