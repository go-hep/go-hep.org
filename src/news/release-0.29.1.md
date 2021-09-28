---
title: "Go HEP release 0.29.1"
date: 2021-09-28T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.5534304.svg)](https://doi.org/10.5281/zenodo.5534304)

Release [`v0.29.1`](https://github.com/go-hep/hep/tree/v0.29.1) is out of the oven.

This release contains no major API breakage, but all around improvements and bug fixes.
Most notably, this release uses the new `gonum/plot@v0.10.0` release that fixes a bunch of graphics bugs and performance improvements in plots drawing (PDF, PNG and SVG.)

## arrow

We are now using a "vendored" version of `github.com/apache/arrow/go` without all the `flight` support, to cut down on the amount of dependencies this was dragging in.

## groot

- `groot/{rdict,rtree}:` handle groot types mirroring TObject inheritance tree
- `groot/rtree:` reduce memory usage in tree-write
- `groot/{riofs,rtree}:` handle small/big file thresholds for long-running baskets
- `groot/rhist`: handle nil-title in YODA histograms
- `cmd/yoda2root:` handle YODA files with a directory structure
- `groot/rhist:` implement YODA (un)marshaler for TGraph{,Asymm}{,Errors}
- `groot/rhist:` correctly write out TGraph{,Asymm}{,Errors}

## hbook

- `hbook:` introduce Rand1D. `Rand1D` represents a 1D distribution created from a `hbook.H1D` histogram.

## hplot

- `hplot:` add Label plotter. This plotter allows user to put labels on a plot in "normalized" or "data" coordinates.

## xrootd

- `xrootd:` fix race in file

## AOB

That's all for today.
Next cycle will probably see some work on the `RNTuple` front.

## Changelog

* [dbb6f2e5](/commit/dbb6f2e5) xrootd: fix race in file
* [2811be8a](/commit/2811be8a) all: apply lint fixes
* [523ebcd8](/commit/523ebcd8) ci: bump golangci-lint@v1.42
* [1e14bc8e](/commit/1e14bc8e) all: bump go-mmap@v0.6.0
* [a39b0282](/commit/a39b0282) all: update go-mmap@v0.5.1, go-cmp@v0.5.6, xz@v0.5.10, x-tools@v0.1.6
* [b636438e](/commit/b636438e) fit: use cmpimg.CheckPlotApprox where needed
* [fb956a53](/commit/fb956a53) all: update reference files
* [64a270d1](/commit/64a270d1) hplot: migrate to new gonum/plot font cache
* [8c95ddff](/commit/8c95ddff) all: bump to gonum.org/v1/{gonum@v0.9,plot@v0.10}
* [26d966e5](/commit/26d966e5) ci: fix (build+module) cache of paths
* [685d3dfb](/commit/685d3dfb) ci: add git-config for Windows
* [660723ff](/commit/660723ff) ci: use GHA+Windows
* [c1c804eb](/commit/c1c804eb) ci: use Go-1.16 for AppVeyor
* [baf51613](/commit/baf51613) all: drop use of io/ioutil
* [c467a582](/commit/c467a582) groot/cmd/root-print: add -regen mechanism
* [cb8ccf5d](/commit/cb8ccf5d) all: bump klauspost/compress@v1.13.4
* [72b99d88](/commit/72b99d88) all: add Go-1.17, drop Go-1.15
* [576edb60](/commit/576edb60) groot/rhist: add att{line,fill,marker} fields to graph
* [4622cac6](/commit/4622cac6) groot/rhist: correctly write out TGraph{,Asymm}{,Errors}
* [ceb43b51](/commit/ceb43b51) groot/rhist: implement YODA (un)marshaler for TGraph{,Asymm}{,Errors}
* [a721c455](/commit/a721c455) all: apply gofmt -w -s
* [b410620b](/commit/b410620b) groot: bump to pierrec/lz4@v4.1.8, klauspost/compress@v1.13.1
* [dcbc8102](/commit/dcbc8102) ci: re-enable static builds where possible
* [4a636e4b](/commit/4a636e4b) ci: use netgo tag in mk-release
* [36092fc2](/commit/36092fc2) all: bump to ROOT-6.24.00
* [c2bd848c](/commit/c2bd848c) cmd/yoda2root: handle YODA files with a directory structure
* [b19b7758](/commit/b19b7758) groot/rhist: handle nil-title in YODA histograms
* [84ca3d76](/commit/84ca3d76) groot/riofs: use SMHiggsToZZTo4L.root from ccxrootdgotest instead of eospublic.cern.ch
* [226a27a5](/commit/226a27a5) ci: add support for Ubuntu-20 and headless X11 mode
* [7024f736](/commit/7024f736) hep: use light-arrow instead of apache/arrow
* [8aeeba84](/commit/8aeeba84) hep: remove Travis-CI badge
* [eda01ec6](/commit/eda01ec6) all: update license blurb
* [73dbb27c](/commit/73dbb27c) all: prepare for go:build directive
* [a2762f6b](/commit/a2762f6b) fastjet/internal/plot: remove old plot code
* [947bb632](/commit/947bb632) hep: bump klauspost/compress@v1.11.8
* [4fe1e1ab](/commit/4fe1e1ab) all: bump to Go-1.16
* [720a7c27](/commit/720a7c27) doc: add link to forum discussion
* [36c2894d](/commit/36c2894d) hep: bump npyio@v0.5.2
* [bd26c299](/commit/bd26c299) fmom: use func-based spatial/r3 API
* [160f5eb2](/commit/160f5eb2) hep: bump klauspost/compress to v1.11.7
* [cc69431f](/commit/cc69431f) all: 2021 is the year of the Gopher
* [8fb7154d](/commit/8fb7154d) hplot,pawgo: use new vg-gio
* [0410c37e](/commit/0410c37e) all: use main instead of master
* [4b45b719](/commit/4b45b719) hplot: add Label plotter
* [262cb504](/commit/262cb504) fmom: implement VecOf
* [9a1d24de](/commit/9a1d24de) fmom: add SetPtEtaPhi{E,M} to PxPyPzE
* [e07f1889](/commit/e07f1889) hbook: introduce Rand1D
* [2edf3b7b](/commit/2edf3b7b) groot/{riofs,rtree}: handle small/big file thresholds for long-running baskets
* [1d1c78d1](/commit/1d1c78d1) groot/rtree: reduce memory usage in tree-write
* [8c71fb89](/commit/8c71fb89) groot/riofs: fix overlaping blocks w/ big-file mode
* [43e74de9](/commit/43e74de9) all: bump versions of arrow, gokrb5, klauspost/compress, npyio and ql
* [f5eb1e63](/commit/f5eb1e63) ci: bump lint to v2
* [a6f1eec6](/commit/a6f1eec6) all: use xyz_example_test Go idiom
* [853cfac7](/commit/853cfac7) fit: remove exponential example from darwin tests
* [c976e7f0](/commit/c976e7f0) all: bump gonum/plot@v0.8.1
* [b2e31fc2](/commit/b2e31fc2) xrootd: update for new ccxrootd test server
* [55f031ef](/commit/55f031ef) groot/rdict: make sure StreamerOf panics when it must (chan,int,uint,func)
* [4759357b](/commit/4759357b) groot/{rdict,rtree}: handle groot types mirroring TObject inheritance tree

