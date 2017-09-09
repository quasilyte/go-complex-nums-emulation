package xmath

// Complex64 implements Go builtin "complex64" type.
//
// This type has value semantics, all operations return a
// new instance of Complex64.
type Complex64 struct {
	r float32
	i float32
}

// Real returns complex number real part.
func (c Complex64) Real() float32 { return c.r }

// Imag returns complex number imaginary part.
func (c Complex64) Imag() float32 { return c.i }

// IsZero returns true if both c.Real() and c.Imag() return 0.
func (c Complex64) IsZero() bool {
	return c.r == 0 && c.i == 0
}

// Eq is "==" operation.
func (c Complex64) Eq(x Complex64) bool {
	return c.r == x.r && c.i == x.i
}

// Neq is "!=" operation.
func (c Complex64) Neq(x Complex64) bool {
	return c.r != x.r || c.i != x.i
}

// Add is "+" operation.
func (c Complex64) Add(x Complex64) Complex64 {
	return Complex64{
		r: c.r + x.r,
		i: c.i + x.i,
	}
}

// Sub is "-" operation.
func (c Complex64) Sub(x Complex64) Complex64 {
	return Complex64{
		r: c.r - x.r,
		i: c.i - x.i,
	}
}

// Mul is "*" operation.
func (c Complex64) Mul(x Complex64) Complex64 {
	r1 := float64(c.r)
	i1 := float64(c.i)
	r2 := float64(x.r)
	i2 := float64(x.i)
	return Complex64{
		r: float32(r1*r2 - i1*i2),
		i: float32(r1*i2 + i1*r2),
	}
}

// Div is "/" operation.
func (c Complex64) Div(x Complex64) Complex64 {
	// The implementation code taken from Go runtime package,
	// "complex128div" function.
	// More borrowed code at "xruntime.go".

	r1 := float64(c.r)
	i1 := float64(c.i)
	r2 := float64(x.r)
	i2 := float64(x.i)

	var e, f float64 // complex(e, f) = n/m

	// Algorithm for robust complex division as described in
	// Robert L. Smith: Algorithm 116: Complex division. Commun. ACM 5(8): 435 (1962).
	if abs(r2) >= abs(i2) {
		ratio := i2 / r2
		denom := r2 + ratio*i2
		e = (r1 + i1*ratio) / denom
		f = (i1 - r1*ratio) / denom
	} else {
		ratio := r2 / i2
		denom := i2 + ratio*r2
		e = (r1*ratio + i1) / denom
		f = (i1*ratio - r1) / denom
	}

	if isNaN(e) && isNaN(f) {
		// Correct final result to infinities and zeros if applicable.
		// Matches C99: ISO/IEC 9899:1999 - G.5.1  Multiplicative operators.

		a, b := r1, i1
		c, d := r2, i2

		switch {
		case x.IsZero() && (!isNaN(a) || !isNaN(b)):
			e = copysign(inf, c) * a
			f = copysign(inf, c) * b

		case (isInf(a) || isInf(b)) && isFinite(c) && isFinite(d):
			a = inf2one(a)
			b = inf2one(b)
			e = inf * (a*c + b*d)
			f = inf * (b*c - a*d)

		case (isInf(c) || isInf(d)) && isFinite(a) && isFinite(b):
			c = inf2one(c)
			d = inf2one(d)
			e = 0 * (a*c + b*d)
			f = 0 * (b*c - a*d)
		}
	}

	return Complex64{r: float32(e), i: float32(f)}
}
