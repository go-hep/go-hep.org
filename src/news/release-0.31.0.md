---
title: "Go HEP release 0.31.0"
date: 2022-03-22T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.6376611.svg)](https://doi.org/10.5281/zenodo.6376611)

Release [`v0.31.0`](https://github.com/go-hep/hep/tree/v0.31.0) is out of the oven.

This release contains no major API breakage, but all around improvements and bug fixes.

This release drops support for Go-1.16 (and earlier).

## groot

- bump to `ROOT-6.26/00`
- implement a riofs Reader over HTTP(s), instead of downloading the whole file locally
- improve performances of `rcmd.Dump(tree)`
- add support for `TMultiGraph`
- add support for `TDatime`
- add support for `TGraphMultiErrors`
- add support for `TProfile{,2D}`
- add support for `TEfficiency`
- add support for `TConfidenceLevel`, `TLimit{,DataSource}`
- add support for `TF1` and `TFormula`
- add support for reading `std::vector<T, myalloc<T>>`
- introduced a new package `groot/rnpy`, to ease ROOT/NumPy conversions (like `groot/rarrow` does for Arrow)
- improve memory usage of `root2npy`, by materializing into memory only one column at a time
- streamline C++ templates parsing for STL containers (vector,deque,map,unordered_{,multi}{map,set}, pair, ...)

## xrootd

- improve the read performances 10-fold, using a finer-grained lock to protect session IDs

## AOB

That's all for today.
Last release we said there would probably be some work on the `RNTuple` front.
We didn't lie, but this didn't happen in the form we thought it would.
A surprise is in the works... stay tuned.

Next cycle will probably see some work on the `RNTuple` front, and that work appear in `main`.

## Changelog

* [b02724a1](/commit/b02724a1) all: bump x/crypto
* [6c359638](/commit/6c359638) all: drop old +build foo stanza, use //go:build
* [24817699](/commit/24817699) all: bump klauspost/compress@v1.15.1 and pierrec/lz4@v4.1.14
* [59c6c675](/commit/59c6c675) all: bump gonum/gonum@v0.11.0, gonum/plot@v0.11.0
* [c6f7c48c](/commit/c6f7c48c) all: drop Go-1.16, add Go-1.18
* [7909cd7d](/commit/7909cd7d) groot/internal/httpio: add a pool of http.Request
* [0e0b9931](/commit/0e0b9931) groot/riofs/plugin/http: first stab at a caching+concurrent http-reader
* [53a256f5](/commit/53a256f5) groot/internal/httpio: first import
* [e6d5f90c](/commit/e6d5f90c) groot: bump to ROOT-6.26/00
* [6f88941b](/commit/6f88941b) xrootd: improve read performances 10-fold
* [c38575f0](/commit/c38575f0) groot/rcmd: improve Dump tree performances
* [ca83c453](/commit/ca83c453) groot/{internal/rtests,riofs}: better ACliC handling
* [e453fa7c](/commit/e453fa7c) groot/{rcmd,rhist,rvers}: add support for TMultiGraph
* [17474123](/commit/17474123) groot/{rbase,riofs,rvers}: add support for TDatime
* [0034b01b](/commit/0034b01b) groot/internal/rtests: make sure macros are run thru ACliC
* [0bd7c62e](/commit/0bd7c62e) groot/{rcmd,riofs,rmeta}: add support for read std::vector<T,my_alloc<T>>
* [551dabd7](/commit/551dabd7) groot/{rcmd,riofs}: add tests for more STL containers
* [f1aa9058](/commit/f1aa9058) groot/rtree: add rvar dispatch for map[K]V
* [b4678544](/commit/b4678544) groot/rdict: add r/w-streamer+type handling for TString
* [480971d2](/commit/480971d2) groot/rdict: handle TString streaming in collections
* [22da5159](/commit/22da5159) groot/rdict: use []T for disk storage of std::set<T>
* [34d6347f](/commit/34d6347f) groot/rdict: proper handling of 'This' streamers
* [0a85f502](/commit/0a85f502) groot: introduce rbytes.Header
* [e16891db](/commit/e16891db) groot/rcmd: add tests for POD data files
* [2a344611](/commit/2a344611) groot/rdict: generate RVersioner if not implemented
* [d23d2f42](/commit/d23d2f42) groot/{rdict,rmeta}: introduce CxxTemplate, streamline template parsing
* [be7ddfa6](/commit/be7ddfa6) groot/rnpy: new package to ease ROOT-Tree/NumPy conversion
* [37a9ac0e](/commit/37a9ac0e) cmd/root2npy: improve memory usage
* [04a2d5fc](/commit/04a2d5fc) groot: streamline r/w buffer ops
* [1b25234f](/commit/1b25234f) groot/{rdict,rhist,rvers}: add r/w support for TGraphMultiErrors
* [73b95060](/commit/73b95060) groot/{rhist,riofs,rvers}: add r/w support for TProfile{,2D}
* [449e81e7](/commit/449e81e7) groot/rdict: add initial support for enums in genGoType
* [e6f8d63d](/commit/e6f8d63d) groot/{rcmd,rdict,rhist,rvers}: add r/w support for TEfficiency
* [929ec73d](/commit/929ec73d) groot/rhist: add write support to TF1 and TFormula
* [129831f6](/commit/129831f6) groot/{rcmd,rdict,rhist,riofs,rvers}: add r/w support for TConfidenceLevel and TLimit{,DataSource}
* [ffc817e8](/commit/ffc817e8) root/rbytes: add WriteStdVectorF64
* [e488e92a](/commit/e488e92a) root/rbytes: add WriteObject
* [fa6827ea](/commit/fa6827ea) groot/{rdict,rhist,rvers}: add initial read support for TF1 and TFormula
* [d8d3eb53](/commit/d8d3eb53) groot/rbytes: add RBuffer.ReadStdVectorI32
* [17373164](/commit/17373164) groot/rbytes: add RBuffer.ReadStdVectorF64
* [ad66388a](/commit/ad66388a) cmd/root2{arrow,csv,npy}: use riofs.Dir to retrieve trees

