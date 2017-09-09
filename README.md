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
Logical-4  2.89µs ± 1%  2.77µs ± 8%    ~     (p=0.108 n=9+10)
Add64-4    1.50µs ± 1%  1.51µs ± 1%    ~     (p=0.174 n=9+10)
Sub64-4    1.50µs ± 1%  1.52µs ± 1%  +0.80%  (p=0.000 n=9+10)
Mul64-4    6.50µs ± 2%  6.74µs ± 7%    ~     (p=0.101 n=10+10)
Div64-4    52.6µs ± 0%  47.9µs ± 7%  -8.99%  (p=0.000 n=9+10)
```

> Note: it is strange that Sub/Add/Mul time differs as their
> machine code is exactly the same on amd64 (see "Machine code comparison").

```
go test -bench=.
BenchmarkLogicalBuiltin-4   	  500000	      2876 ns/op
BenchmarkLogical-4          	  500000	      2711 ns/op
BenchmarkAdd64Builtin-4     	 1000000	      1539 ns/op
BenchmarkAdd64-4            	 1000000	      1516 ns/op
BenchmarkSub64Builtin-4     	 1000000	      1511 ns/op
BenchmarkSub64-4            	 1000000	      1568 ns/op
BenchmarkMul64Builtin-4     	  200000	      6476 ns/op
BenchmarkMul64-4            	  200000	      6456 ns/op
BenchmarkDiv64Builtin-4     	   30000	     52589 ns/op
BenchmarkDiv64-4            	   30000	     44854 ns/op
```

## Benchmark results.

Results stored in "benchmark-results/".
Structure is "benchmark-results/$arch/$machine_id/$version".

Each machine folder can contain some hardware/system info:

- `lscpu.txt`, the result of `lscpu` invocation result
- `uname.txt`, optional `uname -a` invocation result

For each version that was tested on that machine, there is "v$number"
folder, which consist of:

- `builtin.out` & `lib.out` - results of the benchmark runs
- `benchstat.txt` result of `benchstat builtin.out lib.out`

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

## Versions

* 1 : up to 1c7348b2f4683432625a7e1c9c9b434fd48b6ad7
  initial implementation
* 2 : from 41112b4215a2d02aca055a751f8e4047168397c5
  [alternative Eq/Neq/IsZero implementation](https://github.com/Quasilyte/go-complex-nums-emulation/issues/2)