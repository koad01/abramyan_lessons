package begin

import (
	"math"
)

const PI = 3.14

func Begin15(s float64) (d, l float64) {
	r := math.Sqrt(s / PI)
	l = (2 * PI * r)
	d = (2 * r)
	return
}

func Begin20(x1, x2, y1, y2 float64) (s float64) {
	s = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return
}

func Begin21(x1, x2, x3, y1, y2, y3 float64) (s, per float64) {
	per = math.Sqrt(math.Pow(x2-x1, 2)+math.Pow(y2-y1, 2)) + math.Sqrt(math.Pow(x3-x2, 2)+math.Pow(y3-y2, 2)) + math.Sqrt(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2))
	p := per / 2
	s = math.Sqrt(p * ((p - math.Sqrt(math.Pow(x2-x1, 2)+math.Pow(y2-y1, 2))) * (p - math.Sqrt(math.Pow(x3-x2, 2)+math.Pow(y3-y2, 2))) * (p - math.Sqrt(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2)))))
	return
}

func Begin22(a, b float64) (float64, float64) {
	return b, a
}

func Begin23(a, b, c float64) (float64, float64, float64) {
	return c, a, b
}

func Begin24(a, b, c float64) (float64, float64, float64) {
	return b, c, a
}

func Begin25(x float64) (y float64) {
	y = (3 * math.Pow(x, 6)) - (6*(math.Pow(x, 2)) - 7)
	return
}

func Begin26(x float64) (y float64) {
	y = (4*(math.Pow((x-3), 6)) - (7 * (math.Pow((x - 3), 3))) + 2)
	return
}

func Begin27(a float64) (a2, a4, a8 float64) {
	a2 = a * a
	a4 = a2 * a2
	a8 = a4 * a4
	return
}

func Begin28(a float64) (a2, a3, a5, a10, a15 float64) {
	a2 = a * a
	a3 = a2 * a
	a5 = a3 * a2
	a10 = a5 * a5
	a15 = a10 * a5
	return
}

func Begin29(x float64) (p float64) {
	p = (x * PI) / 180
	return
}

func Begin30(p float64) (x float64) {
	x = (p * 180) / PI
	return
}

func Begin31(f float64) (c float64) {
	c = (f - 32) * (5 / 9)
	return
}

func Begin32(c float64) (f float64) {
	f = (9/5)*c + 32
	return
}

func Begin33(x, a, y float64) (xa, ya float64) {
	xa = a / x
	ya = x * xa
	return
}

func Begin34(x, a, y, b float64) (xa, yb, r float64) {
	xa = a / x
	yb = b / y
	r = xa / yb
	return
}

func Begin35(v, u, t1, t2 float64) (s1, s2 float64) {
	s1 = v * t1
	s2 = (v - u) * t2
	return
}

func Begin36(v1, v2, s, t float64) (so float64) {
	so = ((v1 + v2) * t) + s
	return
}

func Begin37(v1, v2, s, t float64) (so float64) {
	so = s - ((v1 + v2) * t)
	if so < 0 {
		so *= -1
	}
	return
}

func Begin38(a, b float64) (x float64) {
	x = b / a * -1
	return
}

func Begin39(a, b, c float64) (x1, x2, d float64) {
	d = (b * b) - 4*a*c
	if d < 0 {
		x1 = 0
		x2 = 0
	} else if d > 0 {
		x1 = (((b * -1) - math.Sqrt(d)) / 2 / a)
		x2 = (((b * -1) + math.Sqrt(d)) / 2 / a)
	} else {
		x1 = ((b * -1) / (2 * a))
		x2 = 0
	}
	if x1 > x2 {
		c := x1
		x1 = x2
		x2 = c
	}
	return x1, x2, d
}

func Begin40(a1, b1, c1, a2, b2, c2 float64) (x, y float64) {
	d := a1*b2 - a2*b1
	x = (c1*b2 - c2*b1) / d
	y = (a1*c2 - a2*c1) / d
	return
}
