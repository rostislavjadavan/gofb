package gofb

import (
	"testing"
)

func TestVec2Distance(t *testing.T) {
	v1 := NewVec2(0, 0)
	d1 := v1.Distance(v1)
	if d1 != 0 {
		t.Errorf("invalid distance d3 %f", d1)
	}

	v2 := NewVec2(1, 0)
	d2 := v1.Distance(v2)
	if d2 != 1 {
		t.Errorf("invalid distance d3 %f", d2)
	}

	v3 := NewVec2(0, 3)
	d3 := v1.Distance(v3)
	if d3 != 3 {
		t.Errorf("invalid distance d3 %f", d3)
	}
}

func TestVec2IsInside(t *testing.T) {
	a := NewVec2(30, 30)

	v1 := NewVec2(5, 5)
	if !v1.IsInside(a) {
		t.Errorf("v1 is not inside a")
	}

	v2 := NewVec2(0, 0)
	if !v2.IsInside(a) {
		t.Errorf("v2 is not inside a")
	}

	v3 := NewVec2(30, 29)
	if v3.IsInside(a) {
		t.Errorf("v3 is inside a")
	}

	v4 := NewVec2(30, 30)
	if v4.IsInside(a) {
		t.Errorf("v4 is inside a")
	}
}

func TestVec2IsInside2(t *testing.T) {
	a := NewVec2(0, 0)
	b := NewVec2(30, 30)

	v1 := NewVec2(5, 5)
	if !v1.IsInside2(a, b) {
		t.Errorf("v1 is not inside a,b")
	}

	v2 := NewVec2(0, 0)
	if !v2.IsInside2(a, b) {
		t.Errorf("v2 is not inside a,b")
	}

	v3 := NewVec2(30, 29)
	if v3.IsInside2(a, b) {
		t.Errorf("v3 is inside a,b")
	}

	v4 := NewVec2(30, 30)
	if v4.IsInside2(a, b) {
		t.Errorf("v4 is inside a,b")
	}
}

func TestVec2ClipInside(t *testing.T) {
	a := NewVec2(3, 3)
	b := NewVec2(30, 30)

	v1 := NewVec2(-5, -5)
	v1.ClipInside2(a, b)
	if v1.X != 3 && v1.Y != 3 {
		t.Errorf("v1 was no clipped correctly")
	}

	v2 := NewVec2(50, 100)
	v2.ClipInside2(a, b)
	if v2.X != 30 && v2.Y != 30 {
		t.Errorf("v2 was no clipped correctly")
	}
}

func TestRegionIsInside(t *testing.T) {
	r := NewRegion(10, 10, 20, 20)

	v1 := NewVec2(15, 15)
	if !r.IsInside(v1) {
		t.Errorf("v1 is not inside r")
	}

	v2 := NewVec2(0, 0)
	if r.IsInside(v2) {
		t.Errorf("v2 is inside r")
	}

	v3 := NewVec2(30, 30)
	if r.IsInside(v3) {
		t.Errorf("v3 is inside r")
	}
}
