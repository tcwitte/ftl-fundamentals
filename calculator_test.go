package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	desc        string
	a, b, want  float64
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{
			desc: "positive",
			a:    2,
			b:    2,
			want: 4,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Add(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{
			desc: "positive",
			a:    4,
			b:    2,
			want: 2,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Subtract(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{
			desc: "positive",
			a:    2,
			b:    2,
			want: 4,
		},
		{
			desc: "negative",
			a:    -2,
			b:    2,
			want: -4,
		},
		{
			desc: "fraction",
			a:    0.5,
			b:    2,
			want: 1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Multiply(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{
			desc: "positive",
			a:    4,
			b:    2,
			want: 2,
		},
		{
			desc: "negative",
			a:    -4,
			b:    2,
			want: -2,
		},
		{
			desc: "fraction",
			a:    4.6,
			b:    2,
			want: 2.3,
		},
		{
			desc:        "division by zero",
			a:           4.6,
			b:           0,
			errExpected: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Divide(tC.a, tC.b)
			if (err != nil) != tC.errExpected {
				t.Errorf("error expected %v, error %s", tC.errExpected, err)
			}
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{
			desc: "positive",
			a:    4,
			want: 2,
		},
		{
			desc:        "negative",
			a:           -4,
			errExpected: true,
		},
		{
			desc: "identity",
			a:    1,
			want: 1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Sqrt(tC.a)
			if (err != nil) != tC.errExpected {
				t.Errorf("error expected %v, error %s", tC.errExpected, err)
			}
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}
