+++
date = "2016-11-24T09:19:32+01:00"
title = "Welcome to Go-HEP"
type = "index"
weight = 1
tags = ["sphinx","documentation"]

+++

`Go-HEP` is a set of libraries and applications allowing physicists from High Energy Physics (HEP) to write efficient analysis code in the [Go](https://golang.org) programming language.


[![GitHub release](https://img.shields.io/github/release/go-hep/hep.svg)](https://github.com/go-hep/hep/releases)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/go-hep.org/x/hep)
[![CI](https://github.com/go-hep/hep/workflows/CI/badge.svg)](https://github.com/go-hep/hep/actions)
[![Build Status](https://secure.travis-ci.org/go-hep/hep.png)](http://travis-ci.org/go-hep/hep)
[![appveyor](https://ci.appveyor.com/api/projects/status/qnnp26vv2c71f560?svg=true)](https://ci.appveyor.com/project/sbinet/hep)
[![codecov](https://codecov.io/gh/go-hep/hep/branch/main/graph/badge.svg)](https://codecov.io/gh/go-hep/hep)
[![License](https://img.shields.io/badge/License-BSD--3-blue.svg)](https://go-hep.org/license)
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.597940.svg)](https://doi.org/10.5281/zenodo.597940)
[![JOSS Paper](/images/joss-badge.svg)](https://doi.org/10.21105/joss.00372)
[![Binder](https://mybinder.org/badge.svg)](https://mybinder.org/v2/gh/go-hep/binder/master)

[Go](https://golang.org) brings the fast edit-compile-run cycle that interpreted language users enjoy and the runtime efficiency that compiled languages users expect.
[Go-HEP](https://go-hep.org) provides the needed `HEP` oriented packages on top of this concurrency-enabled language.
 
`Go-HEP` currently sports the following packages:

- [go-hep.org/x/hep/brio](https://go-hep.org/x/hep/brio): a toolkit to generate serialization code
- [go-hep.org/x/hep/fads](https://go-hep.org/x/hep/fads): a fast detector simulation toolkit
- [go-hep.org/x/hep/fastjet](https://go-hep.org/x/hep/fastjet): a jet clustering algorithms package (WIP)
- [go-hep.org/x/hep/fit](https://go-hep.org/x/hep/fit): a fitting function toolkit (WIP)
- [go-hep.org/x/hep/fmom](https://go-hep.org/x/hep/fmom): a 4-vectors library
- [go-hep.org/x/hep/fwk](https://go-hep.org/x/hep/fwk): a concurrency-enabled framework
- [go-hep.org/x/hep/groot](https://go-hep.org/x/hep/groot): a pure [Go](https://golang.org) package for [ROOT](https://root.cern.ch) I/O (WIP)
- [go-hep.org/x/hep/hbook](https://go-hep.org/x/hep/hbook): histograms and n-tuples (WIP)
- [go-hep.org/x/hep/hplot](https://go-hep.org/x/hep/hplot): interactive plotting (WIP)
- [go-hep.org/x/hep/hepmc](https://go-hep.org/x/hep/hepmc): `HepMC` in pure [Go](https://golang.org) (EDM + I/O)
- [go-hep.org/x/hep/hepevt](https://go-hep.org/x/hep/hepevt): `HEPEVT` bindings
- [go-hep.org/x/hep/heppdt](https://go-hep.org/x/hep/heppdt): `HEP` particle data table
- [go-hep.org/x/hep/lcio](https://go-hep.org/x/hep/lcio): read/write support for `LCIO` event data model
- [go-hep.org/x/hep/lhef](https://go-hep.org/x/hep/lhef): Les Houches Event File format
- [go-hep.org/x/hep/rio](https://go-hep.org/x/hep/rio): `go-hep` record oriented I/O
- [go-hep.org/x/hep/sio](https://go-hep.org/x/hep/sio): basic, low-level, serial I/O used by `LCIO`
- [go-hep.org/x/hep/slha](https://go-hep.org/x/hep/slha): `SUSY` Les Houches Accord I/O
- [go-hep.org/x/hep/xrootd](https://go-hep.org/x/hep/xrootd): [XRootD](http://xrootd.org) client in pure [Go](https://golang.org)
- [go-hep.org/x/cgo/croot](https://go-hep.org/x/cgo/croot): bindings to a subset of [ROOT](https://root.cern.ch) I/O

## Motivation

Writing analyses in HEP involves many steps and one needs a few tools to successfully carry out such an endeavour.
But -- at minima -- one needs to be able to read and write [ROOT](https://root.cern) files to be able to interoperate with the rest of the HEP community or to insert one's work into an already existing analysis pipeline.

Go-HEP provides this necessary interoperability layer, in the [Go](https://golang.org) programming language.
This allows physicists to leverage the great concurrency primitives of Go, together with the surrounding tooling and software engineering ecosystem of Go, to implement physics analyses.

## Installation

`Go-HEP` packages are installable via the `go get` command:

```sh
$ go get go-hep.org/x/hep/fads
```

Just select the package you are interested in and `go get` will take care of fetching, building and installing it, as well as its dependencies, recursively.

`go-hep` is available on all [Go](https://golang.org) supported platforms, mainly:

- `linux-{amd64,386,arm,arm64,mips,mips32,...}`
- `darwin-{amd64,386,...}`
- `windows-{amd64,386,...}`
- `freebsd-{amd64,386,...}`


Binaries for `go-hep` are also available here: [go-hep.org/dist](/dist)

## License

All `go-hep` code is released under a [BSD-3 license](/license).

## Authors and Contributors

``go-hep`` was primarily written by Sebastien Binet ([@sbinet](https://github.com/sbinet)).
The complete [CONTRIBUTORS](https://github.com/go-hep/license/blob/master/CONTRIBUTORS) and [AUTHORS](https://github.com/go-hep/license/blob/master/AUTHORS) list can be consulted on the dedicated [license](https://github.com/go-hep/license) repository.

## Support or Contact

Having trouble with ``Go-HEP``?
 
Check out the documentation:

- on [pkg.go.dev](https://pkg.go.dev/go-hep.org/x/hep),
- or the [tutos](https://go-hep.org/tutos).

There's also the forum:

- [~sbinet/go-hep@lists.sr.ht](mailto:~sbinet/go-hep@lists.sr.ht) (by mail)
- [lists.sr.ht/~sbinet/go-hep](https://lists.sr.ht/~sbinet/go-hep) (via the web interface),

to discuss about `Go-HEP` or ask for help.

## About

The `go-hep` logo was made by [Tom Ingebretsen](https://thenounproject.com/tomplusplus) from the [Noun Project](https://thenounproject.com).
