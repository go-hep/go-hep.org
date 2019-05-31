---
title: "Go HEP release 0.18.0"
date: 2019-05-01T0:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.3236434.svg)](https://doi.org/10.5281/zenodo.3236434)

Release [`v0.18.0`](https://github.com/go-hep/hep/tree/v0.18.0) is fresh-ish from the oven.

This release contains the beginning of an XRootD server in Go (thanks [Mikhail](https://github.com/EgorMatirov)!) as well as yet more `groot` improvements.

## fit

- fit: update for new `gonum.org/v1/gonum/optimize` API
- fit: use new gonum/{diff/fd,optimize} API

## groot

`groot/riofs.Open` and thus `groot.Open` have been modified to only open local files by default.
A "plugin" mechanism has been devised to be able to optionally install plugins to deal with, _e.g._, HTTP-served and XRootD-served files.
To enable these plugins, one just need to "blank import" the following packages:

```go
import (
	_ "go-hep.org/x/hep/groot/riofs/plugin/http"
	_ "go-hep.org/x/hep/groot/riofs/plugin/xrootd"
)
```

The `groot/cmd/root-xxx` commands have been modified to automatically import the above packages.
Users are expected to do the same if they want to transparently be able to read such files.
The rationale being that `groot/riofs` should by default try to always create static binaries: to read the HTTP- and XRootD-served files one needs to import `net/http` that require some amount of dynamic linking.

The `groot` package also gained a new command:

```
$> root-gen-type -h
Usage: root-gen-type [options] input.root

ex:
 $> root-gen-type -p mypkg -t MyType -o streamers_gen.go ./input.root

options:
  -o string
    	output file name
  -p string
    	package import path
  -t string
    	comma-separated list of (regexp) type names (default ".*")
  -v	enable verbose mode
```

`root-gen-type` automatically creates the streamer code from an input ROOT file (not unlike the old `genreflex` or the new `rootcling`) as well as a Go struct compatible with the data contained in the (C++, on-disk) streamers (kind of like the `C++` `TTree::MakeClass`.)

- groot: add support for TMap
- groot/rdict: add support for `StreamerObject{,Pointer}`
- groot/cmd/root-ls: lazily load data off disk
- groot/rvers: migrate to ROOT-6.16/00
- groot: add support for TClonesArray
- groot: introduce ROOT open plugin mechanism
- groot/cmd/root-gen-type: first import
- groot: support TArrayC, TArrayS and TArrayL
- groot/rdict: add support for slices of Object-s or ObjectAny-s
- groot: add initial support for Float16_t and Double32_t in root-gen-type

## hbook

- hbook/ntup: add an example on how to create histos from CSV
- hbook/ntup: add examples for Scan and ScanH1D methods

## hep

- hep: add Version to API

## hplot

- hplot: add support for H1D-logy

## xrootd

- xrootd: add ping, mkdir, rm, and rmdir support to the server
- xrootd/cmd/xrd-ls: fix output format
- xrootd/xrdproto/query: update example for new XRootD version


## AOB

Stay tuned! (and, as always, any kind of help (reviews, patches, nice emails, constructive criticism) deeply appreciated.)
