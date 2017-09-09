// +build ignore

package xmath

// This file defines functions for output asm inspection.
// Functions are single-line to make it possible to grep
// build -S output by line number.
//
// `go build -gcflags -S xruntime.go complex64.go disasm.go 2>&1 | grep 'disasm.go:LINE' | awk '{$1=$2=$3="";print $0}'`
// OR
// `go build -o a.out xruntime.go complex64.go disasm.go` + `go tool objdump -s FUNC_NAME a.out | awk '{$1=$3=""; print $0}'`
//

// 0x22b4  REP MOVSS 0x8(SP), X0
// 0x22ba  REP MOVSS X0, 0x10(SP)
// 0x22c0  RET
func readReal64builtin(c complex64) float32 { return real(c) }

// 0x22ca  REP MOVSS 0x8(SP), X0
// 0x22d0  REP MOVSS X0, 0x10(SP)
// 0x22d6  RET
func readReal64(c Complex64) float32 { return c.Real() }

// 0x22e0  REP MOVSS 0xc(SP), X0
// 0x22e6  REP MOVSS X0, 0x10(SP)
// 0x22ec  RET
func readImag64builtin(c complex64) float32 { return imag(c) }

// 0x22f6  REP MOVSS 0xc(SP), X0
// 0x22fc  REP MOVSS X0, 0x10(SP)
// 0x2302  RET
func readImag64(c Complex64) float32 { return c.Imag() }

// 0x21ae  REP MOVSS 0x10(SP), X0
// 0x21b4  REP MOVSS 0x8(SP), X1
// 0x21ba  REP ADDSS X1, X0
// 0x21be  REP MOVSS X0, 0x18(SP)
// 0x21c4  REP MOVSS 0xc(SP), X0
// 0x21ca  REP MOVSS 0x14(SP), X1
// 0x21d0  REP ADDSS X1, X0
// 0x21d4  REP MOVSS X0, 0x1c(SP)
// 0x21da  RET
func add64builtin(c1, c2 complex64) complex64 { return c1 + c2 }

// 0x21e4  REP MOVSS 0x8(SP), X0
// 0x21ea  REP MOVSS 0x10(SP), X1
// 0x21f0  REP ADDSS X1, X0
// 0x21f4  REP MOVSS X0, 0x18(SP)
// 0x21fa  REP MOVSS 0xc(SP), X0
// 0x2200  REP MOVSS 0x14(SP), X1
// 0x2206  REP ADDSS X1, X0
// 0x220a  REP MOVSS X0, 0x1c(SP)
// 0x2210  RET
func add64(c1, c2 Complex64) Complex64 { return c1.Add(c2) }

// 0x2416  REP MOVSS 0x10(SP), X0
// 0x241c  REP MOVSS 0x8(SP), X1
// 0x2422  REP ADDSS X1, X0
// 0x2426  REP MOVSS 0x18(SP), X1
// 0x242c  REP ADDSS X1, X0
// 0x2430  REP MOVSS X0, 0x20(SP)
// 0x2436  REP MOVSS 0x14(SP), X0
// 0x243c  REP MOVSS 0xc(SP), X1
// 0x2442  REP ADDSS X1, X0
// 0x2446  REP MOVSS 0x1c(SP), X1
// 0x244c  REP ADDSS X1, X0
// 0x2450  REP MOVSS X0, 0x24(SP)
// 0x2456  RET
func chainedAdd64builtin(c1, c2, c3 complex64) complex64 { return c1 + c2 + c3 }

// 0x2460  REP MOVSS 0x10(SP), X0
// 0x2466  REP MOVSS 0x8(SP), X1
// 0x246c  REP ADDSS X1, X0
// 0x2470  REP MOVSS 0x18(SP), X1
// 0x2476  REP ADDSS X1, X0
// 0x247a  REP MOVSS X0, 0x20(SP)
// 0x2480  REP MOVSS 0xc(SP), X0
// 0x2486  REP MOVSS 0x14(SP), X1
// 0x248c  REP ADDSS X1, X0
// 0x2490  REP MOVSS 0x1c(SP), X1
// 0x2496  REP ADDSS X1, X0
// 0x249a  REP MOVSS X0, 0x24(SP)
// 0x24a0  RET
func chainedAdd64(c1, c2, c3 Complex64) Complex64 { return c1.Add(c2).Add(c3) }

// 0x221a  REP MOVSS 0x8(SP), X0
// 0x2220  REP MOVSS 0x10(SP), X1
// 0x2226  REP SUBSS X1, X0
// 0x222a  REP MOVSS X0, 0x18(SP)
// 0x2230  REP MOVSS 0xc(SP), X0
// 0x2236  REP MOVSS 0x14(SP), X1
// 0x223c  REP SUBSS X1, X0
// 0x2240  REP MOVSS X0, 0x1c(SP)
// 0x2246  RET
func sub64builtin(c1, c2 complex64) complex64 { return c1 - c2 }

// 0x2250  REP MOVSS 0x8(SP), X0
// 0x2256  REP MOVSS 0x10(SP), X1
// 0x225c  REP SUBSS X1, X0
// 0x2260  REP MOVSS X0, 0x18(SP)
// 0x2266  REP MOVSS 0xc(SP), X0
// 0x226c  REP MOVSS 0x14(SP), X1
// 0x2272  REP SUBSS X1, X0
// 0x2276  REP MOVSS X0, 0x1c(SP)
// 0x227c  RET
func sub64(c1, c2 Complex64) Complex64 { return c1.Sub(c2) }

// 0x2286  REP MOVSS 0x8(SP), X0
// 0x228c  REP CVTSS2SD X0, X0
// 0x2290  REP MOVSS 0x10(SP), X1
// 0x2296  REP CVTSS2SD X1, X1
// 0x229a  MOVUPS X0, X2
// 0x229d  REPNE MULSD X1, X0
// 0x22a1  REP MOVSS 0xc(SP), X3
// 0x22a7  REP CVTSS2SD X3, X3
// 0x22ab  REP MOVSS 0x14(SP), X4
// 0x22b1  REP CVTSS2SD X4, X4
// 0x22b5  MOVUPS X3, X5
// 0x22b8  REPNE MULSD X4, X3
// 0x22bc  REPNE SUBSD X3, X0
// 0x22c0  REPNE CVTSD2SS X0, X0
// 0x22c4  REP MOVSS X0, 0x18(SP)
// 0x22ca  REPNE MULSD X4, X2
// 0x22ce  REPNE MULSD X5, X1
// 0x22d2  REPNE ADDSD X1, X2
// 0x22d6  REPNE CVTSD2SS X2, X0
// 0x22da  REP MOVSS X0, 0x1c(SP)
// 0x22e0  RET
func mul64builtin(c1, c2 complex64) complex64 { return c1 * c2 }

// 0x22eb  REP MOVSS 0x8(SP), X0
// 0x22f1  REP CVTSS2SD X0, X0
// 0x22f5  REP MOVSS 0x10(SP), X1
// 0x22fb  REP CVTSS2SD X1, X1
// 0x22ff  MOVUPS X0, X2
// 0x2302  REPNE MULSD X1, X0
// 0x2306  REP MOVSS 0xc(SP), X3
// 0x230c  REP CVTSS2SD X3, X3
// 0x2310  REP MOVSS 0x14(SP), X4
// 0x2316  REP CVTSS2SD X4, X4
// 0x231a  MOVUPS X3, X5
// 0x231d  REPNE MULSD X4, X3
// 0x2321  REPNE SUBSD X3, X0
// 0x2325  REPNE CVTSD2SS X0, X0
// 0x2329  REP MOVSS X0, 0x18(SP)
// 0x232f  REPNE MULSD X4, X2
// 0x2333  REPNE MULSD X1, X5
// 0x2337  REPNE ADDSD X5, X2
// 0x233b  REPNE CVTSD2SS X2, X0
// 0x233f  REP MOVSS X0, 0x1c(SP)
// 0x2345  RET
func mul64(c1, c2 Complex64) Complex64 { return c1.Mul(c2) }

// 0x2350  REP MOVSS 0x8(SP), X0
// 0x2356  XORPS X1, X1
// 0x2359  UCOMISS X1, X0
// 0x235c  SETE CL
// 0x235f  SETNP AL
// 0x2362  ANDL AX, CX
// 0x2364  REP MOVSS 0xc(SP), X0
// 0x236a  UCOMISS X1, X0
// 0x236d  SETE DL
// 0x2370  SETNP AL
// 0x2373  ANDL AX, DX
// 0x2375  ANDL DX, CX
// 0x2377  MOVB CL, 0x10(SP)
// 0x237b  RET
func isZero64builtin(c complex64) bool { return c == 0 }

// 0x25d6  REP MOVSS 0x8(SP), X0
// 0x25dc  XORPS X1, X1
// 0x25df  UCOMISS X1, X0
// 0x25e2  JNE 0x25fc
// 0x25e4  JP 0x25fc
// 0x25e6  REP MOVSS 0xc(SP), X0
// 0x25ec  UCOMISS X1, X0
// 0x25ef  SETE CL
// 0x25f2  SETNP AL
// 0x25f5  ANDL AX, CX
// 0x25f7  MOVB CL, 0x10(SP)
// 0x25fb  RET
// 0x25fc  XORL CX, CX
// 0x25fe  JMP 0x25f7
func isZero64(c Complex64) bool { return c.IsZero() }

// 0x23ba  REP MOVSS 0x8(SP), X0
// 0x23c0  REP MOVSS 0x10(SP), X1
// 0x23c6  UCOMISS X1, X0
// 0x23c9  SETE CL
// 0x23cc  SETNP AL
// 0x23cf  ANDL AX, CX
// 0x23d1  REP MOVSS 0xc(SP), X0
// 0x23d7  REP MOVSS 0x14(SP), X1
// 0x23dd  UCOMISS X1, X0
// 0x23e0  SETE DL
// 0x23e3  SETNP AL
// 0x23e6  ANDL AX, DX
// 0x23e8  ANDL DX, CX
// 0x23ea  MOVB CL, 0x18(SP)
// 0x23ee  RET
func eq64builtin(c1, c2 complex64) bool { return c1 == c2 }

// 0x2649  REP MOVSS 0x8(SP), X0
// 0x264f  REP MOVSS 0x10(SP), X1
// 0x2655  UCOMISS X1, X0
// 0x2658  JNE 0x2678
// 0x265a  JP 0x2678
// 0x265c  REP MOVSS 0xc(SP), X0
// 0x2662  REP MOVSS 0x14(SP), X1
// 0x2668  UCOMISS X1, X0
// 0x266b  SETE CL
// 0x266e  SETNP AL
// 0x2671  ANDL AX, CX
// 0x2673  MOVB CL, 0x18(SP)
// 0x2677  RET
// 0x2678  XORL CX, CX
// 0x267a  JMP 0x2673
func eq64(c1, c2 Complex64) bool { return c1.Eq(c2) }

// 0x2480  REP MOVSS 0x8(SP), X0
// 0x2486  REP MOVSS 0x10(SP), X1
// 0x248c  UCOMISS X1, X0
// 0x248f  SETE CL
// 0x2492  SETNP AL
// 0x2495  ANDL AX, CX
// 0x2497  REP MOVSS 0xc(SP), X0
// 0x249d  REP MOVSS 0x14(SP), X1
// 0x24a3  UCOMISS X1, X0
// 0x24a6  SETE DL
// 0x24a9  SETNP AL
// 0x24ac  ANDL AX, DX
// 0x24ae  ANDL DX, CX
// 0x24b0  XORL $0x1, CX
// 0x24b3  MOVB CL, 0x18(SP)
// 0x24b7  RET
func neq64builtin(c1, c2 complex64) bool { return c1 != c2 }

// 0x26c8  REP MOVSS 0x8(SP), X0
// 0x26ce  REP MOVSS 0x10(SP), X1
// 0x26d4  UCOMISS X1, X0
// 0x26d7  JNE 0x26db
// 0x26d9  JNP 0x26e5
// 0x26db  MOVL $0x1, AX
// 0x26e0  MOVB AL, 0x18(SP)
// 0x26e4  RET
// 0x26e5  REP MOVSS 0xc(SP), X0
// 0x26eb  REP MOVSS 0x14(SP), X1
// 0x26f1  UCOMISS X1, X0
// 0x26f4  SETNE CL
// 0x26f7  SETP AL
// 0x26fa  ORL AX, CX
// 0x26fc  MOVL CX, AX
// 0x26fe  JMP 0x26e0
func neq64(c1, c2 Complex64) bool { return c1.Neq(c2) }
