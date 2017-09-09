# Go complex numbers: builtin vs library implementation

See [proposal: remove complex numbers](https://github.com/golang/go/issues/19921).

This repository:
1. Implements most `complex64` operations with user-defined type `Complex64`.
2. Measures performance of `complex64` VS `Complex64`.
3. Provides `amd64` machine code output for both builtin and `Complex64` operations.

Results and conclusions may be extrapolated to `complex128`.

**This is performance comparison of the potential implementation, not a library**.

## TL;DR

```
# old is builtin complex64
# new is user-defined Complex64

name       old time/op  new time/op  delta
Logical-4  2.89µs ± 1%  2.70µs ± 3%   -6.47%  (p=0.000 n=9+10)
Add64-4    1.50µs ± 0%  1.60µs ± 6%   +6.74%  (p=0.000 n=8+10)
Sub64-4    1.51µs ± 2%  1.70µs ± 0%  +12.48%  (p=0.000 n=10+8)
Mul64-4    6.56µs ± 2%  6.91µs ± 6%   +5.45%  (p=0.002 n=10+10)
Div64-4    53.2µs ± 3%  44.6µs ± 0%  -16.06%  (p=0.000 n=10+9)
```

> Note: it is strange that Sub/Add/Mul time differs as their
> machine code is exactly the same on amd64 (see "Machine code comparison").

```
go test -bench=.
BenchmarkLogicalBuiltin-4   	  500000	      3087 ns/op
BenchmarkLogical-4          	  500000	      2698 ns/op
BenchmarkAdd64Builtin-4     	 1000000	      1496 ns/op
BenchmarkAdd64-4            	 1000000	      1501 ns/op
BenchmarkSub64Builtin-4     	 1000000	      1495 ns/op
BenchmarkSub64-4            	 1000000	      1511 ns/op
BenchmarkMul64Builtin-4     	  200000	      6466 ns/op
BenchmarkMul64-4            	  200000	      6451 ns/op
BenchmarkDiv64Builtin-4     	   30000	     52530 ns/op
BenchmarkDiv64-4            	   30000	     44477 ns/op
```

## Benchmark results.

Results stored in "benchmark-results/".
Each result has "benchmark-results/$arch/$machine_id/"
path and contains:

1. `lscpu.txt`, the result of `lscpu` invocation result;
2. `builtin.out` & `lib.out` - results of the benchmark runs;
3. `uname.txt`, optional `uname -a` invocation result;
4. `benchstat.txt` result of `benchstat builtin.out lib.out`;

## Run benchmark

[Benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) is required to
aggregate benchmark statistics.

See [run-benchmarks](run-benchmarks) bash script for detailed instruction.

## Machine code comparison

Most code looks the same, but some code compiled
to different instruction sequences.

Look inside [disasm.go](disasm.go) to inspect objdump output.

## Edge cases / limitations

### Constant semantics

[Spec:complex_numbers](https://golang.org/ref/spec#Complex_numbers): `complex`, `imag` and `real`
can be a constant value.

>  If the operands of these functions are all constants, the return value is a constant.

User-defined struct literal is never a constant.


