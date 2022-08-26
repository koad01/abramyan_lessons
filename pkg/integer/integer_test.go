package integer

import (
	"testing"
)

func TestInteger1(t *testing.T) {
	n := Integer1(2076)
	if n != 20 {
		t.Errorf("n (%v) != %v", n, 20)
	}
}

func TestInteger2(t *testing.T) {
	tn := Integer2(4657)
	if tn != 4 {
		t.Errorf("tn (%v) != %v", tn, 4)
	}
}

func TestInteger3(t *testing.T) {
	kbait := Integer3(4653)
	if kbait != 4 {
		t.Errorf("kbait (%v) != %v", kbait, 4)
	}
}

func TestInteger4(t *testing.T) {
	z := Integer4(20, 6)
	if z != 3 {
		t.Errorf("z (%v) != %v", z, 3)
	}
}

func TestInteger5(t *testing.T) {
	z := Integer5(20, 6)
	if z != 2 {
		t.Errorf("z (%v) != %v", z, 2)
	}
}

func TestInteger6(t *testing.T) {
	d, e := Integer6(23)
	if d != 2 {
		t.Errorf("d (%v) != %v", d, 2)
	}
	if e != 3 {
		t.Errorf("e (%v) != %v", e, 3)
	}
}

func TestInteger7(t *testing.T) {
	s, p := Integer7(23)
	if s != 5 {
		t.Errorf("s (%v) != %v", s, 5)
	}
	if p != 6 {
		t.Errorf("p (%v) != %v", p, 6)
	}
}

func TestInteger8(t *testing.T) {
	a1 := Integer8(26)
	if a1 != 62 {
		t.Errorf("a1 (%v) != %v", a1, 62)
	}
}

func TestInteger9(t *testing.T) {
	s := Integer9(246)
	if s != 2 {
		t.Errorf("s (%v) != %v", s, 2)
	}
}

func TestInteger10(t *testing.T) {
	e, d := Integer10(239)
	if e != 9 {
		t.Errorf("e (%v) != %v", e, 9)
	}
	if d != 3 {
		t.Errorf("d (%v) != %v", d, 3)
	}
}

func TestInteger11(t *testing.T) {
	s, p := Integer11(121)
	if s != 4 {
		t.Errorf("s (%v) != %v", s, 4)
	}
	if p != 2 {
		t.Errorf("p (%v) != %v", p, 2)
	}
}

func TestInteger12(t *testing.T) {
	a1 := Integer12(246)
	if a1 != 642 {
		t.Errorf("a1 (%v) != %v", a1, 642)
	}
}

func TestInteger13(t *testing.T) {
	a1 := Integer13(246)
	if a1 != 462 {
		t.Errorf("a1 (%v) != %v", a1, 462)
	}
}

func TestInteger14(t *testing.T) {
	a1 := Integer14(246)
	if a1 != 624 {
		t.Errorf("a1 (%v) != %v", a1, 624)
	}
}

func TestInteger15(t *testing.T) {
	a1 := Integer15(246)
	if a1 != 426 {
		t.Errorf("a1 (%v) != %v", a1, 426)
	}
}

func TestInteger16(t *testing.T) {
	a1 := Integer16(246)
	if a1 != 264 {
		t.Errorf("a1 (%v) != %v", a1, 264)
	}
}

func TestInteger17(t *testing.T) {
	c := Integer17(1246)
	if c != 2 {
		t.Errorf("c (%v) != %v", c, 2)
	}
}

func TestInteger18(t *testing.T) {
	ts := Integer18(13246)
	if ts != 3 {
		t.Errorf("ts (%v) != %v", ts, 3)
	}
}

func TestInteger19(t *testing.T) {
	m := Integer19(139)
	if m != 2 {
		t.Errorf("m (%v) != %v", m, 2)
	}
}

func TestInteger20(t *testing.T) {
	h := Integer20(139667)
	if h != 38 {
		t.Errorf("h (%v) != %v", h, 38)
	}
}

func TestInteger21(t *testing.T) {
	s1 := Integer21(139667)
	if s1 != 47 {
		t.Errorf("s1 (%v) != %v", s1, 47)
	}
}

func TestInteger22(t *testing.T) {
	s1 := Integer22(139667)
	if s1 != 2867 {
		t.Errorf("s1 (%v) != %v", s1, 2867)
	}
}

func TestInteger23(t *testing.T) {
	m := Integer23(139667)
	if m != 47 {
		t.Errorf("m (%v) != %v", m, 47)
	}
}

func TestInteger24(t *testing.T) {
	vd := Integer24(14)
	if vd != 0 {
		t.Errorf("vd (%v) != %v", vd, 0)
	}
}

func TestInteger25(t *testing.T) {
	vd := Integer25(14)
	if vd != 4 {
		t.Errorf("vd (%v) != %v", vd, 4)
	}
}

func TestInteger26(t *testing.T) {
	vd := Integer26(14)
	if vd != 1 {
		t.Errorf("vd (%v) != %v", vd, 1)
	}
}

func TestInteger27(t *testing.T) {
	vd := Integer27(14)
	if vd != 5 {
		t.Errorf("vd (%v) != %v", vd, 5)
	}
}

func TestInteger28(t *testing.T) {
	nn := Integer28(145, 6)
	if nn != 3 {
		t.Errorf("nn (%v) != %v", nn, 3)
	}
}

func TestInteger29(t *testing.T) {
	n, so := Integer29(5, 6, 2)
	if n != 7 {
		t.Errorf("n (%v) != %v", n, 7)
	}
	if so != 2 {
		t.Errorf("so (%v) != %v", so, 2)
	}
}

func TestInteger30(t *testing.T) {
	c := Integer30(14590)
	if c != 146 {
		t.Errorf("c (%v) != %v", c, 146)
	}
}
