+++
date = "2016-11-24T09:19:32+01:00"
title = "Welcome to Go-HEP"
type = "index"
weight = 0
tags = ["sphinx","documentation"]

+++

`go-hep` is a set of libraries and applications allowing High Energy Physicists to write efficient analysis code in the [Go](https://golang.org) programming language.


[![Build Status](https://secure.travis-ci.org/go-hep/hep.png)](http://travis-ci.org/go-hep/hep)
[![GoDoc](https://godoc.org/go-hep.org/x/hep?status.svg)](https://godoc.org/go-hep.org/x/hep)
[![License](https://img.shields.io/badge/License-BSD--3-blue.svg)](https://go-hep.org/license)
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.597940.svg)](https://doi.org/10.5281/zenodo.597940)
[![JOSS Paper](/images/joss-badge.svg)](https://doi.org/10.21105/joss.00372)
[![Binder](https://mybinder.org/badge.svg)](https://mybinder.org/v2/gh/go-hep/binder/master)

[Go](https://golang.org) brings the fast edit-compile-run cycle that interpreted language users know and the runtime efficiency that compiled languages users expect.
[go-hep](https://go-hep.org) provides the needed `HEP` oriented packages on top of this concurrency-enabled language.
 
`go-hep` currently sports the following packages:

- [go-hep.org/x/hep/brio](https://go-hep.org/x/hep/brio): a toolkit to generate serialization code
- [go-hep.org/x/hep/fads](https://go-hep.org/x/hep/fads): a fast detector simulation toolkit
- [go-hep.org/x/hep/fastjet](https://go-hep.org/x/hep/fastjet): a jet clustering algorithms package (WIP)
- [go-hep.org/x/hep/fit](https://go-hep.org/x/hep/fit): a fitting function toolkit (WIP)
- [go-hep.org/x/hep/fmom](https://go-hep.org/x/hep/fmom): a 4-vectors library
- [go-hep.org/x/hep/fwk](https://go-hep.org/x/hep/fwk): a concurrency-enabled framework
- [go-hep.org/x/hep/hbook](https://go-hep.org/x/hep/hbook): histograms and n-tuples (WIP)
- [go-hep.org/x/hep/hplot](https://go-hep.org/x/hep/hplot): interactive plotting (WIP)
- [go-hep.org/x/hep/hepmc](https://go-hep.org/x/hep/hepmc): `HepMC` in pure [Go](https://golang.org) (EDM + I/O)
- [go-hep.org/x/hep/hepevt](https://go-hep.org/x/hep/hepevt): `HEPEVT` bindings
- [go-hep.org/x/hep/heppdt](https://go-hep.org/x/hep/heppdt): `HEP` particle data table
- [go-hep.org/x/hep/lcio](https://go-hep.org/x/hep/lcio): read/write support for `LCIO` event data model
- [go-hep.org/x/hep/lhef](https://go-hep.org/x/hep/lhef): Les Houches Event File format
- [go-hep.org/x/hep/rio](https://go-hep.org/x/hep/rio): `go-hep` record oriented I/O
- [go-hep.org/x/hep/rootio](https://go-hep.org/x/hep/rootio): a pure [Go](https://golang.org) package to for [ROOT](https://root.cern.ch) I/O (WIP) 
- [go-hep.org/x/hep/sio](https://go-hep.org/x/hep/sio): basic, low-level, serial I/O used by `LCIO`
- [go-hep.org/x/hep/slha](https://go-hep.org/x/hep/slha): `SUSY` Les Houches Accord I/O
- [go-hep.org/x/hep/xrootd](https://go-hep.org/x/hep/xrootd): [XRootD](http://xrootd.org) client in pure [Go](https://golang.org)
- [go-hep.org/x/cgo/croot](https://go-hep.org/x/cgo/croot): bindings to a subset of [ROOT](https://root.cern.ch) I/O

## Installation

`go-hep` packages are installable via the `go get` command:

```sh
$ go get go-hep.org/x/hep/fads
```

Just select the package you are interested in and `go get` will take care of fetching, building and installing it, as well as its dependencies, recursively.

`go-hep` is available on all [Go](https://golang.org) supported platforms, mainly:

- `linux-{amd64,386,arm,arm64,mips,mips32,...}`
- `darwin-{amd64,386,...}`
- `windows-{amd64,386,...}`
- `freebsd-{amd64,386,...}`

## License

All `go-hep` code is released under a [BSD-3 license](https://go-hep.org/license).

## Authors and Contributors

``go-hep`` was primarily written by Sebastien Binet ([@sbinet](https://github.com/sbinet)).
The complete [CONTRIBUTORS](https://github.com/go-hep/license/blob/master/CONTRIBUTORS) and [AUTHORS](https://github.com/go-hep/license/blob/master/AUTHORS) list can be consulted on the dedicated [license](https://github.com/go-hep/license) repository.

## Support or Contact

Having trouble with ``go-hep``?
 
Check out the documentation:

- on [godoc.org](https://godoc.org/?q=go-hep.org),
- or the [tutos](https://go-hep.org/tutos).

You can also contact us at [go-hep@googlegroups.com](mailto:go-hep@googlegroups.com) and we’ll help you sort it out.

The main mailing list for `go-hep` is hosted on [googlegroups](https://groups.google.com/forum/#!forum/go-hep).
You can subscribe to the forum without having a `GMail` account, by just sending a mail to [go-hep+subscribe@googlegroups.com](mailto:go-hep+subscribe@googlegroups.com?subject=subscribe) with `subscribe` as a subject.

## About

The `go-hep` logo was made by [Tom Ingebretsen](https://thenounproject.com/tomplusplus) from the [Noun Project](https://thenounproject.com).
