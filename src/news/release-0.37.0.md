---
title: "Go HEP release 0.37.0"
date: 2025-05-20T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.15470934.svg)](https://doi.org/10.5281/zenodo.15470934)


Release [`v0.37.0`](https://github.com/go-hep/hep/tree/v0.37.0) is out of the oven.

This release contains one major API breakage : we've dropped support for [Gio](https://gioui.org) (following suit of `gonum/plot`), and thus removed `pawgo` and `hplot/cmd/iplot`.
We _may_ reintroduce `pawgo` and `iplot` some time in the future, perhaps based off [modernc.org/tk9.0](https://modernc.org/tk9.0).

`Go-HEP v0.37.0` is the _last_ release to be cooked with `GitHub`.
[Go HEP](https://go-hep.org) is indeed migrating to [codeberg.org/go-hep](https://codeberg.org/go-hep).

## hplot

- `hplot/cmd/iplot`: command removed, following removal of `Gio` support
- add `PreSteps`, `MidSteps` and `PostSteps` options for drawing steps with a `S2D` plotter

## fit

- add [`Poly`](https://pkg.go.dev/go-hep.org/x/hep@v0.37.0/fit#Poly) to fit polynomials
- add `Func1D.Hessian` to (numerically) compute the Hessian matrix at a given point in the parameter space

## groot

- added support for ROOT `6.34/04`

## Changelog

* [ad8e1bdc](/commit/ad8e1bdc) all: bump x/deps
* [6be6485b](/commit/6be6485b) ci: add gitlab-ci support
* [9e79199c](/commit/9e79199c) ci: migrate to woodpecker
* [f1741a87](/commit/f1741a87) all: use codeberg.org/raw instead of github.com/raw
* [f5a6a573](/commit/f5a6a573) hplot: add pre/mid/post-steps kinds to S2D
* [7ca54edc](/commit/7ca54edc) fit: add Poly to fit polynomials
* [0591db64](/commit/0591db64) fit: export Func1D.Hessian
* [02e76ba8](/commit/02e76ba8) all: bump x/net@v0.38
* [b6d84be2](/commit/b6d84be2) all: adapt tests and floating-point precision for Mac-Silicon
* [9d9ea81e](/commit/9d9ea81e) all: migrate to math/rand/v2
* [5129a0f8](/commit/5129a0f8) all: bump Gonum@v0.16
* [542b178e](/commit/542b178e) all: migrate to codeberg.org/gonuts/binary
* [2c315900](/commit/2c315900) all: migrate to codeberg.org/sbinet/npyio
* [d1e4d816](/commit/d1e4d816) all: migrate to codeberg.org/gonuts/commander
* [0a46d44b](/commit/0a46d44b) all: migrate to codeberg.org/go-mmap
* [2087e2f1](/commit/2087e2f1) all: migrate to codeberg.org/astrogo
* [4c381da6](/commit/4c381da6) all: bump compress@v1.18, x/{crypto,image,sync,mod,net,sys,text,tools} and modernc.org/ql
* [fc7e544b](/commit/fc7e544b) all: modernize Go usage
* [f182f98b](/commit/f182f98b) sliceop: generalize Filter
* [132fff9d](/commit/132fff9d) sliceop/f64s: remove
* [8f7f39b2](/commit/8f7f39b2) xrootd/xrdio: ensure File implements io/fs.File
* [5b304e15](/commit/5b304e15) all: bump deps
* [34c208fc](/commit/34c208fc) ci: remove old travis CI file
* [757a1994](/commit/757a1994) ci: improve run-tests command line interface
* [a7939635](/commit/a7939635) groot: add test for ROOT/C++ and CI
* [f33e49e2](/commit/f33e49e2) groot/rtree: adapt test to TTree::Scan output for ROOT-6.34
* [e2a54ecc](/commit/e2a54ecc) ci: bump ubuntu version to 24.04
* [39398787](/commit/39398787) groot: implement TFormula-v14
* [f52618e1](/commit/f52618e1) groot: bump to ROOT-6.34.04
* [c8e8d96a](/commit/c8e8d96a) all: bump sbinet/go-arrow@v0.3.0
* [233f253d](/commit/233f253d) all: drop Go-1.22, add Go-1.24
* [8caf0ce3](/commit/8caf0ce3) all: remove gioui dependency
* [3c86fe7a](/commit/3c86fe7a) pawgo: remove
* [bb0e3b75](/commit/bb0e3b75) hplot/cmd/iplot: remove
* [a9722bcf](/commit/a9722bcf) ci: remove iplot and pawgo special handling
* [1bd95a81](/commit/1bd95a81) ci: remove cgo deps needed for vg/vggio
* [9b7d6b68](/commit/9b7d6b68) ci: bump staticcheck@2025.1
* [daadb188](/commit/daadb188) ci: use ubuntu-latest
* [633172f5](/commit/633172f5) fit: remove debugging remnant
* [ce2a9ced](/commit/ce2a9ced) all: bump golang.org/x deps
