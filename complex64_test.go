package xmath

import (
	"math/cmplx"
	"testing"
)

// Helper functions.

func ttUnpack64Builtin(v ttValueSet) (complex64, complex64) {
	return complex(v.r1, v.i1), complex(v.r2, v.i2)
}

func ttUnpack64(v ttValueSet) (Complex64, Complex64) {
	return Complex64{r: v.r1, i: v.i1}, Complex64{r: v.r2, i: v.i2}
}

// Unit tests.

func TestComplex64Arith(t *testing.T) {
	bothNaN := func(x, y complex64) bool {
		return cmplx.IsNaN(complex128(x)) &&
			cmplx.IsNaN(complex128(y))
	}

	tests := []struct {
		name      string
		builtinOp func(x, y complex64) complex64
		op        func(x, y Complex64) Complex64
	}{
		{
			"+",
			func(x, y complex64) complex64 { return x + y },
			Complex64.Add,
		},
		{
			"-",
			func(x, y complex64) complex64 { return x - y },
			Complex64.Sub,
		},
		{
			"*",
			func(x, y complex64) complex64 { return x * y },
			Complex64.Mul,
		},
		{
			"/",
			func(x, y complex64) complex64 { return x / y },
			Complex64.Div,
		},
	}

	for _, tt := range tests {
		for _, v := range ttValues {
			x, y := ttUnpack64Builtin(v)
			want := tt.builtinOp(x, y)
			res := tt.op(ttUnpack64(v))
			have := complex(res.r, res.i)
			if want != have && !bothNaN(want, have) {
				t.Errorf(
					"`%v%s%v` failed;\nwant: %v\nhave: %v",
					x, tt.name, y, want, have,
				)
			}
		}
	}
}

func TestComplex64Logical(t *testing.T) {
	tests := []struct {
		name      string
		builtinOp func(x, y complex64) bool
		op        func(x, y Complex64) bool
	}{
		{
			"==",
			func(x, y complex64) bool { return x == y },
			Complex64.Eq,
		},
		{
			"!=",
			func(x, y complex64) bool { return x != y },
			Complex64.Neq,
		},
	}

	for _, tt := range tests {
		for _, v := range ttValues {
			x, y := ttUnpack64Builtin(v)
			want := tt.builtinOp(x, y)
			have := tt.op(ttUnpack64(v))
			if want != have {
				t.Errorf(
					"`%v%s%v` failed;\nwant: %v\nhave: %v",
					x, tt.name, y, want, have,
				)
			}
		}
	}
}

func TestComplex64IsZero(t *testing.T) {
	bothZeroBuiltin := func(x, y complex64) bool {
		return x == 0 && y == 0
	}
	bothZero := func(x, y Complex64) bool {
		return x.IsZero() && y.IsZero()
	}

	for _, v := range ttValues {
		x, y := ttUnpack64Builtin(v)
		want := bothZeroBuiltin(x, y)
		have := bothZero(ttUnpack64(v))
		if want != have {
			t.Errorf(
				"`iszero(%v) && iszero(%v)` failed;\nwant: %v\nhave: %v",
				x, y, want, have,
			)
		}
	}
}

// Performance tests.

// Variables that used to add side-effects for tests.
var (
	ttReal32 float32
	ttImag32 float32
	ttBool1  bool
	ttBool2  bool
	ttBool3  bool
	ttBool4  bool
)

func benchBuiltin(n int, fn func(x, y complex64)) {
	for i := 0; i < n; i++ {
		for _, v := range ttValues {
			fn(ttUnpack64Builtin(v))
		}
	}
}

func bench(n int, fn func(x, y Complex64)) {
	for i := 0; i < n; i++ {
		for _, v := range ttValues {
			fn(ttUnpack64(v))
		}
	}
}

func BenchmarkLogicalBuiltin(b *testing.B) {
	benchBuiltin(b.N, func(x, y complex64) {
		ttBool1 = x == 0
		ttBool2 = y == 0
		ttBool3 = x == y
		ttBool4 = x != y
	})
}

func BenchmarkLogical(b *testing.B) {
	bench(b.N, func(x, y Complex64) {
		ttBool1 = x.IsZero()
		ttBool2 = y.IsZero()
		ttBool3 = x.Eq(y)
		ttBool4 = x.Neq(y)
	})
}

func BenchmarkAdd64Builtin(b *testing.B) {
	benchBuiltin(b.N, func(x, y complex64) {
		ttReal32 = real(x + y + x + y)
		ttImag32 = imag(y + y + y + y)
	})
}

func BenchmarkAdd64(b *testing.B) {
	bench(b.N, func(x, y Complex64) {
		ttReal32 = x.Add(y).Add(x).Add(y).Real()
		ttImag32 = y.Add(y).Add(y).Add(y).Imag()
	})
}

func BenchmarkSub64Builtin(b *testing.B) {
	benchBuiltin(b.N, func(x, y complex64) {
		ttReal32 = real(x - y - x - y)
		ttImag32 = imag(y - y - y - y)
	})
}

func BenchmarkSub64(b *testing.B) {
	bench(b.N, func(x, y Complex64) {
		ttReal32 = x.Sub(y).Sub(x).Sub(y).Real()
		ttImag32 = y.Sub(y).Sub(y).Sub(y).Imag()
	})
}

func BenchmarkMul64Builtin(b *testing.B) {
	benchBuiltin(b.N, func(x, y complex64) {
		ttReal32 = real(x * y * x * y)
		ttImag32 = imag(y * y * y * y)
	})
}

func BenchmarkMul64(b *testing.B) {
	bench(b.N, func(x, y Complex64) {
		ttReal32 = x.Mul(y).Mul(x).Mul(y).Real()
		ttImag32 = y.Mul(y).Mul(y).Mul(y).Imag()
	})
}

func BenchmarkDiv64Builtin(b *testing.B) {
	benchBuiltin(b.N, func(x, y complex64) {
		ttReal32 = real(x / y / x / y)
		ttImag32 = imag(y / y / y / y)
	})
}

func BenchmarkDiv64(b *testing.B) {
	bench(b.N, func(x, y Complex64) {
		ttReal32 = x.Div(y).Div(x).Div(y).Real()
		ttImag32 = y.Div(y).Div(y).Div(y).Imag()
	})
}
