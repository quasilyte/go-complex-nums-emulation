# Go complex numbers: builtin vs library implementation

See [proposal: remove complex numbers](https://github.com/golang/go/issues/19921).

This repository:
1. Implements (quite naively) most `complex64`/`complex128` operations with user-defined
   types `Complex64` and `Complex128`.
2. Measures performance of `complex64` VS `Complex64` and `complex128` VS `Complex128`. 

Tests are included.
