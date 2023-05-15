---
title: "Go HEP release 0.33.0"
date: 2023-05-05T15:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.7928500.svg)](https://doi.org/10.5281/zenodo.7928500)

Release [`v0.33.0`](https://github.com/go-hep/hep/tree/v0.33.0) is out of the oven.

This release contains no major API breakage, but all around improvements and bug fixes.

## cmd

- [cmd/hepmc2root](https://pkg.go.dev/go-hep.org/x/hep@v0.33.0/cmd/hepmc2root): is a new command to automatically convert HepMC2 ASCII files to flat n-tuples ROOT files.

```
$> hepmc2root -h
hepmc2root converts a HepMC2 ASCII file into a ROOT file and (flat) tree.

Usage: hepmc2root [OPTIONS] hepmc.ascii

Example:

$> hepmc2root ./hepmc.ascii
$> hepmc2root -o out.root -t mytree ./hepmc.ascii

Options:
  -o string
    	path to output ROOT file name (default "out.root")
  -t string
    	name of the output tree (default "tree")
```

## hplot

`hplot` has been migrated to the new [Gio](https://gioui.org) API, and needs [Gonum/plot@v0.13.0](https://pkg.go.dev/gonum.org/v1/plot@v0.13.0).

Slight differences are to be expected in the rendering of plots.
With the new `Gio` version we are using, you may need development headers for [Vulkan](https://www.vulkan.org) on [Linux](https://kernel.org/).
See the installation instructions of Gio for more details: [Gio Installation](https://gioui.org/doc/install)

- fix a bug in `hplot.TiledPlot.Plot(i, j int)` where the `*hplot.Plot` returned by `Plot(i,j)` was actually the conjugate of what the documentation specified. Now, `Plot(i,j)` returns the plot as documented.

## sliceop

- introduce a generic `sliceop.Resize` that can reuse `[]T` capacity
- introduce generic versions of `sliceop.Take` and `sliceop.Find`

## Changelog

* [cf2d1a6d](/commit/cf2d1a6d) ci: pin to ubuntu-20.04 for Gio's benefit
* [219c13d4](/commit/219c13d4) hplot,pawgo: migrate to latest Gio API
* [97475faf](/commit/97475faf) ci: install vulkan stack for Gio
* [f1469c35](/commit/f1469c35) all: bump Gio, astrogo/fitsio@v0.3.0, klauspost/compress@v1.16.5 and gonum@v0.13
* [ef35e3ac](/commit/ef35e3ac) hplot: fix TiledPlot.Plot to behave as documented
* [229691c8](/commit/229691c8) ci/mk-release: update for new Go-1.20 build requirements
* [faaf33f7](/commit/faaf33f7) ci: use Go cache from setup-go
* [a25161e3](/commit/a25161e3) ci: update gha-actions
* [6b610132](/commit/6b610132) all: bump gokrb5@v8.4.4 and x/exp
* [12fafb41](/commit/12fafb41) all: bump x/net@v0.7.0
* [e0361d0d](/commit/e0361d0d) all: bump x/image@v0.5.0
* [89566f9e](/commit/89566f9e) xrootd: use crypto/rand instead of math/rand
* [d0b7e3e3](/commit/d0b7e3e3) csvutil: remove deprecated use of os.SEEK_END
* [2d85dea0](/commit/2d85dea0) ci: bump staticcheck
* [13370e03](/commit/13370e03) all: bump dependencies
* [c948dc89](/commit/c948dc89) all: bump to Go-1.20, drop Go-1.18
* [5b6e0488](/commit/5b6e0488) hepmc/go-hepmc-dump: use errors.Is to properly detect io.EOF
* [5e5fc636](/commit/5e5fc636) hepmc: use %w to report errors
* [f2987e91](/commit/f2987e91) cmd/lhef2hepmc: add -keep-all flag to keep intermediaries
* [ca416ca0](/commit/ca416ca0) all: 2023 is the year of the Gopher
* [7cd9c25e](/commit/7cd9c25e) all: bump golang/x repos
* [46778d8e](/commit/46778d8e) all: use internal/diff to display diffs
* [e79cf64b](/commit/e79cf64b) internal/diff: first import
* [063750f7](/commit/063750f7) cmd/hepmc2root,hepmc{,/rootcnv}: implement HepMC-ROOT roundtrip
* [66b2a3f4](/commit/66b2a3f4) cmd/hepmc2root: improve documentation
* [87e70954](/commit/87e70954) groot/{riofs,rtree}: de-duplicate automatically generated This-streamers
* [100af9da](/commit/100af9da) cmd/hepmc2root: first import
* [aad018e2](/commit/aad018e2) sliceop: generalize Take and Find signatures
* [bfaadd23](/commit/bfaadd23) sliceop: add Resize

