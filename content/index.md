+++
date = "2016-11-24T09:19:32+01:00"
title = "Welcome to Go-HEP"
type = "index"
weight = 0
tags = ["sphinx","documentation"]

+++

`go-hep` is a set of libraries and applications allowing High Energy Physicists to write efficient analysis code in the [Go](https://golang.org) programming language.


[Go](https://golang.org) brings the fast edit-compile-run cycle that interpreted language users know and the runtime efficiency that compiled languages users expect.
[go-hep](https://go-hep.org) provides the needed `HEP` oriented packages on top of this concurrency-enabled language.
 
`go-hep` currently sports the following packages:

- [go-hep.org/x/fads](https://go-hep.org/x/fads): a fast detector simulation toolkit
- [go-hep.org/x/fastjet](https://go-hep.org/x/fastjet): a jet clustering algorithms package (WIP)
- [go-hep.org/x/fmom](https://go-hep.org/x/fmom): a 4-vectors library
- [go-hep.org/x/fwk](https://go-hep.org/x/fwk): a concurrency-enabled framework
- [go-hep.org/x/hbook](https://go-hep.org/x/hbook): histograms and n-tuples (WIP)
- [go-hep.org/x/hplot](https://go-hep.org/x/hplot): interactive plotting (WIP)
- [go-hep.org/x/hepmc](https://go-hep.org/x/hepmc): `HepMC` in pure [Go](https://golang.org) (EDM + I/O)
- [go-hep.org/x/hepevt](https://go-hep.org/x/hepevt): `HEPEVT` bindings
- [go-hep.org/x/heppdt](https://go-hep.org/x/heppdt): `HEP` particle data table
- [go-hep.org/x/lhef](https://go-hep.org/x/lhef): Les Houches Event File format
- [go-hep.org/x/croot](https://go-hep.org/x/croot): bindings to a subset of [ROOT](https://root.cern.ch) I/O
- [go-hep.org/x/rio](https://go-hep.org/x/rio): `go-hep` record oriented I/O
- [go-hep.org/x/sio](https://go-hep.org/x/sio): `LCIO` I/O
- [go-hep.org/x/slha](https://go-hep.org/x/slha): `SUSY` Les Houches Accord I/O

## Installation

`go-hep` packages are installable via the `go get` command:

```sh
$ go get go-hep.org/x/fads
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
The complete [CONTRIBUTORS](https://github.com/go-hep/license/blob/master/CONTRIBUTORS) and [AUTHORS](https://github.com/go-hep/license/blob/master/AUTHORS) list can be consulted on the dedicated [license](https://go-hep.org/license) repository.

## Support or Contact

Having trouble with ``go-hep``?
 
Check out the documentation:

- on [godoc.org](https://godoc.org/?q=go-hep.org),
- or the [tutos](https://go-hep.org/tutos).

You can also contact us at [go-hep@googlegroups.com](mailto:go-hep@googlegroups.com) and we’ll help you sort it out.

The main mailing list for `go-hep` is hosted on [googlegroups](https://groups.google.com/forum/#!forum/go-hep).
You can subscribe to the forum without having a `GMail` account, by just sending a mail to [go-hep+subscribe@googlegroups.com](mailto:go-hep+subscribe@googlegroups.com?subject=subscribe) with `subscribe` as a subject.
