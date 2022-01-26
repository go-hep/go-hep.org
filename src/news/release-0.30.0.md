---
title: "Go HEP release 0.30.0"
date: 2022-01-26T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.5905884.svg)](https://doi.org/10.5281/zenodo.5905884)

Release [`v0.30.0`](https://github.com/go-hep/hep/tree/v0.30.0) is out of the oven.

This release contains no major API breakage, but all around improvements and bug fixes.

## fwk

- updated `fwk-app` to use the latest `gonuts/commander` version that provides auto-completion of commands, flags and sub-commands.

## groot

- add support for `TBranchObject` and `TLeafObject`
- add support for `TLorentzVector` (in `root-dump` and friends)
- add support for `TNtuple` and `TNtupleD` (in `root-dump` and friends)

## hplot

- fixed a race in `hplot.New` where the global variable from `gonum/plot` was being modified without a critical section. This race appeared when one would create multiple `hplot.Plot` concurrently.

## sliceop

- improved performances of `sliceop/f64s.Take`. Do note that `sliceop/f64s` will probably be rewritten in terms of a generic implementation (and deprecated) once `Go-1.18` is generally available.

## AOB

That's all for today.
Next cycle will probably see some work on the `RNTuple` front.

## Changelog

* [baa058b5](/commit/baa058b5) bump to ROOT-6.24/06
* [bc000321](/commit/bc000321) all: bump peterh/liner, x/crypto, x/image, x/tools and ql
* [b57330a1](/commit/b57330a1) groot/{rcmd,riofs,rtree,rvers}: add rtree.Reader support for TNtuple{,D}
* [24ff587b](/commit/24ff587b) all: bump klauspost/compress@v1.14.2
* [d1e8a6fe](/commit/d1e8a6fe) all: 2022 is the year of the Gopher
* [505ab934](/commit/505ab934) fwk/cmd/fwk-app: use gonuts/commander@v0.3.x
* [c5711342](/commit/c5711342) hplot: fix race in hplot.New on gonum/plot.DefaultFont
* [eaeb641a](/commit/eaeb641a) groot/rphys: test RStreamer with Bind+Unmarshaler for TLorentzVector
* [679c8867](/commit/679c8867) groot/rphys: make Vector{2,3} implement fmt.Stringer
* [7851644b](/commit/7851644b) groot/rphys: make LorentzVector implement fmt.Stringer
* [07e2951c](/commit/07e2951c) groot/rdict: leverage rbytes.{Unm,M}arshaler implementation in Bind
* [e370fced](/commit/e370fced) groot/riofs: add generation of TLorentzVector testdata
* [c117054d](/commit/c117054d) groot/{rdict,rtree,rvers}: add support for T{Branch,Leaf}Object
* [975023e2](/commit/975023e2) groot: remove old comment about write-mode of TTrees
* [7735669e](/commit/7735669e) sliceop/f64s: improve Take performances
* [4ca091cd](/commit/4ca091cd) all: bump x/tools@v0.1.7


