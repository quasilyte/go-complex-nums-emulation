package xmath

import "unsafe"

// These functions are also copy/pasted from the Go runtime.

var inf = float64frombits(0x7FF0000000000000)

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func abs(x float64) float64 {
	const sign = 1 << 63
	return float64frombits(float64bits(x) &^ sign)
}

// isNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func isNaN(f float64) (is bool) {
	// IEEE 754 says that only NaNs satisfy f != f.
	return f != f
}

// isInf reports whether f is an infinity.
func isInf(f float64) bool {
	return !isNaN(f) && !isFinite(f)
}

// isFinite reports whether f is neither NaN nor an infinity.
func isFinite(f float64) bool {
	return !isNaN(f - f)
}

// copysign returns a value with the magnitude
// of x and the sign of y.
func copysign(x, y float64) float64 {
	const sign = 1 << 63
	return float64frombits(float64bits(x)&^sign | float64bits(y)&sign)
}

// inf2one returns a signed 1 if f is an infinity and a signed 0 otherwise.
// The sign of the result is the sign of f.
func inf2one(f float64) float64 {
	g := 0.0
	if isInf(f) {
		g = 1.0
	}
	return copysign(g, f)
}

// Float64bits returns the IEEE 754 binary representation of f.
func float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

// Float64frombits returns the floating point number corresponding
// the IEEE 754 binary representation b.
func float64frombits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}
