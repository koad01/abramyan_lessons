package if20_30

import (
	"fmt"
	"testing"
)

func TestIf20(t *testing.T) {
	type dataForBegin15 struct {
		a     float64
		b     float64
		c     float64
		wantS string
		wantR float64
	}
	var tests = []dataForBegin15{
		{8, 6, 12, "b", 2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v, %v, %v, %v", tt.a, tt.b, tt.c, tt.wantS, tt.wantS)
		functionForTest := func(t *testing.T) {
			ansS, ansR := If20(tt.a, tt.b, tt.c)
			if ansS != tt.wantS {
				t.Errorf("got %v, wantS %v", ansS, tt.wantS)
			}
			if ansR != tt.wantR {
				t.Errorf("got %v, wantR %v", ansR, tt.wantR)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf21(t *testing.T) {
	type dataForBegin15 struct {
		x     float64
		y     float64
		wantS int16
	}
	var tests = []dataForBegin15{
		{8, 6, 3},
		{0, 0, 0},
		{1, 0, 1},
		{0, 6, 2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v, %v", tt.x, tt.y, tt.wantS)
		functionForTest := func(t *testing.T) {
			ansS := If21(tt.x, tt.y)
			if ansS != tt.wantS {
				t.Errorf("got %v, wantS %v", ansS, tt.wantS)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf22(t *testing.T) {
	type dataForBegin15 struct {
		x     float64
		y     float64
		wantS int16
	}
	var tests = []dataForBegin15{
		{8, 6, 1},
		{2, -1, 4},
		{-9, 3, 2},
		{-7, -5, 3},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v, %v", tt.x, tt.y, tt.wantS)
		functionForTest := func(t *testing.T) {
			ansS := If22(tt.x, tt.y)
			if ansS != tt.wantS {
				t.Errorf("got %v, wantS %v", ansS, tt.wantS)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf23(t *testing.T) {
	type dataForBegin15 struct {
		x1     int16
		x2     int16
		x3     int16
		y1     int16
		y2     int16
		y3     int16
		wantX4 int16
		wantY4 int16
	}
	var tests = []dataForBegin15{
		{1, 1, 2, 1, 2, 2, 1, 2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v, %v", tt.x1, tt.x2, tt.x3, tt.y1, tt.y2, tt.y3, tt.wantX4, tt.wantY4)
		functionForTest := func(t *testing.T) {
			ansX4, ansY4 := If23(tt.x1, tt.x2, tt.x3, tt.y1, tt.y2, tt.y3)
			if ansX4 != tt.wantX4 {
				t.Errorf("got %v, wantX4 %v", ansX4, tt.wantX4)
			}
			if ansY4 != tt.wantY4 {
				t.Errorf("got %v, wantY4 %v", ansY4, tt.wantY4)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf24(t *testing.T) {
	type dataForBegin15 struct {
		x     float64
		wantF float64
	}
	var tests = []dataForBegin15{
		{8, 1.9787164932467634},
		{-2, 8},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.x, tt.wantF)
		functionForTest := func(t *testing.T) {
			ansF := If24(tt.x)
			if ansF != tt.wantF {
				t.Errorf("got %v, wantF %v", ansF, tt.wantF)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf25(t *testing.T) {
	type dataForBegin15 struct {
		x     int16
		wantF int16
	}
	var tests = []dataForBegin15{
		{8, 16},
		{-2, 6},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.x, tt.wantF)
		functionForTest := func(t *testing.T) {
			ansF := If25(tt.x)
			if ansF != tt.wantF {
				t.Errorf("got %v, wantF %v", ansF, tt.wantF)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf26(t *testing.T) {
	type dataForBegin15 struct {
		x     float64
		wantF float64
	}
	var tests = []dataForBegin15{
		{8, 64},
		{-2, 2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.x, tt.wantF)
		functionForTest := func(t *testing.T) {
			ansF := If26(tt.x)
			if ansF != tt.wantF {
				t.Errorf("got %v, wantF %v", ansF, tt.wantF)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf27(t *testing.T) {
	type dataForBegin15 struct {
		x     float64
		wantF float64
	}
	var tests = []dataForBegin15{
		{8, 1},
		{-2, 0},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.x, tt.wantF)
		functionForTest := func(t *testing.T) {
			ansF := If27(tt.x)
			if ansF != tt.wantF {
				t.Errorf("got %v, wantF %v", ansF, tt.wantF)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf28(t *testing.T) {
	type dataForBegin15 struct {
		x     int16
		wantF int16
	}
	var tests = []dataForBegin15{
		{8, 366},
		{678, 365},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.x, tt.wantF)
		functionForTest := func(t *testing.T) {
			ansF := If28(tt.x)
			if ansF != tt.wantF {
				t.Errorf("got %v, wantF %v", ansF, tt.wantF)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf29(t *testing.T) {
	type dataForBegin15 struct {
		x       int16
		wantXar string
	}
	var tests = []dataForBegin15{
		{8, "Положительное четное"},
		{671, "Положительное нечетное"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.x, tt.wantXar)
		functionForTest := func(t *testing.T) {
			ansXar := If29(tt.x)
			if ansXar != tt.wantXar {
				t.Errorf("got %v, wantXar %v", ansXar, tt.wantXar)
			}
		}
		t.Run(testname, functionForTest)

	}
}

func TestIf30(t *testing.T) {
	type dataForBegin15 struct {
		x       int16
		wantXar string
	}
	var tests = []dataForBegin15{
		{8, "Четное однозначное"},
		{671, "Нечетное трёхзначное"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.x, tt.wantXar)
		functionForTest := func(t *testing.T) {
			ansXar := If30(tt.x)
			if ansXar != tt.wantXar {
				t.Errorf("got %v, wantXar %v", ansXar, tt.wantXar)
			}
		}
		t.Run(testname, functionForTest)

	}
}
