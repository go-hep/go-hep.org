---
title: "Go HEP release 0.23.0"
date: 2020-02-21T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.3678646.svg)](https://doi.org/10.5281/zenodo.3678646)


Release [`v0.23.0`](https://github.com/go-hep/hep/tree/v0.23.0) is fresh from the oven.

This release contains a couple of `groot` related features and bugfixes.

## groot

`groot` gained 2 new commands:

- [cmd/root2fits](https://go-hep.org/x/hep/cmd/root2fits) converts a ROOT `TTree` into a [FITS](https://en.wikipedia.org/wiki/FITS) table:

```
$> root2fits -h
root2fits converts the content of a ROOT tree to a FITS (binary) table.

Usage: root2fits [OPTIONS] -f input.root

Example:

 $> root2fits -f ./input.root -t tree

Options:
  -f string
    	path to input ROOT file name
  -o string
    	path to output FITS file name (default "output.fits")
  -t string
    	name of the ROOT tree to convert
```

- [cmd/fits2root](https://go-hep.org/x/hep/cmd/fits2root) converts a [FITS](https://en.wikipedia.org/wiki/FITS) table into a ROOT `TTree`:

```
$> fits2root -h
fits2root converts the content of a FITS table to a ROOT file and tree.

Usage: fits2root [OPTIONS] -f input.fits

Example:

 $> fits2root -f ./input.fits -t MyHDU

Options:
  -f string
    	path to input FITS file name
  -o string
    	path to output ROOT file name (default "output.root")
  -t string
    	name of the FITS table to convert
```

Also, many `root-xyz` commands have been refactored into a simple shim executable that calls into the new [groot/rcmd](https://go-hep.org/x/hep/groot/rcmd) so users (and `groot` tests) can more easily customize or re-use, say, `root-ls`, `root-dump`, ... through an API instead of spawning a sub-process.
This refactor is especially useful for [rcmd.Merge](https://pkg.go.dev/go-hep.org/x/hep/groot/rcmd#Merge), where users can register their own `merge` strategy for their own types (by implementing the [root.Merger](https://pkg.go.dev/go-hep.org/x/hep/groot/root#Merger) interface.)

`groot/rtree` can now correctly read multi-leaves branches.

And support for `TLeafF16`, `TLeafD32`, `TProcessID`, `TRef`, `TRefArray`, `TVector2`, `TVector3`, `TLorentzVector` and `TFeldmanCousins` has been added.

## hplot

- fixes for `HLine` and `VLine` have been applied: `HLine`/`VLine` outside a canvas were incorrectly drawn.

## xrootd

- migration to `gokrb5/v8`

## ChangeLog

* [398059a](/commit/398059a) groot/rbase: add initial support for tracking Refs
* [8c42445](/commit/8c42445) groot: introduce root.UIDer interface
* [8d26a4a](/commit/8d26a4a) groot{,rdict,rvers}: first stab at TBits
* [4b94d19](/commit/4b94d19) groot{,rdict,rvers}: first stab at TProcessUUID
* [84a6e03](/commit/84a6e03) ci: bump appveyor to Go-1.13
* [6e0dddd](/commit/6e0dddd) groot{,rdict,rphys,rvers}: add TFeldmanCousins
* [be0efa3](/commit/be0efa3) groot/rphys: properly handle version-2 of TVector{2,3}
* [c9a6535](/commit/c9a6535) groot{,rdict,rphys,rvers}: add TVector2
* [77598d3](/commit/77598d3) hplot: fix HLine/VLine for out of canvas lines
* [64086e9](/commit/64086e9) hep: update Go modules
* [20a7b06](/commit/20a7b06) hep,xrootd/xrdproto/auth: migrate to gokrb5/v8
* [4900ca2](/commit/4900ca2) groot: implement TVector3, TLorentzVector
* [808cc5b](/commit/808cc5b) groot: add TRefArray
* [dbf217c](/commit/dbf217c) groot: add first stab at a TRef implementation
* [683c737](/commit/683c737) hbook: fix doc-example of H1D.Integral after Binning1D API migration
* [c1fb32e](/commit/c1fb32e) cmd/root2arrow: improve automatic installation of arrow-cat
* [237736a](/commit/237736a) groot/{internal,rdict}: add list-groot-{sinfos,types}
* [12725a7](/commit/12725a7) groot: generate textual representation of streamers
* [9ae982c](/commit/9ae982c) groot/rdict: handling of Long64_t members in ROOT->groot streamer generation
* [ca69d60](/commit/ca69d60) groot/rdict: better handling of groot-wrapped ROOT types in ROOT->groot streamer generation
* [c19782e](/commit/c19782e) groot/rdict: add range-parsing to StreamerElement unmarshaling
* [499a8f4](/commit/499a8f4) groot/riofs: add generation of std::map-based ROOT data tree
* [2376dc0](/commit/2376dc0) groot/riofs/gendata: simplify code generation
* [d5830ae](/commit/d5830ae) groot/{rdict,riofs,rtree}: improve STL-container name arguments parsing
* [7118537](/commit/7118537) groot/{rdict,rtree}: rename StreamerSTL.STLVectorType into STLType
* [8e85fb9](/commit/8e85fb9) xroot/xrdproto/auth/krb5: migrate to gokrb5.v7
* [a44b1ea](/commit/a44b1ea) hep: go mod tidy
* [3216604](/commit/3216604) hplot: make sure embedmd is installed+required
* [fb7f30f](/commit/fb7f30f) hep: remove link to opencollective
* [64d644a](/commit/64d644a) groot: implement r/w TLeafF16
* [d17364c](/commit/d17364c) groot: implement r/w TLeafD32
* [c9904cb](/commit/c9904cb) groot/rdict: add support for TLeaf{D32,F16} comment/range parsing
* [98e75b5](/commit/98e75b5) groot{,/riofs,/rvers}: add support for TLeafF16 and TLeafD32
* [5621593](/commit/5621593) groot/{cmd/root-diff,rcmd}: refactor root-diff into rcmd.Diff
* [4c1484a](/commit/4c1484a) groot/rcmd: refactor Dump into a dumpCmd struct
* [9e53d27](/commit/9e53d27) groot/{cmd/root-cp,rcmd}: export root-cp as rcmd.Copy
* [a1734ed](/commit/a1734ed) groot/{cmd/root-merge,rcmd}: export root-merge as rcmd.Merge
* [6851230](/commit/6851230) cmd/arrow2root: use rcmd.Dump
* [944c4a9](/commit/944c4a9) cmd/fits2root: use groot/rcmd.Dump in tests
* [e20ccab](/commit/e20ccab) groot: export rcmd package
* [098e848](/commit/098e848) groot: more pervasive use of rtests.RunCxxROOT
* [af6ef10](/commit/af6ef10) groot/rtree: handle merging of Atlas flat-tuples
* [1df725b](/commit/1df725b) groot/riofs: add support for TProcessID StreamerInfo
* [0e1f43f](/commit/0e1f43f) groot/{rbase,rtypes}: implement TProcessID
* [1978a74](/commit/1978a74) groot{,/rvers}: add support for TProcessID
* [ebbd064](/commit/ebbd064) groot/rdict: better error message
* [d3b1b41](/commit/d3b1b41) groot/rmeta: better error message
* [cee0755](/commit/cee0755) groot/rtree: add more Branch.setAddress tests
* [aad9221](/commit/aad9221) groot/riofs: use rtests for generating testdata
* [98fe128](/commit/98fe128) groot/{internal/rcmd,riofs}: add testfile w/ & w/o padding
* [4419621](/commit/4419621) groot/internal/rcmd: properly display branches w/ multi-leaves in root-dump
* [e25f2ad](/commit/e25f2ad) groot/rtree: support reading multi-leaves branches
* [3c58df0](/commit/3c58df0) cmd/root2fits: first import
* [f15977d](/commit/f15977d) cmd/fits2root: first import
* [3e92a08](/commit/3e92a08) all: 2020 is the year of the Gopher



