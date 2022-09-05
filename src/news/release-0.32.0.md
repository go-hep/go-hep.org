---
title: "Go HEP release 0.32.0"
date: 2022-09-05T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.7050683.svg)](https://doi.org/10.5281/zenodo.7050683)

Release [`v0.32.0`](https://github.com/go-hep/hep/tree/v0.32.0) is out of the oven.

This release contains no major API breakage, but all around improvements and bug fixes.

This release drops support for Go-1.17 (and earlier) and is the first release with the introduction of a few generics-based APIs.

Thanks to Olivier Mengué ([@dolmen](https://github.com/dolmen)) there were a couple of house cleaning changes that have been applied (as part of the [CERN+Google Hackathon](https://gohack.devpost.com/)).
Thanks again Olivier.

There has been some work on `RNTuple` front, but nothing that can be released yet.

## brio

General clean-ups by Olivier M.

## fwk

General clean-ups by Olivier M.

## groot

- `riofs`: add generic `Get[T]` to retrieve ROOT objects
- `riofs`: export in-memory read-only `RMemFile` ROOT file

## hbook

General clean-ups by Olivier M.

## lhef

General clean-ups by Oliver M.
Olivier also reduced the amount of memory allocations in `lhef.NewDecoder`.

## sliceop

- introduce `sliceop` and generics-based functions that mimick what `sliceop/f64s` did
- implement `sliceop/f64s` in terms of `sliceop`

## Changelog

* [c8c6fbf9](/commit/c8c6fbf9) fmom: update for new gonum.org/v1/gonum@v0.12.0/r3 API
* [3e7183c9](/commit/3e7183c9) all: update gonum@v0.12, gokrb5@v8.4.3, compress@v1.15.9, lz4@v4.1.15, npyio@v0.7.0 and go/x
* [535619c2](/commit/535619c2) ci: add Go-1.19
* [c6ebf251](/commit/c6ebf251) all: use git.sr.ht/~sbinet/go-arrow@v0.2.0 fork
* [770e2dc1](/commit/770e2dc1) all: bump x/tools@v0.1.11
* [ece275b1](/commit/ece275b1) groot: fix go:generate portability for root-gen-type
* [0c1ce510](/commit/0c1ce510) brio-gen: cleanup files produced by test on test success
* [8c3e3fa2](/commit/8c3e3fa2) ci: bump actions/{setup-go,cache,checkout}/v3
* [fafbeb67](/commit/fafbeb67) all: bump yaml@v3.0.1
* [ff19fc7b](/commit/ff19fc7b) groot: document interop model with ROOT/C++
* [0eeae227](/commit/0eeae227) groot/rmeta: regnerate stringer
* [f9b43650](/commit/f9b43650) groot/{cmd/root-gen-type,rdict}: add support for generating RVec types
* [dc2dcf63](/commit/dc2dcf63) groot/{rcmd,rdict}: add support for reading ROOT::VecOps::RVec<T>
* [fd8c2dc4](/commit/fd8c2dc4) cmd/lhef2hepmc: improve warning message and display event number
* [f979a1cb](/commit/f979a1cb) all: apply gofmt-1.19 formating rules
* [2d8ce49a](/commit/2d8ce49a) fwk/internal/fwktest: apply staticcheck cosmetic change
* [f81a84e5](/commit/f81a84e5) fwk: gofmt
* [f596c97c](/commit/f596c97c) fwk: manual fix of the last reference of pkg testdata
* [d4ec96b7](/commit/d4ec96b7) fwk: fix references to pkg testdata
* [c04eef86](/commit/c04eef86) fwk: rename pkg ./testdata to ./internal/fwktest
* [07f22fc6](/commit/07f22fc6) sliceop: use errors.New for constant errors
* [40700b82](/commit/40700b82) hbook: regenerate files produced by brio-gen
* [25b579b8](/commit/25b579b8) hbook: fix go:generate to not require to install brio-gen
* [7fe81f63](/commit/7fe81f63) brio-gen: fix generated code to follow the Go standard
* [3396f5f1](/commit/3396f5f1) brio-gen: add .gitignore for temporary files produced by tests
* [19923699](/commit/19923699) lhef: reduce allocs on event decoding
* [9c5c5662](/commit/9c5c5662) lhef: more Decoder tests
* [cbb49350](/commit/cbb49350) lhef: reduce allocs in NewDecoder
* [163d127b](/commit/163d127b) lhef: use errors.New instead of fmt.Errorf for constants
* [020835d7](/commit/020835d7) cmd/root2yoda: removes unreacheable statement
* [cf9adbb1](/commit/cf9adbb1) ci: use codecov-action@v2
* [474d07ad](/commit/474d07ad) groot/riofs: export in-memory read-only RMemFile ROOT file
* [bb8ad3d4](/commit/bb8ad3d4) groot/{cmd/root-gen-type,riofs}: add testdata for base+derived classes
* [c52e43d3](/commit/c52e43d3) groot/rdict: correctly handle generation of marshaling for base classes
* [7f90bc7e](/commit/7f90bc7e) all: bump x/crypto, x/exp and x/sys
* [4f927c35](/commit/4f927c35) sliceop: improve performances of Take
* [c2ec83ed](/commit/c2ec83ed) sliceop{,/f64s}: introduce generics sliceop
* [582dce6c](/commit/582dce6c) groot/riofs: add generic Get[T] to retrieve ROOT objects
* [ccbaeb33](/commit/ccbaeb33) all: apply staticcheck fixes
* [e93f8a17](/commit/e93f8a17) ci: use staticcheck instead of golangci-lint
* [7a3ff60c](/commit/7a3ff60c) ci: reduce git-checkout depth
* [ddf2dc75](/commit/ddf2dc75) all: drop Go-1.17

