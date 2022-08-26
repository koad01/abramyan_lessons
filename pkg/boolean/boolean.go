package boolean

import "math"

func Boolean1(a int) (bol bool) {
	if a > 0 {
		bol = true
	}
	return
}

func Boolean2(a int) (bol bool) {
	c := a % 2
	if c == 1 {
		bol = true
	}
	return
}

func Boolean3(a int) (bol bool) {
	c := a % 2
	if c == 0 {
		bol = true
	}
	return
}

func Boolean4(a, b int) (bol bool) {
	if a > 2 && b <= 3 {
		bol = true
	}
	return
}

func Boolean5(a, b int) (bol bool) {
	if a >= 0 && b < -2 {
		bol = true
	}
	return
}

func Boolean6(a, b, c int) (bol bool) {
	if a < b && b < c {
		bol = true
	}
	return
}

func Boolean7(a, b, c int) (bol bool) {
	if b > a && b < c {
		bol = true
	}
	return
}

func Boolean8(a, b int) (bol bool) {
	c := a % 2
	d := b % 2
	if c == 1 && d == 1 {
		bol = true
	}
	return
}

func Boolean9(a, b int) (bol bool) {
	c := a % 2
	d := b % 2
	if c == 1 || d == 1 {
		bol = true
	}
	return
}

func Boolean10(a, b int) (bol bool) {
	c := a % 2
	d := b % 2
	if (c == 1 && d == 0) || (c == 0 && d == 1) {
		bol = true
	}
	return
}

func Boolean11(a, b int) (bol bool) {
	c := a % 2
	d := b % 2
	if (c == 1 && d == 1) || (c == 0 && d == 0) {
		bol = true
	}
	return
}

func Boolean12(a, b, c int) (bol bool) {
	if a > 0 && b > 0 && c > 0 {
		bol = true
	}
	return
}

func Boolean13(a, b, c int) (bol bool) {
	if a > 0 || b > 0 || c > 0 {
		bol = true
	}
	return
}

func Boolean14(a, b, c int) (bol bool) {
	if (a > 1 && b > 0 && c > 0) || (a > 0 && b > 1 && c > 0) || (a > 0 && b > 0 && c > 1) {
		bol = true
	}
	return
}

func Boolean15(a, b, c int) (bol bool) {
	if (a > 1 && b > 1 && c > 0) || (a > 0 && b > 1 && c > 1) || (a > 1 && b > 0 && c > 1) {
		bol = true
	}
	return
}

func Boolean16(a int) (bol bool) {
	c := a % 2
	if a < 100 && a > 9 && c == 0 {
		bol = true
	}
	return
}

func Boolean17(a int) (bol bool) {
	c := a % 2
	if a < 1000 && a > 100 && c == 1 {
		bol = true
	}
	return
}

func Boolean18(a, b, c int) (bol bool) {
	if a == b || b == c || c == a {
		bol = true
	}
	return
}

func Boolean19(a, b, c int) (bol bool) {
	if a == -b || b == -c || c == -a {
		bol = true
	}
	return
}

func Boolean20(a, b, c int) (bol bool) {
	if a != b && b != c && c != a {
		bol = true
	}
	return
}

func Boolean21(a int) (bol bool) {
	e := a % 100
	d := a%10 - e
	c := a / 100
	if c < d && d < e {
		bol = true
	}
	return
}

func Boolean22(a int) (bol bool) {
	e := a % 100
	d := a%10 - e
	c := a / 100
	if (c < d && d < e) || (c > d && d > e) {
		bol = true
	}
	return
}

func Boolean23(a int) (bol bool) {
	e := a % 1000
	d := a%100 - e
	c := a%10 - d
	t := a / 1000
	if e == t && d == c {
		bol = true
	}
	return
}

func Boolean24(a, b, c int64) (bol bool) {
	d := (b * b) - 4*a*c
	if d >= 0 {
		bol = true
	}
	return
}

func Boolean25(x, y int64) (bol bool) {
	if x < 0 && y > 0 {
		bol = true
	}
	return
}

func Boolean26(x, y int64) (bol bool) {
	if x > 0 && y < 0 {
		bol = true
	}
	return
}

func Boolean27(x, y int64) (bol bool) {
	if x < 0 && y > 0 || x < 0 && y < 0 {
		bol = true
	}
	return
}

func Boolean28(x, y int64) (bol bool) {
	if x > 0 && y > 0 || x < 0 && y < 0 {
		bol = true
	}
	return
}

func Boolean29(x, y, x1, y1, x2, y2 int64) (bol bool) {
	if x > x1 && x < x2 && y > y1 && y < y2 {
		bol = true
	}
	return
}

func Boolean30(a, b, c int64) (bol bool) {
	if a == b && b == c {
		bol = true
	}
	return
}

func Boolean31(a, b, c int64) (bol bool) {
	if (a == b && b != c) || (a == c && c != b) || (b == c && c != a) {
		bol = true
	}
	return
}

func Boolean32(a, b, c float64) (bol bool) {
	if c == math.Sqrt(a*a+b*b) || b == math.Sqrt(a*a+c*c) || a == math.Sqrt(c*c+c*c) {
		bol = true
	}
	return
}

func Boolean33(a, b, c int64) (bol bool) {
	if (a+b < c) || (a+c < b) || (b+c < a) {
		bol = true
	}
	return
}

func Boolean34(x, y int64) (bol bool) {
	c := (x + y) % 2
	if c == 1 {
		bol = true
	}
	return
}

func Boolean35(x, y, x1, y1 int64) (bol bool) {
	c := (x + y) % 2
	c1 := (x1 + y1) % 2
	if c == c1 {
		bol = true
	}
	return
}

func Boolean36(x1, y1, x2, y2 int64) (bol bool) {
	if x1 == x2 || y2 == y1 {
		bol = true
	}
	return
}

func Boolean37(x1, y1, x2, y2 int64) (bol bool) {
	c1 := x1 - x2
	c2 := x1 - x2
	if c1 < 0 {
		c1 = -c1
	}
	if c2 < 0 {
		c2 = -c2
	}
	if c1 == 1 || c2 == 1 {
		bol = true
	}
	return
}

func Boolean38(x1, y1, x2, y2 int64) (bol bool) {
	c1 := x1 - x2
	c2 := x1 - x2
	if c1 < 0 {
		c1 = -c1
	}
	if c2 < 0 {
		c2 = -c2
	}
	if c1 > 0 && c2 > 0 {
		bol = true
	}
	return
}

func Boolean39(x1, y1, x2, y2 int64) (bol bool) {
	c1 := x1 - x2
	c2 := x1 - x2
	if c1 < 0 {
		c1 = -c1
	}
	if c2 < 0 {
		c2 = -c2
	}
	if c1 > 0 && c2 > 0 || x1 == x2 || y2 == y1 {
		bol = true
	}
	return
}

func Boolean40(x1, y1, x2, y2 int64) (bol bool) {
	c1 := x1 - x2
	c2 := x1 - x2
	if c1 < 0 {
		c1 = -c1
	}
	if c2 < 0 {
		c2 = -c2
	}
	if c1 == 1 && c2 == 2 || c1 == 2 && c2 == 1 {
		bol = true
	}
	return
}
