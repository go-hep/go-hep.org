+++
title = "Go HEP release 0.1"
date = "2017-04-06T14:13:32+02:00"
weight = 10
+++

It's with great pleasure that we can announce that the `Go-HEP` release `v0.1` is out:

- Changelog: https://github.com/go-hep/hep/releases/v0.1
- Sources (zip): https://github.com/go-hep/hep/archive/v0.1.zip
- Sources (tar): https://github.com/go-hep/hep/archive/v0.1.tar.gz

## What's new

What's new?
Well, lots of things since it is our very first release.

But since the merge of the `github.com/go-hep/xyz` repositories into the mono-repository `go-hep/hep` and the introduction of the `go-hep.org/x/hep/...` vanity imports, quite a few things.

First: as we have tagged a release, `Go-HEP` is now citable.
Thanks to [Zenodo](https://zenodo.org), we have a DOI: 

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.495433.svg)](https://doi.org/10.5281/zenodo.495433)

### Improved ROOT support

Support of [ROOT](https://root.cern.ch) files and reading their content dramatically improved:

- read `TH1x` histograms
- read `TH2x` histograms
- read `TGraph`, `TGraphErrors` and `TGraphAsymmErrors`
- read (flat) `TTrees`
- scan/iterate/select content of `TTrees`
- improved the [root-ls](https://godoc.org/go-hep.org/x/hep/rootio/cmd/root-ls) command display
- introduction of [root-srv](https://godoc.org/go-hep.org/x/hep/rootio/cmd/root-srv), a command to browse `ROOT` files' content.
  One can also plot `TH1x`, `TH2x`, `TGraph{,{,Asymm}Errors}` and (flat) `TTrees`.
- `root-srv` is also served on [Google AppEngine](https://cloud.google.com/appengine/):  [`rootio-inspector.appspot.com`](http://rootio-inspector.appspot.com). This allows to inspect `ROOT` files without having `ROOT` nor `Go-HEP` installed.

![rootio-inspector](/images/news/release-0.1/rootio-inspector-gui.png)

### Improved fads & hbook

[go-hep/hbook](https://go-hep.org/x/hep/hbook) gained inter-operability with [YODA](https://yoda.hepforge.org/) (Yet more Objects for Data Analysis).
`hbook` can now save `hbook.H1D`, `hbook.H2D`, ... into `YODA` ASCII files and load `hbook.H1D`, ... from `YODA` ASCII files.

Two new commands, [yoda2rio](https://godoc.org/go-hep.org/x/hep/cmd/yoda2rio) and [rio2yoda](https://godoc.org/go-hep.org/x/hep/cmd/rio2yoda), were also introduced to allow converting `YODA` files to [rio](https://godoc.org/go-hep.org/x/hep/rio) ones, and back.

This allowed to create a new command: [`fads-rivet-mc-generic`](https://godoc.org/go-hep.org/x/hep/fads/cmd/fads-rivet-mc-generic).
This command reproduces the default MonteCarlo analysis from [Rivet](https://rivet.hepforge.org): [MC_GENERIC.cc](https://rivet.hepforge.org/trac/browser/src/Analyses/MC_GENERIC.cc).
It was nice to see the results were the same :) (except that `fads` was twice faster.)

![](https://raw.githubusercontent.com/go-hep/hep/master/fit/testdata/h1d-gauss-plot.png)

### Miscellaneous

As always, the usual string of fixes, documentation improvements and other things were applied to the repository.
Noteworthy, the `sio` package has seen some attention, preparing for a pure-Go `lcio` package for inter-operability with the [LCIO](https://github.com/iLCSoft/LCIO) C++ I/O library for the future linear collider.

## Next steps

The plan is to provide support for reading of `TTrees` containing user classes.
This implies being able to interact with the `ROOT` streamers that are stored in the `TFiles`.
That would be a solid `v0.2` release.

Release `v0.3` (or `v0.4`) should also get the first iteration(s) on getting write support for `ROOT` files.
On the non-ROOT front, future releases could perhaps integrate the `rio` storage with `hbook/ntup.Ntuple` so one could get a complete, pure-Go stack with data persistency and an analysis workstation (`pawgo` and its web-based equivalent.)

In the meantime, please give `v0.1` a spin, file [issues](https://github.com/go-hep/hep/issues), send patches (via [pull requests](https://github.com/go-hep/hep/pulls)) and/or discuss anything `Go-HEP` related on the [go-hep](mailto:go-hep@googlegroups.com) mailing list.
