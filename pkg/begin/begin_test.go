package begin

import (
	"fmt"
	"testing"
)

//  Пример использования unit-тестов

func TestBegin15(t *testing.T) {
	type dataForBegin15 struct {
		s     float64
		wantD float64
		wantL float64
	}
	var tests = []dataForBegin15{
		{1, 1.1286652959662007, 7},
		{7, 2.9861676865556794, 14},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v, %v", tt.s, tt.wantD, tt.wantL)
		functionForTest := func(t *testing.T) {
			ansD, ansL := Begin15(tt.s)
			if ansD != tt.wantD {
				t.Errorf("got %v, wantD %v", ansD, tt.wantD)
			}
			if ansL != tt.wantL {
				t.Errorf("got %v, wantL %v", ansL, tt.wantL)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestBegin20(t *testing.T) {
	s := Begin20(0, 3, 0, 2)
	if s != 3.605551275463989 {
		t.Errorf("s (%v) != %v", s, 3.605551275463989)
	}
}

func TestBegin21(t *testing.T) {
	s, per := Begin21(0, 1, 2, 0, 1, 2)
	if s != 0.866025403784439 {
		t.Errorf("s (%v) != %v", s, 0.866025403784439)
	}
	if per != 4.242640687119286 {
		t.Errorf("per (%v) != %v", per, 4.242640687119286)
	}
}

func TestBegin22(t *testing.T) {
	a, b := Begin22(2, 4)
	if a != 4 {
		t.Errorf("a (%v) != %v", a, 4)
	}
	if b != 2 {
		t.Errorf("b (%v) != %v", b, 2)
	}
}

func TestBegin23(t *testing.T) {
	a, b, c := Begin23(2, 4, 6)
	if a != 6 {
		t.Errorf("a (%v) != %v", a, 6)
	}
	if b != 2 {
		t.Errorf("b (%v) != %v", b, 2)
	}
	if c != 4 {
		t.Errorf("c (%v) != %v", c, 4)
	}
}

func TestBegin24(t *testing.T) {
	a, b, c := Begin23(2, 4, 6)
	if a != 6 {
		t.Errorf("a (%v) != %v", a, 6)
	}
	if b != 2 {
		t.Errorf("b (%v) != %v", b, 2)
	}
	if c != 4 {
		t.Errorf("c (%v) != %v", c, 4)
	}
}

func TestBegin25(t *testing.T) {
	y := Begin25(2)
	if y != 175 {
		t.Errorf("y (%v) != %v", y, 175)
	}
}

func TestBegin26(t *testing.T) {
	y := Begin26(3)
	if y != 2 {
		t.Errorf("y (%v) != %v", y, 2)
	}
}

func TestBegin27(t *testing.T) {
	a2, a4, a8 := Begin27(2)
	if a2 != 4 {
		t.Errorf("a2 (%v) != %v", a2, 4)
	}
	if a4 != 16 {
		t.Errorf("a4 (%v) != %v", a4, 16)
	}
	if a8 != 256 {
		t.Errorf("a8 (%v) != %v", a8, 256)
	}
}

func TestBegin28(t *testing.T) {
	a2, a3, a5, a10, a15 := Begin28(2)
	if a2 != 4 {
		t.Errorf("a2 (%v) != %v", a2, 4)
	}
	if a3 != 8 {
		t.Errorf("a3 (%v) != %v", a3, 8)
	}
	if a5 != 32 {
		t.Errorf("a5 (%v) != %v", a5, 32)
	}
	if a10 != 1024 {
		t.Errorf("a10 (%v) != %v", a10, 1024)
	}
	if a15 != 32768 {
		t.Errorf("a15 (%v) != %v", a15, 32768)
	}
}

func TestBegin29(t *testing.T) {
	p := Begin29(360)
	if p != 6.28 {
		t.Errorf("p (%v) != %v", p, 6.28)
	}
}

func TestBegin30(t *testing.T) {
	p := Begin30(6.28)
	if p != 360 {
		t.Errorf("p (%v) != %v", p, 360)
	}
}

func TestBegin31(t *testing.T) {
	c := Begin31(32)
	if c != 0 {
		t.Errorf("c (%v) != %v", c, 0)
	}
}

func TestBegin32(t *testing.T) {
	f := Begin32(89.6)
	if f != 121.6 {
		t.Errorf("f (%v) != %v", f, 121.6)
	}
}

func TestBegin33(t *testing.T) {
	xa, ya := Begin33(2, 342, 65)
	if xa != 171 {
		t.Errorf("xa (%v) != %v", xa, 171)
	}
	if ya != 342 {
		t.Errorf("ya (%v) != %v", ya, 342)
	}
}

func TestBegin34(t *testing.T) {
	xa, yb, r := Begin34(2, 342, 2, 45)
	if xa != 171 {
		t.Errorf("xa (%v) != %v", xa, 171)
	}
	if yb != 22.5 {
		t.Errorf("xq (%v) != %v", yb, 22.5)
	}
	if r != 7.6 {
		t.Errorf("r (%v) != %v", r, 7.6)
	}
}

func TestBegin35(t *testing.T) {
	s1, s2 := Begin35(50, 2, 2, 3)
	if s1 != 100 {
		t.Errorf("s1 (%v) != %v", s1, 100)
	}
	if s2 != 144 {
		t.Errorf("s2 (%v) != %v", s2, 144)
	}
}

func TestBegin36(t *testing.T) {
	so := Begin36(50, 20, 27, 5)
	if so != 377 {
		t.Errorf("so (%v) != %v", so, 377)
	}
}

func TestBegin37(t *testing.T) {
	so := Begin37(50, 20, 70, 5)
	if so != 280 {
		t.Errorf("so (%v) != %v", so, 280)
	}
}

func TestBegin38(t *testing.T) {
	x := Begin38(2, 4)
	if x != -2 {
		t.Errorf("x (%v) != %v", x, -2)
	}
}

func TestBegin39(t *testing.T) {
	x1, x2, d := Begin39(1, 1, 1)
	if x1 != 0 {
		t.Errorf("x1 (%v) != %v", x1, 0)
	}
	if x2 != 0 {
		t.Errorf("x2 (%v) != %v", x2, 0)
	}
	if d != -3 {
		t.Errorf("d (%v) != %v", d, -3)
	}
}

func TestBegin40(t *testing.T) {
	x, y := Begin40(2, 3, 2, 2, 8, 8)
	if x != -0.8 {
		t.Errorf("x (%v) != %v", x, -0.8)
	}
	if y != 1.2 {
		t.Errorf("y (%v) != %v", y, 1.2)
	}
}
