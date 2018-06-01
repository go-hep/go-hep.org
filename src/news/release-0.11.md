---
title: "Go HEP release 0.11"
date: 2018-04-05T15:30:51+02:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1213189.svg)](https://doi.org/10.5281/zenodo.1213189)

Release [`v0.11`](https://github.com/go-hep/hep/tree/v0.11) is fresh from the oven.

This release drops official support for Go-1.7 and earlier, adds support for Go-1.10.

This allowed to migrate the `fwk` package away from `golang.org/x/net/context`.
`fwk` now directly uses the `context.Context` type from the standard library.

## rootio

This release brings quite a few improvements in the `rootio` area:

- add support for empty ROOT files,
- add support for reading remote ROOT files over HTTP.
  This is implemented in a quite naive way, using `net/http.Get` to
	download the whole file under a temporary directory.
	See [go-hep/hep#142](https://github.com/go-hep/hep/issues/142) for ideas on
	how to improve this,
- add support for `TH1`-v6,
- add support for streamer-less `TDirectoryFile`,
- add support for `TBranchElements` in `cmd/root-dump`,
- add support for displaying `TH1`, `TH2` and `TGraph{,Error}s` in `cmd/root-dump`,
- add support for branches and leaves with names only differing by case in
  `cmd/root-gen-datareader`.
	Now, the `struct` type being generated contains fields whose names start
	with `ROOT_`.
	The original branch or leaf name is then appended after `ROOT_`, unmodified.
- add support for `TKey`s with a large binary payload,
- add preliminary support for creating new Go types at runtime directly from
  a `StreamerInfo` value,
- add more documentation about ROOT binary file format (`TFile`, `TKey` and a bit about
  `TStreamerInfo` so far.)

## hbook

- fixed a bug in the `YODA` ASCII file parsing code that would choke on
  analysis objects (histograms, scatters, ...) that contain whitespace.

## AOB

[LPC-Clermont](http://clrwww.in2p3.fr/) has funded a 5 months internship
student, starting now.
[Mohamed Amine El Gnaoui (@maloft)](https://github.com/maloft) will:

- implement reading all ROOT files created with ROOT-6,
- implement writing ROOT files, in a ROOT-6 compatible way,
- improve the read/write performances of Go-HEP to be on par with that of ROOT/C++ (using the builtin performance tools of the Go toolchain: pprof, execution tracer, ... and/or `linux perf`)
- extend the test suite of Go-HEP for reading and writing ROOT files,
- add benchmark tests for reading and writing ROOT files,
- document the ROOT file format as understood by Go-HEP.

Welcome to Go-HEP, [Amine](https://github.com/maloft)!
