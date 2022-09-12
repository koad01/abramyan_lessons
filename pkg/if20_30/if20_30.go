package if20_30

import "math"

func If20(a, b, c float64) (s string, r float64) {
	rb := a - b
	if rb < 0 {
		rb = -rb
	}
	rc := a - c
	if rc < 0 {
		rc = -rc
	}
	if rb < rc {
		s = "b"
		r = rb
	} else {
		s = "c"
		r = rc
	}
	return
}

func If21(x, y float64) (s int16) {
	if x == 0 && y == 0 {
		s = 0
	}
	if x == 0 && y != 0 {
		s = 2
	}
	if x != 0 && y == 0 {
		s = 1
	}
	if x != 0 && y != 0 {
		s = 3
	}
	return
}

func If22(x, y float64) (s int16) {
	if x > 0 && y > 0 {
		s = 1
	}
	if x < 0 && y > 0 {
		s = 2
	}
	if x < 0 && y < 0 {
		s = 3
	}
	if x > 0 && y < 0 {
		s = 4
	}
	return
}

func If23(x1, y1, x2, y2, x3, y3 int16) (x4, y4 int16) {
	if x2 == x3 {
		x4 = x1
	}
	if x3 == x1 {
		x4 = x2
	}
	if x1 == x2 {
		x4 = x3
	}
	if y2 == y3 {
		y4 = y1
	}
	if y3 == y1 {
		y4 = y2
	}
	if y1 == y2 {
		y4 = y3
	}
	return
}

func If24(x float64) (f float64) {
	if x > 0 {
		f = 2 * math.Sin(x)
	} else {
		f = 6 - x
	}
	return
}

func If25(x int16) (f int16) {
	if x < -2 || x > 2 {
		f = 2 * x
	} else {
		f = -3 * x
	}
	return
}

func If26(x float64) (f float64) {
	if x <= 0 {
		f = -x
	} else if x > 0 || x < 2 {
		f = x * x
	} else {
		f = 4
	}
	return
}

func If27(x float64) (f float64) {
	var y int = int(x)
	if x < 0 {
		f = 0
	} else if y%2 == 0 {
		f = 1
	} else if y%2 == 1 {
		f = -1
	}
	return
}

func If28(x int16) (nd int16) {
	if ((x%4) == 0) && ((x%100) != 0) || ((x % 400) == 0) {
		nd = 366
	} else {
		nd = 365
	}
	return
}

func If29(x int16) (xar string) {
	if x < 0 && -x%2 == 1 {
		xar = "Отрицательное нечетное"
	}
	if x < 0 && -x%2 == 0 {
		xar = "Отрицательное четное"
	}
	if x > 0 && x%2 == 0 {
		xar = "Положительное четное"
	}
	if x > 0 && x%2 == 1 {
		xar = "Положительное нечетное"
	}
	if x == 0 {
		xar = "Нулевое"
	}
	return
}

func If30(x int16) (xar string) {
	if x%2 == 1 && x <= 9 {
		xar = "Нечетное однозначное"
	}
	if x%2 == 0 && x <= 9 {
		xar = "Четное однозначное"
	}
	if x%2 == 1 && x <= 99 && x >= 10 {
		xar = "Нечетное двухзначное"
	}
	if x%2 == 0 && x <= 99 && x >= 10 {
		xar = "Четное двухзначное"
	}
	if x%2 == 1 && x <= 999 && x >= 100 {
		xar = "Нечетное трёхзначное"
	}
	if x%2 == 0 && x <= 999 && x >= 100 {
		xar = "Четное трёхзначное"
	}
	return
}
