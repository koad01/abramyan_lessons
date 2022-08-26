package integer

func Integer1(l int64) (n int64) {
	n = l / 100
	return
}

func Integer2(kg int64) (tn int64) {
	tn = kg / 1000
	return
}

func Integer3(bait int64) (kbait int64) {
	kbait = bait / 1024
	return
}

func Integer4(a, b int64) (z int64) {
	z = a / b
	return
}

func Integer5(a, b int64) (z int64) {
	z = a % b
	return
}

func Integer6(a int64) (d, e int64) {
	d = a / 10
	e = a % 10
	return
}

func Integer7(a int64) (s, p int64) {
	d := a / 10
	e := a % 10
	s = d + e
	p = d * e
	return
}

func Integer8(a int64) (a1 int64) {
	d := a / 10
	e := a % 10
	a1 = e*10 + d
	return
}

func Integer9(a int64) (s int64) {
	s = a / 100
	return
}

func Integer10(a int64) (e, d int64) {
	e = a % 10
	d = (a / 10) % 10
	return
}

func Integer11(a int64) (s, p int64) {
	e := a % 10
	d := (a / 10) % 10
	c := a / 100
	s = e + d + c
	p = e * d * c
	return
}

func Integer12(a int64) (a1 int64) {
	e := a % 10
	d := (a / 10) % 10
	c := a / 100
	a1 = e*100 + d*10 + c
	return
}

func Integer13(a int64) (a1 int64) {
	e := a % 10
	d := (a / 10) % 10
	c := a / 100
	a1 = d*100 + e*10 + c
	return
}

func Integer14(a int64) (a1 int64) {
	e := a % 10
	d := (a / 10) % 10
	c := a / 100
	a1 = e*100 + c*10 + d
	return
}

func Integer15(a int64) (a1 int64) {
	e := a % 10
	d := (a / 10) % 10
	c := a / 100
	a1 = d*100 + c*10 + e
	return
}

func Integer16(a int64) (a1 int64) {
	e := a % 10
	d := (a / 10) % 10
	c := a / 100
	a1 = c*100 + e*10 + d
	return
}

func Integer17(a int64) (c int64) {
	c = a / 100 % 10
	return
}

func Integer18(a int64) (ts int64) {
	ts = a / 1000 % 10
	return
}

func Integer19(s int64) (m int64) {
	m = s / 60
	return
}

func Integer20(s int64) (h int64) {
	h = s / 3600
	return
}

func Integer21(s int64) (s1 int64) {
	s1 = s % 60
	return
}

func Integer22(s int64) (s1 int64) {
	s1 = s % 3600
	return
}

func Integer23(s int64) (m int64) {
	m = s / 60 % 60
	return
}

func Integer24(d int64) (vn int64) {
	vn = d % 7
	return
}

func Integer25(d int64) (vn int64) {
	vn = d%7 + 4
	return
}

func Integer26(d int64) (vn int64) {
	vn = d%7 + 1
	return
}

func Integer27(d int64) (vn int64) {
	vn = d%7 + 5
	return
}

func Integer28(d, vn int64) (nn int64) {
	nn = ((d + vn - 2) % 7) + 1
	return
}

func Integer29(a, b, c int64) (n, so int64) {
	sp := a * b
	sd := c * c
	n = sp / sd
	so = sp % sd
	return
}

func Integer30(y int64) (c int64) {
	c = ((y - 1) / 100) + 1
	return
}
